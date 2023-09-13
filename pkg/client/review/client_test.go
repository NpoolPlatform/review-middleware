package review

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
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
	ret = &npool.Review{
		AppID:         uuid.NewString(),
		Domain:        uuid.NewString(),
		ObjectType:    types.ReviewObjectType_ObjectKyc,
		ObjectTypeStr: types.ReviewObjectType_ObjectKyc.String(),
		ObjectID:      uuid.NewString(),
		ReviewerID:    uuid.NewString(),
		State:         types.ReviewState_Wait,
		StateStr:      types.ReviewState_Wait.String(),
		Trigger:       types.ReviewTriggerType_InsufficientGas,
		TriggerStr:    types.ReviewTriggerType_InsufficientGas.String(),
		Message:       "",
	}
)

var (
	app = appuserpb.App{
		Name:                     uuid.NewString(),
		Logo:                     uuid.NewString(),
		CreatedBy:                uuid.NewString(),
		Description:              uuid.NewString(),
		RecaptchaMethod:          basetypes.RecaptchaMethod_GoogleRecaptchaV3,
		CreateInvitationCodeWhen: basetypes.CreateInvitationCodeWhen_Registration,
	}
)

func setupApp(t *testing.T) func(*testing.T) {
	info, err := appusercli.CreateApp(context.Background(), &appuserpb.AppReq{
		Name:                     &app.Name,
		Logo:                     &app.Logo,
		Description:              &app.Description,
		CreatedBy:                &app.CreatedBy,
		RecaptchaMethod:          &app.RecaptchaMethod,
		CreateInvitationCodeWhen: &app.CreateInvitationCodeWhen,
	})

	assert.Nil(t, err)
	assert.NotNil(t, info.ID)

	ret.AppID = info.ID
	return func(t *testing.T) {
		_, _ = appusercli.DeleteApp(context.Background(), info.ID)
	}
}

func createReview(t *testing.T) {
	info, err := CreateReview(context.Background(), &npool.ReviewReq{
		AppID:      &ret.AppID,
		Domain:     &ret.Domain,
		ObjectID:   &ret.ObjectID,
		ObjectType: &ret.ObjectType,
		Trigger:    &ret.Trigger,
		ReviewerID: &ret.ReviewerID,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
	}
}

func updateReview(t *testing.T) {
	ret.State = types.ReviewState_Rejected
	ret.StateStr = types.ReviewState_Rejected.String()
	ret.Message = uuid.NewString()
	info, err := UpdateReview(context.Background(), &npool.ReviewReq{
		ID:      &ret.ID,
		State:   &ret.State,
		Message: &ret.Message,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getReviews(t *testing.T) {
	conds := &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		State: &basetypes.Int32Val{
			Op:    cruder.EQ,
			Value: types.ReviewState_value[ret.State.String()],
		},
	}
	infos, _, err := GetReviews(context.Background(), conds, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, 0, len(infos))
	}
}

func getReview(t *testing.T) {
	info, err := GetReview(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func existReviewConds(t *testing.T) {
	exist, err := ExistReviewConds(context.Background(), &npool.ExistReviewCondsRequest{
		Conds: &npool.Conds{
			AppID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: ret.AppID,
			},
			ID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: ret.ID,
			},
		},
	})
	if assert.Nil(t, err) {
		assert.True(t, exist)
	}
}

func getObjectReviews(t *testing.T) {
	infos, err := GetObjectReviews(context.Background(), ret.AppID, ret.Domain, []string{ret.ObjectID}, ret.ObjectType)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, infos[0])
	}
}

func getObjectReview(t *testing.T) {
	info, err := GetObjectReview(context.Background(), ret.AppID, ret.Domain, ret.ObjectID, ret.ObjectType)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteReview(t *testing.T) {
	_, err := DeleteReview(context.Background(), ret.ID)
	assert.Nil(t, err)
	info, err := GetReview(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupApp(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
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
	t.Run("getObjectReviews", getObjectReviews)
	t.Run("getObjectReview", getObjectReview)

	t.Run("deleteReview", deleteReview)

	patch.Unpatch()
}
