package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/proto"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/server/grpc/interceptor"
	"github.com/casnerano/go-url-shortener/internal/app/service"
)

// ShortenerServer grpc server struct.
type ShortenerServer struct {
	proto.UnimplementedShortenerServer

	cfg        *config.Config
	urlService *service.URL
}

// NewShortenerServer constructor.
func NewShortenerServer(cfg *config.Config, urlService *service.URL) *ShortenerServer {
	return &ShortenerServer{cfg: cfg, urlService: urlService}
}

// Get method.
func (s *ShortenerServer) Get(ctx context.Context, in *proto.GetShortURLRequest) (*proto.GetShortURLResponse, error) {
	response := &proto.GetShortURLResponse{}

	if in.GetShortCode() == "" {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	shortURL, err := s.urlService.GetByCode(in.GetShortCode())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response.Result = shortURL.Original

	return response, nil
}

// GetUserHistory method.
func (s *ShortenerServer) GetUserHistory(ctx context.Context, in *emptypb.Empty) (*proto.GetUserHistoryShortURLResponse, error) {
	response := &proto.GetUserHistoryShortURLResponse{}

	ctxUserUUID := ctx.Value(interceptor.MetaUserUUIDKey)
	userUUID, ok := ctxUserUUID.(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	userURLHistory, err := s.urlService.FindByUserUUID(userUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, userURL := range userURLHistory {
		item := &proto.GetUserHistoryShortURLResponse_Item{
			ShortUrl:    s.buildAbsoluteShortURL(userURL.Code),
			OriginalUrl: userURL.Original,
		}
		response.Items = append(response.Items, item)
	}

	return response, nil
}

// GetStats method.
func (s *ShortenerServer) GetStats(ctx context.Context, in *emptypb.Empty) (*proto.GetStatsShortURLResponse, error) {
	response := proto.GetStatsShortURLResponse{}

	urlsCount, err := s.urlService.GetTotalURLCount()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	usersCount, err := s.urlService.GetTotalUserCount()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.Links = int64(urlsCount)
	response.Users = int64(usersCount)

	return &response, nil
}

// CreateURL method.
func (s *ShortenerServer) CreateURL(ctx context.Context, in *proto.CreateShortURLRequest) (*proto.CreateShortURLResponse, error) {
	response := proto.CreateShortURLResponse{}

	if in.GetUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctxUserUUID := ctx.Value(interceptor.MetaUserUUIDKey)
	userUUID, ok := ctxUserUUID.(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	shortURLModel, err := s.urlService.Create(in.GetUrl(), userUUID)
	if err != nil {
		if errors.Is(err, repository.ErrURLAlreadyExist) {
			shortURLModel, err = s.urlService.GetByUserUUIDAndOriginal(userUUID, in.GetUrl())
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

			response.Result = s.buildAbsoluteShortURL(shortURLModel.Code)
			return &response, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	response.Result = s.buildAbsoluteShortURL(shortURLModel.Code)
	return &response, nil
}

// CreateBatch method.
func (s *ShortenerServer) CreateBatch(ctx context.Context, in *proto.CreateBatchRequest) (*proto.CreateBatchResponse, error) {
	response := proto.CreateBatchResponse{}

	ctxUserUUID := ctx.Value(interceptor.MetaUserUUIDKey)
	userUUID, ok := ctxUserUUID.(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	var batchRequest []*model.ShortURLBatchRequest
	for _, item := range in.Items {
		batchRequest = append(batchRequest, &model.ShortURLBatchRequest{
			CorrelationID: item.CorrelationId,
			OriginalURL:   item.OriginalUrl,
		})
	}

	models, err := s.urlService.CreateBatch(batchRequest, userUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, item := range models {
		response.Items = append(response.Items, &proto.CreateBatchResponse_Item{
			CorrelationId: item.CorrelationID,
			ShortUrl:      s.buildAbsoluteShortURL(item.ShortURL),
		})
	}

	return &response, nil
}

// DeleteBatch method.
func (s *ShortenerServer) DeleteBatch(ctx context.Context, in *proto.DeleteBatchRequest) (*emptypb.Empty, error) {
	ctxUserUUID := ctx.Value(interceptor.MetaUserUUIDKey)
	userUUID, ok := ctxUserUUID.(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	go func() {
		_ = s.urlService.DeleteBatch(in.Items, userUUID)
	}()

	return nil, nil
}

func (s *ShortenerServer) buildAbsoluteShortURL(shortCode string) string {
	return fmt.Sprintf("%s/%s", s.cfg.Server.BaseURL, shortCode)
}
