package review

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/mw/review"
)

func Prepare(body string) (interface{}, error) {
	req := npool.ReviewReq{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Apply(ctx context.Context, req interface{}) error {
	_req, ok := req.(*npool.ReviewReq)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	handler, err := review1.NewHandler(ctx,
		review1.WithAppID(_req.AppID),
		review1.WithDomain(_req.Domain),
		review1.WithReviewerID(_req.ReviewerID),
		review1.WithObjectID(_req.ObjectID),
		review1.WithObjectType(_req.ObjectType),
		review1.WithTrigger(_req.Trigger),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Apply",
			"Req", req,
			"Error", err,
		)
		return err
	}

	if _, err := handler.CreateReview(ctx); err != nil {
		return err
	}

	return nil
}
