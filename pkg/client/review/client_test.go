package review

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	reviewtypes "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	reviewmwpb "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"github.com/NpoolPlatform/review-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = reviewmwpb.Review{
		EntID:         uuid.NewString(),
		AppID:         uuid.NewString(),
		Domain:        uuid.NewString(),
		ObjectType:    types.ReviewObjectType_ObjectKyc,
		ObjectTypeStr: types.ReviewObjectType_ObjectKyc.String(),
		ObjectID:      uuid.NewString(),
		ReviewerID:    uuid.Nil.String(),
		State:         types.ReviewState_Wait,
		StateStr:      types.ReviewState_Wait.String(),
		Trigger:       types.ReviewTriggerType_DefaultTriggerType,
		TriggerStr:    types.ReviewTriggerType_DefaultTriggerType.String(),
		Message:       "",
	}
)

func createReview(t *testing.T) {
	info, err := CreateReview(context.Background(), &reviewmwpb.ReviewReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		Domain:     &ret.Domain,
		ObjectID:   &ret.ObjectID,
		ObjectType: &ret.ObjectType,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateReview(t *testing.T) {
	ret.State = types.ReviewState_Rejected
	ret.StateStr = types.ReviewState_Rejected.String()
	ret.ReviewerID = uuid.NewString()
	ret.Message = uuid.NewString()

	info, err := UpdateReview(context.Background(), &reviewmwpb.ReviewReq{
		ID:         &ret.ID,
		State:      &ret.State,
		ReviewerID: &ret.ReviewerID,
		Message:    &ret.Message,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getReview(t *testing.T) {
	info, err := GetReview(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

var (
	conds = &reviewmwpb.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		ObjectID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.ObjectID},
		Domain:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Domain},
		Trigger:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewTriggerType_DefaultTriggerType)},
		ObjectType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewObjectType_ObjectKyc)},
		State:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(reviewtypes.ReviewState_Rejected)},
	}
)

func getReviews(t *testing.T) {
	infos, total, err := GetReviews(context.Background(), conds, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, &ret, infos[0])
	}
}

func existReviewConds(t *testing.T) {
	exist, err := ExistReviewConds(context.Background(), &reviewmwpb.ExistReviewCondsRequest{
		Conds: conds,
	})
	if assert.Nil(t, err) {
		assert.True(t, exist)
	}
}

func deleteReview(t *testing.T) {
	_, err := DeleteReview(context.Background(), ret.ID)
	assert.Nil(t, err)

	info, err := GetReview(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createReview", createReview)
	t.Run("updateReview", updateReview)
	t.Run("getReviews", getReviews)
	t.Run("getReview", getReview)
	t.Run("existReviewConds", existReviewConds)
	t.Run("deleteReview", deleteReview)
}
