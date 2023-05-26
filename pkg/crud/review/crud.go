package review

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"

	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	ReviewerID *uuid.UUID
	Domain     *string
	ObjectID   *uuid.UUID
	Trigger    *npool.ReviewTriggerType
	ObjectType *npool.ReviewObjectType
	State      *npool.ReviewState
	Message    *string
	DeletedAt  *uint32
}

func CreateSet(c *ent.ReviewCreate, in *Req) *ent.ReviewCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.ReviewerID != nil {
		c.SetReviewerID(*in.ReviewerID)
	}
	if in.Domain != nil {
		c.SetDomain(*in.Domain)
	}
	if in.ObjectID != nil {
		c.SetObjectID(*in.ObjectID)
	}
	if in.Trigger != nil {
		c.SetTrigger(in.Trigger.String())
	}
	if in.ObjectType != nil {
		c.SetObjectType(in.ObjectID.String())
	}
	c.SetState(npool.ReviewState_Wait.String())
	return c
}

func UpdateSet(u *ent.ReviewUpdateOne, in *Req) *ent.ReviewUpdateOne {
	if in.ReviewerID != nil {
		u.SetReviewerID(*in.ReviewerID)
	}

	// TODO: Only ReviewState_Wait can be update
	if in.State != nil {
		// switch info.State {
		// case npool.ReviewState_Wait.String():
		// default:
		// 	return nil, fmt.Errorf("permission denied")
		// }
		// stm = stm.SetState(in.GetState().String())
		u.SetState(in.State.String())
	}
	if in.Message != nil {
		u.SetMessage(*in.Message)
	}
	if in.DeletedAt != nil {
		u.SetDeletedAt(*in.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	ReviewerID *cruder.Cond
	Domain     *cruder.Cond
	ObjectID   *cruder.Cond
	Trigger    *cruder.Cond
	ObjectType *cruder.Cond
	State      *cruder.Cond
}

func SetQueryConds(q *ent.ReviewQuery, conds *Conds) (*ent.ReviewQuery, error) { //nolint
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(review.ID(id))
		default:
			return nil, fmt.Errorf("invalid id operator field %v", conds.ID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}

		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(review.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appid operator field %v", conds.AppID.Op)
		}
	}
	if conds.ReviewerID != nil {
		reviewerID, ok := conds.ReviewerID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid reviewer id")
		}

		switch conds.ReviewerID.Op {
		case cruder.EQ:
			q.Where(review.ReviewerID(reviewerID))
		default:
			return nil, fmt.Errorf("invalid reviewer id operator field %v", conds.ReviewerID.Op)
		}
	}

	if conds.Domain != nil {
		domain, ok := conds.Domain.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid domain")
		}
		switch conds.Domain.Op {
		case cruder.EQ:
			q.Where(review.Domain(domain))
		default:
			return nil, fmt.Errorf("invalid domain operator field %v", conds.Domain.Op)
		}
	}
	if conds.ObjectID != nil {
		objectID, ok := conds.ObjectID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid reviewer id")
		}
		switch conds.ObjectID.Op {
		case cruder.EQ:
			q.Where(review.ObjectID(objectID))
		default:
			return nil, fmt.Errorf("invalid object id operator field %v", conds.ObjectID.Op)
		}
	}
	if conds.Trigger != nil {
		trigger, ok := conds.Trigger.Val.(npool.ReviewTriggerType)
		if !ok {
			return nil, fmt.Errorf("invalid trigger")
		}
		switch conds.Trigger.Op {
		case cruder.EQ:
			q.Where(review.Trigger(trigger.String()))
		default:
			return nil, fmt.Errorf("invalid trigger operator field %v", conds.Trigger.Op)
		}
	}
	if conds.ObjectType != nil {
		objectType, ok := conds.ObjectType.Val.(npool.ReviewObjectType)
		if !ok {
			return nil, fmt.Errorf("invalid object type")
		}
		switch conds.ObjectType.Op {
		case cruder.EQ:
			q.Where(review.ObjectType(objectType.String()))
		default:
			return nil, fmt.Errorf("invalid object type operator field %v", conds.ObjectType.Op)
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(npool.ReviewState)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(review.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid state operator field %v", conds.State.Op)
		}
	}
	return q, nil
}
