package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReview(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateReview",
			"In", in,
		)
		return &npool.CreateReviewResponse{}, status.Error(codes.InvalidArgument, "invalid argument")
	}
	handler, err := review1.NewHandler(ctx,
		review1.WithID(req.ID, true),
		review1.WithAppID(req.AppID, true),
		review1.WithDomain(req.Domain, true),
		review1.WithReviewerID(req.ReviewerID, false),
		review1.WithObjectID(req.ObjectID, true),
		review1.WithObjectType(req.ObjectType, true),
		review1.WithTrigger(req.Trigger, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReview",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReview",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateReviewResponse{
		Info: info,
	}, nil
}
