package review

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"
	mgrcli "github.com/NpoolPlatform/review-manager/pkg/client/review"
)

func CreateReview(ctx context.Context, in *mgrpb.ReviewReq) (*mgrpb.Review, error) {
	// TODO: if we have one wait or approved for object, do not create again

	info, err := mgrcli.CreateReview(ctx, in)
	if err != nil {
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	return info, nil
}
