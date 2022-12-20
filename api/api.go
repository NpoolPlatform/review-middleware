package api

import (
	review "github.com/NpoolPlatform/message/npool/review/mw/v2"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	review.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	review.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
