package review

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mw/v2"

	// "github.com/NpoolPlatform/review-manager/pkg/db"
	// "github.com/NpoolPlatform/review-manager/pkg/db/ent"
	// entreview "github.com/NpoolPlatform/review-manager/pkg/db/ent/review"

	// converter "github.com/NpoolPlatform/review-manager/pkg/converter/review"
	// crud "github.com/NpoolPlatform/review-manager/pkg/crud/review"

	"github.com/google/uuid"
)

func CreateReview(ctx context.Context, in *mgrpb.ReviewReq) (*mgrpb.Review, error) {
	var info *ent.Review

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		exist, err := tx.
			Review.
			Query().
			Where(
				entreview.AppID(uuid.MustParse(in.GetAppID())),
				entreview.Domain(in.GetDomain()),
				entreview.ObjectID(uuid.MustParse(in.GetObjectID())),
				entreview.ObjectType(in.GetObjectType().String()),
				entreview.Or(
					entreview.State(mgrpb.ReviewState_Wait.String()),
					entreview.State(mgrpb.ReviewState_Approved.String()),
				),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("object review exist")
		}

		info, err = crud.CreateSet(tx.Review.Create(), in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return converter.Ent2Grpc(info), nil
}
