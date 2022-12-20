package api

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"

	mgrapi "github.com/NpoolPlatform/review-manager/api"
	mgrcli "github.com/NpoolPlatform/review-manager/pkg/client/review"
	constant "github.com/NpoolPlatform/review-middleware/pkg/const"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/review"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CreateReview(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	if err := mgrapi.ValidateCreate(in.GetInfo()); err != nil {
		return &npool.CreateReviewResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := review1.CreateReview(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateReviewResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateReview(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	info, err := mgrcli.UpdateReview(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateReviewResponse{
		Info: info,
	}, nil
}

func (s *Server) GetObjectReview(ctx context.Context, in *npool.GetObjectReviewRequest) (*npool.GetObjectReviewResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	if _, err := uuid.Parse(in.GetObjectID()); err != nil {
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	if in.GetDomain() == "" {
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, "Domain is invalid")
	}
	switch in.GetObjectType() {
	case mgrpb.ReviewObjectType_ObjectKyc:
	case mgrpb.ReviewObjectType_ObjectWithdrawal:
	default:
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, "Object is invalid")
	}

	info, err := review1.GetObjectReview(ctx, in.GetAppID(), in.GetDomain(), in.GetObjectID(), in.GetObjectType())
	if err != nil {
		return &npool.GetObjectReviewResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetObjectReviewResponse{
		Info: info,
	}, nil
}

func (s *Server) GetReviews(ctx context.Context, in *npool.GetReviewsRequest) (*npool.GetReviewsResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := mgrcli.GetReviews(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReviewsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetObjectReviews(ctx context.Context, in *npool.GetObjectReviewsRequest) (*npool.GetObjectReviewsResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}
	if in.GetDomain() == "" {
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, "Domain is invalid")
	}
	for _, id := range in.GetObjectIDs() {
		if _, err := uuid.Parse(id); err != nil {
			return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	switch in.GetObjectType() {
	case mgrpb.ReviewObjectType_ObjectKyc:
	case mgrpb.ReviewObjectType_ObjectWithdrawal:
	default:
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, "ObjectType is invalid")
	}

	infos, err := review1.GetObjectReviews(ctx, in.GetAppID(), in.GetDomain(), in.GetObjectIDs(), in.GetObjectType())
	if err != nil {
		return &npool.GetObjectReviewsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetObjectReviewsResponse{
		Infos: infos,
	}, nil
}
