package review

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"

	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
	entreview "github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.ReviewSelect
	infos []*npool.Review
	total uint32
	IDs   []uuid.UUID
}

func (h *queryHandler) selectReviews(stm *ent.ReviewQuery) {
	h.stm = stm.Select(
		entreview.FieldID,
		entreview.FieldAppID,
		entreview.FieldDomain,
		entreview.FieldReviewerID,
		entreview.FieldObjectID,
		entreview.FieldTrigger,
		entreview.FieldObjectType,
		entreview.FieldState,
		entreview.FieldMessage,
		entreview.FieldCreatedAt,
		entreview.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryReview(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	h.selectReviews(
		cli.
			Review.
			Query().
			Where(
				entreview.ID(*h.ID),
			),
	)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Trigger = npool.ReviewTriggerType(npool.ReviewTriggerType_value[info.TriggerStr])
		info.ObjectType = npool.ReviewObjectType(npool.ReviewObjectType_value[info.ObjectTypeStr])
		info.State = npool.ReviewState(npool.ReviewState_value[info.StateStr])
	}
}

func (h *Handler) GetReview(ctx context.Context) (info *npool.Review, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReview(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *queryHandler) queryCondsByBonds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.Review.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectReviews(stm)
	return nil
}

func (h *Handler) GetReviews(ctx context.Context) ([]*npool.Review, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCondsByBonds(_ctx, cli); err != nil {
			return err
		}
		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
