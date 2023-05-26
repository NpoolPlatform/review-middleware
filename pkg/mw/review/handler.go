package review

import (
	"context"
	"fmt"

	appcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	constant "github.com/NpoolPlatform/review-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	ReviewerID *uuid.UUID
	Domain     *string
	ObjectID   *uuid.UUID
	Trigger    *npool.ReviewTriggerType
	ObjectType *npool.ReviewObjectType
	State      *npool.ReviewState
	Message    *string
	Offset     int32
	Limit      int32
	Conds      npool.Conds
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
		if appID == nil {
			return nil
		}
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

func WithDomain(domain *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if domain == nil {
			return fmt.Errorf("invalid domain")
		}
		h.Domain = domain
		return nil
	}
}

func WithTrigger(trigger *npool.ReviewTriggerType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *trigger {
		case npool.ReviewTriggerType_InsufficientFunds:
		case npool.ReviewTriggerType_InsufficientGas:
		case npool.ReviewTriggerType_InsufficientFundsGas:
		case npool.ReviewTriggerType_LargeAmount:
		case npool.ReviewTriggerType_AutoReviewed:
		default:
			return fmt.Errorf("invalid trigger type")
		}

		h.Trigger = trigger
		return nil
	}
}

func WithObjectType(_type *npool.ReviewObjectType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *_type {
		case npool.ReviewObjectType_ObjectKyc:
		case npool.ReviewObjectType_ObjectWithdrawal:
		default:
			return fmt.Errorf("invalid object type")
		}

		h.ObjectType = _type
		return nil
	}
}

func WithState(state *npool.ReviewState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *state {
		case npool.ReviewState_Rejected:
		case npool.ReviewState_Approved:
		default:
			return fmt.Errorf("invalid review state")
		}

		h.State = state
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
