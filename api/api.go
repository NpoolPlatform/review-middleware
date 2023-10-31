package api

import (
	"context"

	v2 "github.com/NpoolPlatform/message/npool/review/mw/v2"
	"github.com/NpoolPlatform/review-middleware/api/review"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v2.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	v2.RegisterMiddlewareServer(server, &Server{})
	review.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v2.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := review.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
