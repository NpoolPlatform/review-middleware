//nolint:dupl
package review

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"

	servicename "github.com/NpoolPlatform/review-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get review connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateReview(ctx context.Context, in *npool.ReviewReq) (*npool.Review, error) {
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
	return info.(*npool.Review), nil
}

func UpdateReview(ctx context.Context, in *npool.ReviewReq) (*npool.Review, error) {
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
	return info.(*npool.Review), nil
}

func GetReviews(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Review, uint32, error) {
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
	return infos.([]*npool.Review), total, nil
}

func GetReview(ctx context.Context, id string) (*npool.Review, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetReview(ctx, &npool.GetReviewRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get review: %v", err)
	}
	return info.(*npool.Review), nil
}

func DeleteReview(ctx context.Context, id uint32) (*npool.Review, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteReview(ctx, &npool.DeleteReviewRequest{
			Info: &npool.ReviewReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete review: %v", err)
	}

	return info.(*npool.Review), nil
}

func ExistReviewConds(ctx context.Context, conds *npool.ExistReviewCondsRequest) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistReviewConds(ctx, conds)
		if err != nil {
			return nil, fmt.Errorf("fail exist review: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail exist review: %v", err)
	}

	return info.(bool), nil
}
