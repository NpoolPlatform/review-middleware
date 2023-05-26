package review

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"

	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
	entreview "github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"

	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createReview(ctx context.Context) error {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.Review.Create(),
			&crud.Req{
				AppID:      h.AppID,
				ReviewerID: h.ReviewerID,
				Domain:     h.Domain,
				ObjectID:   h.ObjectID,
				ObjectType: h.ObjectType,
				Trigger:    h.Trigger,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (h *Handler) CreateReview(ctx context.Context) (*npool.Review, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		exist, err := tx.
			Review.
			Query().
			Where(
				entreview.AppID(*h.AppID),
				entreview.Domain(*h.Domain),
				entreview.ObjectID(*h.ObjectID),
				entreview.ObjectType(h.ObjectType.String()),
				entreview.Or(
					entreview.State(npool.ReviewState_Wait.String()),
					entreview.State(npool.ReviewState_Approved.String()),
				),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("object review exist")
		}

		handler := &createHandler{
			Handler: h,
		}
		if err := handler.createReview(ctx); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return h.GetReview(ctx)
}
