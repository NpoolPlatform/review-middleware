package review

import (
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
		TriggerStr:    npool.ReviewTriggerType(npool.ReviewTriggerType_value[row.Trigger]).String(),
		Trigger:       npool.ReviewTriggerType(npool.ReviewTriggerType_value[row.Trigger]),
		ObjectTypeStr: npool.ReviewObjectType(npool.ReviewObjectType_value[row.ObjectType]).String(),
		ObjectType:    npool.ReviewObjectType(npool.ReviewObjectType_value[row.ObjectType]),
		StateStr:      npool.ReviewState(npool.ReviewState_value[row.State]).String(),
		State:         npool.ReviewState(npool.ReviewState_value[row.State]),
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
