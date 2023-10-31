package review

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"

	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/google/uuid"
)

func (h *Handler) CreateReview(ctx context.Context) (info *npool.Review, err error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Review.Create(),
			&crud.Req{
				ID:         h.ID,
				EntID:      h.EntID,
				AppID:      h.AppID,
				ReviewerID: h.ReviewerID,
				Domain:     h.Domain,
				ObjectID:   h.ObjectID,
				ObjectType: h.ObjectType,
				Trigger:    h.Trigger,
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
