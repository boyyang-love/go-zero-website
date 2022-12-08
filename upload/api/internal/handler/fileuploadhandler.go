package handler

import (
	"blog/helper"
	"blog/response"
	"blog/upload/api/internal/logic"
	"blog/upload/api/internal/svc"
	"blog/upload/api/internal/types"
	"blog/upload/models"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"
)

func fileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 文件处理
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		hash, _ := helper.MakeFileHash(file, fileHeader)
		fileInfo := new(models.Upload)
		has := svcCtx.DB.Model(&models.Upload{}).Where("hash = ?", hash).First(&fileInfo)

		if has.RowsAffected != 0 {
			response.Response(w, &types.UploadFileRes{
				FileName: fileInfo.FileName,
				FilePath: fileInfo.FileName,
			}, nil)

			return
		}

		url, err := helper.CosFileUpload(svcCtx.Client, file, fileHeader, "images")
		if err != nil {
			httpx.Error(w, err)

			return
		}

		svcCtx.DB.Model(&models.Upload{}).Create(&models.Upload{
			Hash:     hash,
			FileName: fileHeader.Filename,
			FilePath: url,
			Ext:      path.Ext(fileHeader.Filename),
			Size:     fileHeader.Size,
			UserId:   0,
		})

		req.FilePath = url
		req.FileName = fileHeader.Filename
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Ext = path.Ext(fileHeader.Filename)

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)

		response.Response(w, resp, err)

		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
