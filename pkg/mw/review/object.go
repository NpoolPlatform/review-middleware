package review

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
	entreview "github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"
	converter "github.com/NpoolPlatform/review-middleware/pkg/mw/converter/review"
)

type queryObjectHandler struct {
	*Handler
	infos     []*npool.Review
	ObjectIDs []uuid.UUID
}

func (h *queryObjectHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.ObjectID == nil {
		return fmt.Errorf("invalid object id")
	}
	if h.Domain == nil {
		return fmt.Errorf("invalid domain")
	}
	if h.ObjectType == nil {
		return fmt.Errorf("invalid object type")
	}
	return nil
}

func (h *Handler) GetObjectReview(ctx context.Context) (*npool.Review, error) {
	handler := &queryObjectHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
		return nil, err
	}

	var info *ent.Review
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Review.
			Query().
			Where(
				entreview.AppID(*h.AppID),
				entreview.Domain(*h.Domain),
				entreview.ObjectID(*h.ObjectID),
				entreview.ObjectType(h.ObjectType.String()),
			).
			Order(
				ent.Desc(entreview.FieldUpdatedAt),
			).
			Limit(1).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(infos) > 0 {
			info = infos[0]
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return converter.Ent2Grpc(info), nil
}

func (h *Handler) GetObjectReviews(ctx context.Context) ([]*npool.Review, error) {
	handler := &queryObjectHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Review.
			Query().
			Where(
				entreview.AppID(*h.AppID),
				entreview.Domain(*h.Domain),
				entreview.ObjectIDIn(handler.ObjectIDs...),
				entreview.ObjectType(h.ObjectType.String()),
			).
			Order(ent.Desc(entreview.FieldUpdatedAt)).
			All(_ctx)
		if err != nil {
			return err
		}

		handler.infos = converter.Ent2GrpcMany(infos)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return handler.infos, nil
}
