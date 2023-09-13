package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	object1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetObjectReview(ctx context.Context, in *npool.GetObjectReviewRequest) (*npool.GetObjectReviewResponse, error) {
	handler, err := object1.NewHandler(
		ctx,
		object1.WithAppID(&in.AppID, true),
		object1.WithObjectID(&in.ObjectID, true),
		object1.WithDomain(&in.Domain, true),
		object1.WithObjectType(&in.ObjectType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetObjectReview",
			"In", in,
			"Error", err,
		)
		return &npool.GetObjectReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetObjectReview(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetObjectReview",
			"In", in,
			"Error", err,
		)
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetObjectReviewResponse{
		Info: info,
	}, nil
}

func (s *Server) GetObjectReviews(ctx context.Context, in *npool.GetObjectReviewsRequest) (*npool.GetObjectReviewsResponse, error) {
	handler, err := object1.NewHandler(
		ctx,
		object1.WithAppID(&in.AppID, true),
		object1.WithObjectIDs(in.GetObjectIDs(), true),
		object1.WithDomain(&in.Domain, true),
		object1.WithObjectType(&in.ObjectType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetObjectReviews",
			"In", in,
			"Error", err,
		)
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.GetObjectReviews(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetObjectReviews",
			"In", in,
			"Error", err,
		)
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetObjectReviewsResponse{
		Infos: infos,
	}, nil
}
