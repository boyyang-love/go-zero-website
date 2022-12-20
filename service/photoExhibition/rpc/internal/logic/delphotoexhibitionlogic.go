package logic

import (
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/svc"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPhotoExhibitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPhotoExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPhotoExhibitionLogic {
	return &DelPhotoExhibitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPhotoExhibitionLogic) DelPhotoExhibition(in *photoExhibition.DelPhotoExhibitionReq) (*photoExhibition.DelPhotoExhibitionRes, error) {
	err := l.svcCtx.DB.
		Model(&models.PhotoExhibition{}).
		Where("id = ?", in.Id).
		Delete(&models.PhotoExhibition{}).
		Error
	if err == nil {
		return &photoExhibition.DelPhotoExhibitionRes{
			Code: 1,
			Msg:  "删除成功",
		}, nil
	} else {
		return &photoExhibition.DelPhotoExhibitionRes{
			Code: 0,
			Msg:  "删除失败",
		}, err
	}
}
