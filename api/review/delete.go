package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteReview(ctx context.Context, in *npool.DeleteReviewRequest) (*npool.DeleteReviewResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteReview",
			"In", in,
		)
		return &npool.DeleteReviewResponse{}, status.Error(codes.InvalidArgument, "invalid argument")
	}
	handler, err := review1.NewHandler(ctx,
		review1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReview",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReview",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteReviewResponse{
		Info: info,
	}, nil
}
