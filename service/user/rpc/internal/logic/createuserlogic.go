package logic

import (
	"blog/service/user/models"
	"blog/service/user/rpc/internal/svc"
	"blog/service/user/rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserRes, error) {
	res := l.svcCtx.DB.
		Where("username = ? or tel = ?", in.Username, in.Tel).
		Find(&models.User{})

	if res.RowsAffected == 0 {
		userinfo := models.User{Username: in.Username, Password: in.Password, Tel: &in.Tel}
		err := l.svcCtx.DB.Create(&userinfo).Error
		if err == nil {
			return &user.CreateUserRes{Code: 1, Msg: "创建用户成功", Id: int32(userinfo.Id)}, nil
		} else {
			return &user.CreateUserRes{Code: 0, Msg: err.Error()}, err
		}
	} else {
		return &user.CreateUserRes{Code: 0, Msg: "该用户名已存在或手机号已注册"}, nil
	}
}
