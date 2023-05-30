package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReviews(ctx context.Context, in *npool.GetReviewsRequest) (*npool.GetReviewsResponse, error) {
	handler, err := review1.NewHandler(ctx,
		review1.WithConds(in.GetConds()),
		review1.WithOffset(in.Offset),
		review1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReviews",
			"Req", in,
			"Error", err,
		)
		return &npool.GetReviewsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetReviews(ctx)
	if err != nil {
		return &npool.GetReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReviewsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
