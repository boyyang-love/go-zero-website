// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"blog/service/photoExhibition/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/exhibtion/create",
				Handler: createPhotoExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/exhibtion/info",
				Handler: photoExhibitionInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
