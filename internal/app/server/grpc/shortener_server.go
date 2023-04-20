package grpc

import (
	"context"
	"github.com/casnerano/go-url-shortener/internal/app/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ShortenerServer struct {
	proto.UnimplementedShortenerServer
}

func (s *ShortenerServer) Get(ctx context.Context, in *proto.GetShortURLRequest) (*proto.GetShortURLResponse, error) {
	response := proto.GetShortURLResponse{
		Result: "http://ya.ru",
	}
	return &response, nil
}

func (s *ShortenerServer) GetUserHistory(context.Context, *emptypb.Empty) (*proto.GetUserHistoryShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserHistory not implemented")
}

func (s *ShortenerServer) GetStats(context.Context, *emptypb.Empty) (*proto.GetStatsShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}

func (s *ShortenerServer) CreateURL(context.Context, *proto.CreateShortURLRequest) (*proto.CreateShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateURL not implemented")
}

func (s *ShortenerServer) CreateBatch(context.Context, *proto.CreateBatchRequest) (*proto.CreateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatch not implemented")
}

func (s *ShortenerServer) DeleteBatch(context.Context, *proto.DeleteBatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBatch not implemented")
}
