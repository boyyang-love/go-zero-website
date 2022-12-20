package logic

import (
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/svc"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhotoExhibitionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhotoExhibitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhotoExhibitionInfoLogic {
	return &PhotoExhibitionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhotoExhibitionInfoLogic) PhotoExhibitionInfo(in *photoExhibition.PhotoExhibitionInfoReq) (*photoExhibition.PhotoExhibitionInfoRes, error) {
	var info models.PhotoExhibition
	err := l.svcCtx.DB.
		Model(&models.PhotoExhibition{}).
		Where("id = ?", in.Id).
		First(&info).
		Error
	if err == nil {
		return &photoExhibition.PhotoExhibitionInfoRes{
			Code:      1,
			Msg:       "获取信息成功!",
			Id:        int32(info.Id),
			Title:     info.Title,
			SubTitle:  info.SubTitle,
			Des:       info.Des,
			Cover:     info.Cover,
			UserId:    int32(info.UserId),
			Status:    int32(info.Status),
			RejectRes: info.RejectRes,
		}, nil
	} else {
		return &photoExhibition.PhotoExhibitionInfoRes{
			Code: 0,
			Msg:  "获取信息失败!",
		}, err
	}
}
