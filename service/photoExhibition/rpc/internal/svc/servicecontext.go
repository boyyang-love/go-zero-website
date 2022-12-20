package svc

import (
	"blog/helper"
	"blog/service/photoExhibition/models"
	"blog/service/photoExhibition/rpc/internal/config"
	"fmt"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	username := c.Mysql.Username
	password := c.Mysql.Password
	host := c.Mysql.Host
	port := c.Mysql.Port
	database := c.Mysql.Database
	charset := c.Mysql.Charset
	timeout := c.Mysql.Timeout
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		timeout,
	)
	db, err := helper.InitMysql(args)
	if err == nil {
		fmt.Printf("数据库连接成功!")
		db.AutoMigrate(&models.PhotoExhibition{})
	} else {
		fmt.Printf("数据库连接失败!(%s)", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
