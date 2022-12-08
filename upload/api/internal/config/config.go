package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UploadRpc zrpc.RpcClientConf
	Mysql     struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
		Charset  string
		Timeout  string
	}
	CloudBase struct {
		ClientUrl       string
		ClientSecretId  string
		ClientSecretKey string
	}
}
