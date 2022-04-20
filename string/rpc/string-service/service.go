package string_service

import (
	"context"
	"github.com/FFFcomewhere/University-course-selection/string/rpc/pb"
)

// Service constants
const (
	StrMaxSize = 1024
)

type StringService struct {
	*pb.UnimplementedStringServiceServer
}

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := pb.StringResponse{Ret: req.A + req.B}
	return &response, nil
}

func (s *StringService) mustEmbedUnimplementedStringServiceServer() {
	panic("implement me")
}
