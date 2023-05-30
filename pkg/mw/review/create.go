package review

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"

	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.ObjectID == nil {
		return fmt.Errorf("invalid object id")
	}
	if h.Domain == nil {
		return fmt.Errorf("invalid domain")
	}
	if h.ReviewerID == nil {
		return fmt.Errorf("invalid reviewer id")
	}
	return nil
}

func (h *Handler) CreateReview(ctx context.Context) (info *npool.Review, err error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
		return nil, err
	}

	return h.GetReview(ctx)
}
