package handler

import (
	"blog/response"
	"net/http"

	"blog/service/photoExhibition/api/internal/logic"
	"blog/service/photoExhibition/api/internal/svc"
	"blog/service/photoExhibition/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createPhotoExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePhotoExhibitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreatePhotoExhibitionLogic(r.Context(), svcCtx)
		resp, err := l.CreatePhotoExhibition(&req)
		response.Response(w, resp, err)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
