package logic

import (
	"blog/service/upload/api/internal/svc"
	"blog/service/upload/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.UploadFileReq) (resp *types.UploadFileRes, err error) {
	return &types.UploadFileRes{
		FileName: req.FileName,
		FilePath: req.FilePath,
	}, nil
}
