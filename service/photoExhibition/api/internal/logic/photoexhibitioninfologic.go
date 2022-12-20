package logic

import (
	"context"

	"blog/service/photoExhibition/api/internal/svc"
	"blog/service/photoExhibition/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhotoExhibitionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhotoExhibitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhotoExhibitionInfoLogic {
	return &PhotoExhibitionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhotoExhibitionInfoLogic) PhotoExhibitionInfo(req *types.PhotoExhibitionInfoReq) (resp *types.PhotoExhibitionInfoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
