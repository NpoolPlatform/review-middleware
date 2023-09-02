package review

import (
	"context"
	"fmt"

	appcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	constant "github.com/NpoolPlatform/review-middleware/pkg/const"
	crud "github.com/NpoolPlatform/review-middleware/pkg/crud/review"
	"github.com/google/uuid"
)

type Handler struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	ReviewerID *uuid.UUID
	Domain     *string
	ObjectID   *uuid.UUID
	ObjectIDs  []uuid.UUID
	Trigger    *types.ReviewTriggerType
	ObjectType *types.ReviewObjectType
	State      *types.ReviewState
	Message    *string
	Offset     int32
	Limit      int32
	Conds      *crud.Conds
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}
		exist, err := appcli.ExistApp(ctx, *appID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}

		h.AppID = &_appID
		return nil
	}
}

func WithReviewerID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ReviewerID = &_id
		return nil
	}
}

func WithObjectID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ObjectID = &_id
		return nil
	}
}

func WithObjectIDs(ids []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range ids {
			uid, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			h.ObjectIDs = append(h.ObjectIDs, uid)
		}
		return nil
	}
}

func WithDomain(domain *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if domain == nil {
			return fmt.Errorf("invalid domain")
		}
		h.Domain = domain
		return nil
	}
}

func WithTrigger(trigger *types.ReviewTriggerType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if trigger == nil {
			return nil
		}
		switch *trigger {
		case types.ReviewTriggerType_InsufficientFunds:
		case types.ReviewTriggerType_InsufficientGas:
		case types.ReviewTriggerType_InsufficientFundsGas:
		case types.ReviewTriggerType_LargeAmount:
		case types.ReviewTriggerType_AutoReviewed:
		default:
			return fmt.Errorf("invalid trigger type")
		}

		h.Trigger = trigger
		return nil
	}
}

func WithObjectType(_type *types.ReviewObjectType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *_type {
		case types.ReviewObjectType_ObjectKyc:
		case types.ReviewObjectType_ObjectWithdrawal:
		default:
			return fmt.Errorf("invalid object type")
		}
		h.ObjectType = _type
		return nil
	}
}

func WithState(state *types.ReviewState, message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case types.ReviewState_Rejected:
		case types.ReviewState_Approved:
		default:
			return fmt.Errorf("invalid review state")
		}
		if *state == types.ReviewState_Rejected && message == nil {
			return fmt.Errorf("message is empty")
		}
		h.State = state
		return nil
	}
}

func WithMessage(message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			return nil
		}
		h.Message = message
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.ID.Op, Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.AppID.Op, Val: id}
		}
		if conds.ReviewerID != nil {
			id, err := uuid.Parse(conds.GetReviewerID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ReviewerID = &cruder.Cond{Op: conds.ReviewerID.Op, Val: id}
		}
		if conds.ObjectID != nil {
			id, err := uuid.Parse(conds.GetObjectID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ObjectID = &cruder.Cond{Op: conds.ObjectID.Op, Val: id}
		}
		if conds.Domain != nil {
			h.Conds.Domain = &cruder.Cond{Op: conds.Domain.Op, Val: conds.Domain.GetValue()}
		}
		if conds.Trigger != nil {
			trigger := conds.GetTrigger().GetValue()
			h.Conds.Trigger = &cruder.Cond{Op: conds.Trigger.Op, Val: types.ReviewTriggerType(trigger)}
		}
		if conds.ObjectType != nil {
			_type := conds.GetObjectType().GetValue()
			h.Conds.ObjectType = &cruder.Cond{Op: conds.ObjectType.Op, Val: types.ReviewObjectType(_type)}
		}
		if conds.State != nil {
			_state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{Op: conds.State.Op, Val: types.ReviewState(_state)}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
