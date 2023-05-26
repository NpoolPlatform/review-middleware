package review

import (
	"context"
	"encoding/json"
	"fmt"

	reviewmgrpb "github.com/NpoolPlatform/message/npool/review/mw/v2"
	review1 "github.com/NpoolPlatform/review-middleware/pkg/review"
)

func Prepare(body string) (interface{}, error) {
	req := reviewmgrpb.ReviewReq{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Apply(ctx context.Context, req interface{}) error {
	_req, ok := req.(*reviewmgrpb.ReviewReq)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	if _, err := review1.CreateReview(ctx, _req); err != nil {
		return err
	}

	return nil
}
