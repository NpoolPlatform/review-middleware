package api

import (
	reviewmw "github.com/NpoolPlatform/message/npool/review/mw/v2"
	"github.com/NpoolPlatform/review-middleware/api/review"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	reviewmw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	reviewmw.RegisterMiddlewareServer(server, &Server{})
	review.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
