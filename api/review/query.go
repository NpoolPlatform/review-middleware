package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
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
		logger.Sugar().Errorw(
			"GetReviews",
			"Req", in,
			"Error", err,
		)
		return &npool.GetReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReviewsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetReview(ctx context.Context, in *npool.GetReviewRequest) (*npool.GetReviewResponse, error) {
	handler, err := review1.NewHandler(ctx,
		review1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReview",
			"Req", in,
			"Error", err,
		)
		return &npool.GetReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReview",
			"Req", in,
			"Error", err,
		)
		return &npool.GetReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReviewResponse{
		Info: info,
	}, nil
}
