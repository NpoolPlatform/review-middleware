package review

import (
	"context"

	review1 "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	review1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	review1.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return review1.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
