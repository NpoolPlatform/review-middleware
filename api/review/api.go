package review

import (
	review1 "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"google.golang.org/grpc"
)

type Server struct {
	review1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	review1.RegisterMiddlewareServer(server, &Server{})
}
