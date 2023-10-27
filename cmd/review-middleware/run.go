package main

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/action"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/review-middleware/api"
	"github.com/NpoolPlatform/review-middleware/pkg/db"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/review-middleware/pkg/migrator"

	"github.com/NpoolPlatform/review-middleware/pkg/pubsub"
	cli "github.com/urfave/cli/v2"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		return action.Run(
			c.Context,
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
		)
	},
}

func run(ctx context.Context) error {
	if err := migrator.Migrate(ctx); err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}
	return pubsub.Subscribe(ctx)
}

func shutdown(ctx context.Context) {
	<-ctx.Done()
	logger.Sugar().Infow(
		"Watch",
		"State", "Done",
		"Error", ctx.Err(),
	)
	_ = pubsub.Shutdown(ctx)
}

func watch(ctx context.Context, cancel context.CancelFunc) error {
	go shutdown(ctx)
	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apicli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	_ = apicli.Register(mux)
	return nil
}
