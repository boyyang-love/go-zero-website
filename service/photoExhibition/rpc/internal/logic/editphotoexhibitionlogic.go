package logic

import (
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/svc"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditPhotoExhibitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditPhotoExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPhotoExhibitionLogic {
	return &EditPhotoExhibitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditPhotoExhibitionLogic) EditPhotoExhibition(in *photoExhibition.EditPhotoExhibitionReq) (*photoExhibition.EditPhotoExhibitionRes, error) {
	err := l.svcCtx.DB.
		Model(&models.PhotoExhibition{}).
		Select("title", "sub_title", "des").
		Where("id = ?", in.Id).
		Updates(models.PhotoExhibition{
			Title:    in.Title,
			SubTitle: in.SubTitle,
			Des:      in.Des,
		}).
		Error
	if err == nil {
		return &photoExhibition.EditPhotoExhibitionRes{
			Code: 1,
			Msg:  "更新成功!",
		}, nil
	} else {
		return &photoExhibition.EditPhotoExhibitionRes{
			Code: 0,
			Msg:  "更新失败!",
		}, err
	}
}
