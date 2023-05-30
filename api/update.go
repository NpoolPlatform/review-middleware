package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateReview(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	req := in.GetInfo()
	handler, err := review1.NewHandler(ctx,
		review1.WithID(req.ID),
		review1.WithReviewerID(req.ReviewerID),
		review1.WithState(req.State, req.Message),
		review1.WithMessage(req.Message),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateReview",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateReview",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateReviewResponse{
		Info: info,
	}, nil
}