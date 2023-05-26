package review

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
)

func CreateReview(ctx context.Context, in *npool.ReviewReq) (*npool.Review, error) {
	handler, err := review1.NewHandler(ctx,
		review1.WithAppID(in.AppID),
		review1.WithReviewerID(in.ReviewerID),
		review1.WithDomain(in.Domain),
		review1.WithObjectID(in.ObjectID),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateReview(ctx)
}
