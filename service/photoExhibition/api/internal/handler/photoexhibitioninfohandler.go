package handler

import (
	"net/http"

	"blog/service/photoExhibition/api/internal/logic"
	"blog/service/photoExhibition/api/internal/svc"
	"blog/service/photoExhibition/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func photoExhibitionInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhotoExhibitionInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPhotoExhibitionInfoLogic(r.Context(), svcCtx)
		resp, err := l.PhotoExhibitionInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
