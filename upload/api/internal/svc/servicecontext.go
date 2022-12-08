package svc

import (
	"blog/helper"
	"blog/upload/api/internal/config"
	"blog/upload/models"
	"blog/upload/rpc/uploadclient"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UploadRpc uploadclient.Upload
	Client    *cos.Client
	DB        *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 对象存储
	client := helper.InitCloudBase(c.CloudBase.ClientUrl, c.CloudBase.ClientSecretId, c.CloudBase.ClientSecretKey)
	if client == nil {
		fmt.Printf("对象存储初始化失败! \n")
	} else {
		fmt.Printf("对象存储初始化成功! \n")
	}
	// 数据库
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.Database,
		c.Mysql.Charset,
		c.Mysql.Timeout,
	)
	db, err := helper.InitMysql(args)
	if err == nil {
		fmt.Printf("数据库初始化成功! \n")
		db.AutoMigrate(&models.Upload{})
	} else {
		fmt.Printf("数据库初始化失败! (%s) \n", err)
	}
	return &ServiceContext{
		Config:    c,
		UploadRpc: uploadclient.NewUpload(zrpc.MustNewClient(c.UploadRpc)),
		Client:    client,
		DB:        db,
	}
}
