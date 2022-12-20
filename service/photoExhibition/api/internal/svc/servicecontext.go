package svc

import (
	"blog/service/photoExhibition/api/internal/config"
	"blog/service/photoExhibition/rpc/photoexhibitionclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	PhotoExhibitionRpc photoexhibitionclient.PhotoExhibition
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		PhotoExhibitionRpc: photoexhibitionclient.NewPhotoExhibition(zrpc.MustNewClient(c.PhotoExhibitionRpc)),
	}
}
