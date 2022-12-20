package logic

import (
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/svc"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoExhibitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePhotoExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoExhibitionLogic {
	return &CreatePhotoExhibitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePhotoExhibitionLogic) CreatePhotoExhibition(in *photoExhibition.CreatePhotoExhibitionReq) (*photoExhibition.CreatePhotoExhibitionRes, error) {
	pe := &models.PhotoExhibition{
		Title:    in.Title,
		SubTitle: in.SubTitle,
		Des:      in.Des,
		Cover:    in.Cover,
		UserId:   uint(in.UserId),
	}

	res := l.svcCtx.DB.Create(pe)

	if res.Error == nil {
		return &photoExhibition.CreatePhotoExhibitionRes{
			Code: 1,
			Msg:  "创建成功",
			Id:   int32(pe.Id),
		}, nil
	} else {
		return &photoExhibition.CreatePhotoExhibitionRes{
			Code: 1,
			Msg:  "创建失败",
		}, res.Error
	}
}
