package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistReviewConds(ctx context.Context, in *npool.ExistReviewCondsRequest) (*npool.ExistReviewCondsResponse, error) {
	handler, err := review1.NewHandler(
		ctx,
		review1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReview",
			"In", in,
			"Error", err,
		)
		return &npool.ExistReviewCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistReviewConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReview",
			"In", in,
			"Error", err,
		)
		return &npool.ExistReviewCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistReviewCondsResponse{
		Info: info,
	}, nil
}
