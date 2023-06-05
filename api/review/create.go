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
	handler, err := review1.NewHandler(ctx,
		review1.WithID(req.ID),
		review1.WithAppID(req.AppID),
		review1.WithDomain(req.Domain),
		review1.WithReviewerID(req.ReviewerID),
		review1.WithObjectID(req.ObjectID),
		review1.WithObjectType(req.ObjectType),
		review1.WithTrigger(req.Trigger),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReview",
			"Req", req,
			"Error", err,
		)
		return &npool.CreateReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReview",
			"Req", req,
			"Error", err,
		)
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateReviewResponse{
		Info: info,
	}, nil
}
