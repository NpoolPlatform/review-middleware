package review

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mw/v2"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"

	"github.com/NpoolPlatform/review-middleware/pkg/db"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
	entreview "github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"
	converter "github.com/NpoolPlatform/review-middleware/pkg/mw/converter/review"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm       *ent.ReviewSelect
	infos     []*npool.Review
	total     uint32
	ObjectIDs []uuid.UUID
}

func (h *queryHandler) selectObjectReviews(stm *ent.ReviewQuery) {
	h.stm = stm.Select(
		entreview.FieldID,
		entreview.FieldAppID,
		entreview.FieldDomain,
		entreview.FieldReviewerID,
		entreview.FieldObjectID,
		entreview.FieldTrigger,
		entreview.FieldObjectType,
		entreview.FieldState,
		entreview.FieldCreatedAt,
		entreview.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryObjectReview(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	h.selectObjectReviews(
		cli.
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
			Limit(1),
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

func (h *Handler) GetObjectReview(ctx context.Context, appID, domain, objectID string, objectType mgrpb.ReviewObjectType) (info *npool.Review, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryObjectReview(cli); err != nil {
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

func (h *Handler) GetObjectReviews(ctx context.Context) ([]*mgrpb.Review, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		// for _, objectID := range handler.ObjectIDs {
		// 	h.ObjectID = &objectID
		// 	if err := handler.queryObjectReview(cli); err != nil {
		// 		return err
		// 	}
		// }
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
