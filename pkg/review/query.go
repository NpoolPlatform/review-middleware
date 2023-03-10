package review

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"

	converter "github.com/NpoolPlatform/review-manager/pkg/converter/review"
	"github.com/NpoolPlatform/review-manager/pkg/db"
	"github.com/NpoolPlatform/review-manager/pkg/db/ent"
	entreview "github.com/NpoolPlatform/review-manager/pkg/db/ent/review"

	"github.com/google/uuid"
)

func GetObjectReview(ctx context.Context, appID, domain, objectID string, objectType mgrpb.ReviewObjectType) (*mgrpb.Review, error) {
	var info *ent.Review

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Review.
			Query().
			Where(
				entreview.AppID(uuid.MustParse(appID)),
				entreview.Domain(domain),
				entreview.ObjectID(uuid.MustParse(objectID)),
				entreview.ObjectType(objectType.String()),
			).
			Order(ent.Desc(entreview.FieldUpdatedAt)).
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

func GetObjectReviews(
	ctx context.Context,
	appID, domain string,
	objectIDs []string,
	objectType mgrpb.ReviewObjectType,
) (
	[]*mgrpb.Review, error,
) {
	var infos []*mgrpb.Review

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, id := range objectIDs {
			_infos, err := cli.
				Review.
				Query().
				Where(
					entreview.AppID(uuid.MustParse(appID)),
					entreview.Domain(domain),
					entreview.ObjectID(uuid.MustParse(id)),
					entreview.ObjectType(objectType.String()),
				).
				Order(ent.Desc(entreview.FieldUpdatedAt)).
				Limit(1).
				All(_ctx)
			if err != nil {
				return err
			}

			if len(_infos) == 0 {
				continue
			}

			infos = append(infos, converter.Ent2Grpc(_infos[0]))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
