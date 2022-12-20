package logic

import (
	"blog/errorx"
	"blog/service/photoExhibition/api/internal/svc"
	"blog/service/photoExhibition/api/internal/types"
	"blog/service/photoExhibition/rpc/types/photoExhibition"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePhotoExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoExhibitionLogic {
	return &CreatePhotoExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePhotoExhibitionLogic) CreatePhotoExhibition(req *types.CreatePhotoExhibitionReq) (resp *types.CreatePhotoExhibitionRes, err error) {
	userId, _ := l.ctx.Value("id").(json.Number).Int64()

	res, err := l.svcCtx.PhotoExhibitionRpc.CreatePhotoExhibition(l.ctx, &photoExhibition.CreatePhotoExhibitionReq{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Des:      req.Des,
		Cover:    req.Cover,
		UserId:   int32(userId),
	})

	if err == nil && res.Code == 1 {
		return &types.CreatePhotoExhibitionRes{
			Id: int(res.Id),
		}, nil
	} else {
		return nil, errorx.NewDefaultError(res.Msg)
	}
}
