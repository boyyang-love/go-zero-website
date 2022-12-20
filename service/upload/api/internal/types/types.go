// Code generated by goctl. DO NOT EDIT.
package types

type UploadFileReq struct {
	Hash     string `json:"hash,optional"`
	FileName string `json:"file_name,optional"`
	Ext      string `json:"ext,optional"`
	Size     int64  `json:"size,optional"`
	FilePath string `json:"file_path,optional"`
}

type UploadFileRes struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}