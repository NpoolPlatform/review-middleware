package review

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"
)

func GetObjectReview(ctx context.Context, appID, domain, objectID string, objectType mgrpb.ReviewObjectType) (*mgrpb.Review, error) {
	return nil, nil
}
