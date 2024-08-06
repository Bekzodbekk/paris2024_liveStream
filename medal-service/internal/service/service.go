package service

import (
	"context"
	pb "github.com/Bekzodbekk/protofiles/genproto/medals"
	
	"medal-service/internal/repository"
)

type MedalService struct {
	pb.UnimplementedMedalServiceServer
	medalRepo repository.MedalRepository
}

func NewMedalService(medal repository.MedalRepository) *MedalService {
	return &MedalService{
		medalRepo: medal,
	}
}

func (s *MedalService) CreateMedal(ctx context.Context, req *pb.CreateMedalRequest) (*pb.CreateMedalResponse, error) {
	return s.medalRepo.CreateMedal(req)
}

func (s *MedalService) UpdateMedal(ctx context.Context, req *pb.UpdateMedalRequest) (*pb.UpdateMedalResponse, error) {
	return s.medalRepo.UpdateMedal(req)
}

func (s *MedalService) DeleteMedal(ctx context.Context, req *pb.DeleteMedalRequest) (*pb.DeleteMedalResponse, error) {
	return s.medalRepo.DeleteMedal(req)
}

func (s *MedalService) GetMedalById(ctx context.Context, req *pb.GetMedalByIdRequest) (*pb.GetMedalByIdResponse, error) {
	return s.medalRepo.GetMedalById(req)
}

func (s *MedalService) GetMedals(ctx context.Context, req *pb.VoidMedal) (*pb.GetMedalsResponse, error) {
	return s.medalRepo.GetMedals(req)
}

func (s *MedalService) GetMedalByFilter(ctx context.Context, req *pb.GetMedalByFilterRequest) (*pb.GetMedalByFilterResponse, error) {
	return s.medalRepo.GetMedalByFilter(req)
}
