// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pubsubmessageMixin := schema.PubsubMessage{}.Mixin()
	pubsubmessage.Policy = privacy.NewPolicies(pubsubmessageMixin[0], schema.PubsubMessage{})
	pubsubmessage.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := pubsubmessage.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	pubsubmessageMixinFields0 := pubsubmessageMixin[0].Fields()
	_ = pubsubmessageMixinFields0
	pubsubmessageFields := schema.PubsubMessage{}.Fields()
	_ = pubsubmessageFields
	// pubsubmessageDescCreatedAt is the schema descriptor for created_at field.
	pubsubmessageDescCreatedAt := pubsubmessageMixinFields0[0].Descriptor()
	// pubsubmessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	pubsubmessage.DefaultCreatedAt = pubsubmessageDescCreatedAt.Default.(func() uint32)
	// pubsubmessageDescUpdatedAt is the schema descriptor for updated_at field.
	pubsubmessageDescUpdatedAt := pubsubmessageMixinFields0[1].Descriptor()
	// pubsubmessage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pubsubmessage.DefaultUpdatedAt = pubsubmessageDescUpdatedAt.Default.(func() uint32)
	// pubsubmessage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pubsubmessage.UpdateDefaultUpdatedAt = pubsubmessageDescUpdatedAt.UpdateDefault.(func() uint32)
	// pubsubmessageDescDeletedAt is the schema descriptor for deleted_at field.
	pubsubmessageDescDeletedAt := pubsubmessageMixinFields0[2].Descriptor()
	// pubsubmessage.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	pubsubmessage.DefaultDeletedAt = pubsubmessageDescDeletedAt.Default.(func() uint32)
	// pubsubmessageDescMessageID is the schema descriptor for message_id field.
	pubsubmessageDescMessageID := pubsubmessageFields[1].Descriptor()
	// pubsubmessage.DefaultMessageID holds the default value on creation for the message_id field.
	pubsubmessage.DefaultMessageID = pubsubmessageDescMessageID.Default.(string)
	// pubsubmessageDescState is the schema descriptor for state field.
	pubsubmessageDescState := pubsubmessageFields[2].Descriptor()
	// pubsubmessage.DefaultState holds the default value on creation for the state field.
	pubsubmessage.DefaultState = pubsubmessageDescState.Default.(string)
	// pubsubmessageDescRespToID is the schema descriptor for resp_to_id field.
	pubsubmessageDescRespToID := pubsubmessageFields[3].Descriptor()
	// pubsubmessage.DefaultRespToID holds the default value on creation for the resp_to_id field.
	pubsubmessage.DefaultRespToID = pubsubmessageDescRespToID.Default.(func() uuid.UUID)
	// pubsubmessageDescUndoID is the schema descriptor for undo_id field.
	pubsubmessageDescUndoID := pubsubmessageFields[4].Descriptor()
	// pubsubmessage.DefaultUndoID holds the default value on creation for the undo_id field.
	pubsubmessage.DefaultUndoID = pubsubmessageDescUndoID.Default.(func() uuid.UUID)
	// pubsubmessageDescArguments is the schema descriptor for arguments field.
	pubsubmessageDescArguments := pubsubmessageFields[5].Descriptor()
	// pubsubmessage.DefaultArguments holds the default value on creation for the arguments field.
	pubsubmessage.DefaultArguments = pubsubmessageDescArguments.Default.(string)
	reviewMixin := schema.Review{}.Mixin()
	review.Policy = privacy.NewPolicies(reviewMixin[0], schema.Review{})
	review.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := review.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	reviewMixinFields0 := reviewMixin[0].Fields()
	_ = reviewMixinFields0
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescCreatedAt is the schema descriptor for created_at field.
	reviewDescCreatedAt := reviewMixinFields0[0].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the created_at field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() uint32)
	// reviewDescUpdatedAt is the schema descriptor for updated_at field.
	reviewDescUpdatedAt := reviewMixinFields0[1].Descriptor()
	// review.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	review.DefaultUpdatedAt = reviewDescUpdatedAt.Default.(func() uint32)
	// review.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	review.UpdateDefaultUpdatedAt = reviewDescUpdatedAt.UpdateDefault.(func() uint32)
	// reviewDescDeletedAt is the schema descriptor for deleted_at field.
	reviewDescDeletedAt := reviewMixinFields0[2].Descriptor()
	// review.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	review.DefaultDeletedAt = reviewDescDeletedAt.Default.(func() uint32)
	// reviewDescAppID is the schema descriptor for app_id field.
	reviewDescAppID := reviewFields[1].Descriptor()
	// review.DefaultAppID holds the default value on creation for the app_id field.
	review.DefaultAppID = reviewDescAppID.Default.(func() uuid.UUID)
	// reviewDescReviewerID is the schema descriptor for reviewer_id field.
	reviewDescReviewerID := reviewFields[2].Descriptor()
	// review.DefaultReviewerID holds the default value on creation for the reviewer_id field.
	review.DefaultReviewerID = reviewDescReviewerID.Default.(func() uuid.UUID)
	// reviewDescDomain is the schema descriptor for domain field.
	reviewDescDomain := reviewFields[3].Descriptor()
	// review.DefaultDomain holds the default value on creation for the domain field.
	review.DefaultDomain = reviewDescDomain.Default.(string)
	// reviewDescObjectID is the schema descriptor for object_id field.
	reviewDescObjectID := reviewFields[4].Descriptor()
	// review.DefaultObjectID holds the default value on creation for the object_id field.
	review.DefaultObjectID = reviewDescObjectID.Default.(func() uuid.UUID)
	// reviewDescTrigger is the schema descriptor for trigger field.
	reviewDescTrigger := reviewFields[5].Descriptor()
	// review.DefaultTrigger holds the default value on creation for the trigger field.
	review.DefaultTrigger = reviewDescTrigger.Default.(string)
	// reviewDescObjectType is the schema descriptor for object_type field.
	reviewDescObjectType := reviewFields[6].Descriptor()
	// review.DefaultObjectType holds the default value on creation for the object_type field.
	review.DefaultObjectType = reviewDescObjectType.Default.(string)
	// reviewDescState is the schema descriptor for state field.
	reviewDescState := reviewFields[7].Descriptor()
	// review.DefaultState holds the default value on creation for the state field.
	review.DefaultState = reviewDescState.Default.(string)
	// reviewDescMessage is the schema descriptor for message field.
	reviewDescMessage := reviewFields[8].Descriptor()
	// review.DefaultMessage holds the default value on creation for the message field.
	review.DefaultMessage = reviewDescMessage.Default.(string)
	// reviewDescID is the schema descriptor for id field.
	reviewDescID := reviewFields[0].Descriptor()
	// review.DefaultID holds the default value on creation for the id field.
	review.DefaultID = reviewDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)
