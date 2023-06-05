package review

import (
	"github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"google.golang.org/grpc"
)

type Server struct {
	review.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	review.RegisterMiddlewareServer(server, &Server{})
}
