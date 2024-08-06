package token

import (
	"time"
	pb "github.com/Bekzodbekk/protofiles/genproto/user"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "HelloWorld"

// CreateTokens function creates both access and refresh tokens.
func CreateTokens(user *pb.User) (string, string, error) {
	// Create access token
	accessTokenClaims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Access token expires in 1 hour
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	// Create refresh token
	refreshTokenClaims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Refresh token expires in 24 hours
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
