//nolint:dupl
package review

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"

	constant "github.com/NpoolPlatform/review-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get review connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateReview(ctx context.Context, in *mgrpb.ReviewReq) (*mgrpb.Review, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateReview(ctx, &npool.CreateReviewRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create review: %v", err)
	}
	return info.(*mgrpb.Review), nil
}

func UpdateReview(ctx context.Context, in *mgrpb.ReviewReq) (*mgrpb.Review, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateReview(ctx, &npool.UpdateReviewRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update review: %v", err)
	}
	return info.(*mgrpb.Review), nil
}

func GetObjectReview(ctx context.Context, appID, domain, objectID string, objectType mgrpb.ReviewObjectType) (*mgrpb.Review, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetObjectReview(ctx, &npool.GetObjectReviewRequest{
			AppID:      appID,
			Domain:     domain,
			ObjectID:   objectID,
			ObjectType: objectType,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get review: %v", err)
	}
	return info.(*mgrpb.Review), nil
}

func GetObjectReviews(
	ctx context.Context,
	appID, domain string,
	objectIDs []string,
	objectType mgrpb.ReviewObjectType,
) (
	[]*mgrpb.Review, error,
) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetObjectReviews(ctx, &npool.GetObjectReviewsRequest{
			AppID:      appID,
			Domain:     domain,
			ObjectIDs:  objectIDs,
			ObjectType: objectType,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get review: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get review: %v", err)
	}
	return infos.([]*mgrpb.Review), nil
}

func GetReviews(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Review, uint32, error) {
	var total uint32

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetReviews(ctx, &npool.GetReviewsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get reviews: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get reviews: %v", err)
	}
	return infos.([]*mgrpb.Review), total, nil
}
