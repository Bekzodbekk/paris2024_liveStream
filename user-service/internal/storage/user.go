package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"user-service/internal/repository"
	"user-service/token"

	pb "github.com/Bekzodbekk/protofiles/genproto/user"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db  *sql.DB
	rds *redis.Client
}

func NewUserRepo(db *sql.DB, rds *redis.Client) repository.UserRepository {
	return &UserRepo{
		db:  db,
		rds: rds,
	}
}

func (u *UserRepo) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &pb.User{}
	now := time.Now().Format(time.RFC3339)

	// Parolni hashlash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.CreateUserResponse{Success: false, Message: "Failed to hash password"}, err
	}

	err = u.db.QueryRow(
		"INSERT INTO users (username, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, username, role, created_at, updated_at",
		req.Username, string(hashedPassword), req.Role, now, now,
	).Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return &pb.CreateUserResponse{Success: false, Message: "Failed to create user"}, err
	}

	return &pb.CreateUserResponse{
		Success: true,
		Message: "User created successfully",
		User:    user,
	}, nil
}
func (s *UserRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userResp, err := s.GetUserByFilter(ctx, &pb.UserFilter{
		Username: req.Username,
	})

	if err != nil || len(userResp.Users) == 0 {
		return &pb.LoginResponse{
			Success: false,
			Message: "Invalid credentials",
		}, nil
	}

	user := userResp.Users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &pb.LoginResponse{
			Success: false,
			Message: "Invalid credentials",
		}, nil
	}

	accessToken, refreshToken, err := token.CreateTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.rds.Set(ctx, refreshToken, user.Id, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Success:      true,
		Message:      "Login successful",
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *UserRepo) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	userId, err := s.rds.Get(ctx, req.RefreshToken).Result()
	if err != nil {
		return &pb.RefreshTokenResponse{
			Success: false,
			Message: "Invalid refresh token",
		}, nil
	}

	userResp, err := s.GetUserById(ctx, &pb.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	user := userResp.User

	accessToken, refreshToken, err := token.CreateTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.rds.Set(ctx, refreshToken, user.Id, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}

	return &pb.RefreshTokenResponse{
		Success:      true,
		Message:      "Token refreshed successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	now := time.Now().Format(time.RFC3339)

	var err error
	if req.User.Password != "" {
		// Agar parol yangilanayotgan bo'lsa, uni hashlash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
		if err != nil {
			return &pb.UpdateUserResponse{Success: false, Message: "Failed to hash password"}, err
		}
		_, err = u.db.Exec(
			"UPDATE users SET username = $1, role = $2, password = $3, updated_at = $4 WHERE id = $5",
			req.User.Username, req.User.Role, string(hashedPassword), now, req.User.Id,
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = u.db.Exec(
			"UPDATE users SET username = $1, role = $2, updated_at = $3 WHERE id = $4",
			req.User.Username, req.User.Role, now, req.User.Id,
		)
	}

	if err != nil {
		return &pb.UpdateUserResponse{Success: false, Message: "Failed to update user"}, err
	}

	req.User.UpdatedAt = now

	return &pb.UpdateUserResponse{
		Success: true,
		Message: "User updated successfully",
		User:    req.User,
	}, nil
}
func (u *UserRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := u.db.Exec(
		"UPDATE users SET deleted_at = $1 WHERE id = $2",
		time.Now().Unix(), req.Id,
	)

	if err != nil {
		return &pb.DeleteUserResponse{Success: false, Message: "Failed to delete user"}, err
	}

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}
func (u *UserRepo) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &pb.User{}
	err := u.db.QueryRow(
		"SELECT id, username, role, created_at, updated_at FROM users WHERE id = $1 AND deleted_at = 0",
		req.Id,
	).Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetUserResponse{Success: false, Message: "User not found"}, nil
		}
		return &pb.GetUserResponse{Success: false, Message: "Failed to get user"}, err
	}

	return &pb.GetUserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User:    user,
	}, nil
}
func (u *UserRepo) GetUsers(ctx context.Context, req *pb.Void) (*pb.GetUsersResponse, error) {
	rows, err := u.db.Query("SELECT id, username, role, created_at, updated_at FROM users WHERE deleted_at = 0")
	if err != nil {
		return &pb.GetUsersResponse{Success: false, Message: "Failed to get users"}, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		user := &pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return &pb.GetUsersResponse{Success: false, Message: "Failed to scan user"}, err
		}
		users = append(users, user)
	}

	return &pb.GetUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Users:   users,
	}, nil
}
func (u *UserRepo) GetUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.GetUsersResponse, error) {
	query := "SELECT id, username, password, role, created_at, updated_at FROM users WHERE deleted_at = 0"
	args := []interface{}{}

	if req.Username != "" {
		query += " AND username LIKE $1"
		args = append(args, "%"+req.Username+"%")
	}

	if req.Role != "" {
		query += " AND role = $" + fmt.Sprint(len(args)+1)
		args = append(args, req.Role)
	}

	rows, err := u.db.Query(query, args...)
	if err != nil {
		return &pb.GetUsersResponse{Success: false, Message: "Failed to get users"}, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		user := &pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return &pb.GetUsersResponse{Success: false, Message: "Failed to scan user"}, err
		}
		users = append(users, user)
	}

	return &pb.GetUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Users:   users,
	}, nil
}
