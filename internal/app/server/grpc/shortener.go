package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/casnerano/go-url-shortener/internal/app/server/grpc/proto"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerServer
}

func (s *ShortenerServer) Get(ctx context.Context, in *pb.GetShortURLRequest) (*pb.GetShortURLResponse, error) {
	response := pb.GetShortURLResponse{
		Result: "http://ya.ru",
	}
	return &response, nil
}
func (s *ShortenerServer) GetUserHistory(context.Context, *emptypb.Empty) (*pb.GetUserHistoryShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserHistory not implemented")
}
func (s *ShortenerServer) GetStats(context.Context, *emptypb.Empty) (*pb.GetStatsShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (s *ShortenerServer) CreateURL(context.Context, *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateURL not implemented")
}
func (s *ShortenerServer) CreateBatch(context.Context, *pb.CreateBatchRequest) (*pb.CreateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatch not implemented")
}
func (s *ShortenerServer) DeleteBatch(context.Context, *pb.DeleteBatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBatch not implemented")
}
