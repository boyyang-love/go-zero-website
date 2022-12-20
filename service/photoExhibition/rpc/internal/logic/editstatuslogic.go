package logic

import (
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/svc"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditStatusLogic {
	return &EditStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditStatusLogic) EditStatus(in *photoExhibition.EditStatusReq) (*photoExhibition.EditStatusRes, error) {
	err := l.svcCtx.DB.
		Model(&models.PhotoExhibition{}).
		Select("status", "reject_res").
		Where("id = ?", in.Id).
		Updates(&models.PhotoExhibition{Status: int(in.Status), RejectRes: in.RejectRes}).
		Error
	if err == nil {
		return &photoExhibition.EditStatusRes{
			Code: 1,
			Msg:  "状态修改成功",
		}, nil
	} else {
		return &photoExhibition.EditStatusRes{
			Code: 0,
			Msg:  "状态修改失败",
		}, err
	}
}
