package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/casnerano/go-url-shortener/internal/app/service/crypter"
)

// ContextUserUUIDType for context keys.
type ContextUserUUIDType string

// Meta key for user uuid.
const MetaUserUUIDKey ContextUserUUIDType = "SEC_USER_UUID"

// UnaryServer interceptor for server
// used to forward the user uuid
func UnaryServer(key []byte) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		var (
			userUUID string
			err      error
		)

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if values := md.Get(string(MetaUserUUIDKey)); len(values) > 0 {
				userUUID, err = crypter.DecryptString(values[0], key)
				if err != nil {
					return nil, status.Error(codes.Internal, err.Error())
				}
			}
		}

		if userUUID == "" {
			gUUID, err := uuid.NewUUID()
			if err != nil {
				return nil, err
			}

			userUUID = gUUID.String()
		}

		ctx = context.WithValue(ctx, MetaUserUUIDKey, userUUID)
		return handler(ctx, req)
	}
}
