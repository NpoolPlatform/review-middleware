package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	"github.com/NpoolPlatform/review-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

func (Review) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("reviewer_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("domain").
			Optional().
			Default(""),
		field.
			UUID("object_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("trigger").
			Optional().
			Default(types.ReviewTriggerType_DefaultTriggerType.String()),
		field.
			String("object_type").
			Optional().
			Default(types.ReviewObjectType_DefaultObjectType.String()),
		field.
			String("state").
			Optional().
			Default(types.ReviewState_DefaultReviewState.String()),
		field.
			String("message").
			Optional().
			Default(""),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}
