package review

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
)

func (h *Handler) UpdateReview(ctx context.Context) (info *npool.Review, err error) {
	info, err = h.GetReview(ctx)
	if err != nil {
		return nil, err
	}
	switch info.State {
	case npool.ReviewState_DefaultReviewState:
	case npool.ReviewState_Wait:
	default:
		return nil, fmt.Errorf("current state can not be update")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Review.UpdateOneID(*h.ID),
			&crud.Req{
				State:      h.State,
				ReviewerID: h.ReviewerID,
				Message:    h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetReview(ctx)
}
