package repository

import pb "github.com/Bekzodbekk/protofiles/genproto/medals"

type MedalRepository interface {
	CreateMedal(req *pb.CreateMedalRequest) (*pb.CreateMedalResponse, error)
	UpdateMedal(req *pb.UpdateMedalRequest) (*pb.UpdateMedalResponse, error)
	DeleteMedal(req *pb.DeleteMedalRequest) (*pb.DeleteMedalResponse, error)
	GetMedalById(req *pb.GetMedalByIdRequest) (*pb.GetMedalByIdResponse, error)
	GetMedals(req *pb.VoidMedal) (*pb.GetMedalsResponse, error)
	GetMedalByFilter(req *pb.GetMedalByFilterRequest) (*pb.GetMedalByFilterResponse, error)
}
