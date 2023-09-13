package review

import (
	types "github.com/NpoolPlatform/message/npool/basetypes/review/v1"
	npool "github.com/NpoolPlatform/message/npool/review/mw/v2/review"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent"
)

func Ent2Grpc(row *ent.Review) *npool.Review {
	if row == nil {
		return nil
	}

	return &npool.Review{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		ReviewerID:    row.ReviewerID.String(),
		Domain:        row.Domain,
		ObjectID:      row.ObjectID.String(),
		TriggerStr:    types.ReviewTriggerType(types.ReviewTriggerType_value[row.Trigger]).String(),
		Trigger:       types.ReviewTriggerType(types.ReviewTriggerType_value[row.Trigger]),
		ObjectTypeStr: types.ReviewObjectType(types.ReviewObjectType_value[row.ObjectType]).String(),
		ObjectType:    types.ReviewObjectType(types.ReviewObjectType_value[row.ObjectType]),
		StateStr:      types.ReviewState(types.ReviewState_value[row.State]).String(),
		State:         types.ReviewState(types.ReviewState_value[row.State]),
		Message:       row.Message,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Review) []*npool.Review {
	infos := []*npool.Review{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
