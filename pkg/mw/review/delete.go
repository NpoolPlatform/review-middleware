package review

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	reviewcrud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
)

func (h *Handler) DeleteReview(ctx context.Context) (info *npool.Review, err error) {
	info, err = h.GetReview(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := reviewcrud.UpdateSet(
			cli.Review.UpdateOneID(*h.ID),
			&reviewcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
