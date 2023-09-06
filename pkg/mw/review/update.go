package review

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) UpdateReview(ctx context.Context) (info *npool.Review, err error) {
	info, err = h.GetReview(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid review")
	}
	switch info.State {
	case types.ReviewState_DefaultReviewState:
	case types.ReviewState_Wait:
	default:
		return nil, fmt.Errorf("permission denied")
	}

	if info.ReviewerID != uuid.Nil.String() && h.ReviewerID != nil {
		return nil, fmt.Errorf("permission denied")
	}
	if h.State != nil && *h.State == types.ReviewState_Rejected && (h.Message == nil || *h.Message == "") {
		return nil, fmt.Errorf("message is empty")
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
