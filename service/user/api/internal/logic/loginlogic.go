package logic

import (
	"blog/errorx"
	"blog/helper"
	"blog/service/user/api/internal/svc"
	"blog/service/user/api/internal/types"
	"blog/service/user/rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	auth := l.svcCtx.Config.Auth
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Username: req.Username,
	})

	if err == nil && userInfo.Code == 1 && req.Password == userInfo.Password {
		token, _ := helper.GenerateJwtToken(
			&helper.GenerateJwtStruct{
				Id:       int(userInfo.Id),
				Username: userInfo.Username,
				Password: userInfo.Password,
			},
			auth.AccessSecret,
		)
		if err == nil {
			return &types.LoginRes{
				Id:       int64(userInfo.Id),
				Username: userInfo.Username,
				Token:    token,
			}, nil
		} else {
			return nil, errorx.NewDefaultError("token生成失败")
		}
	} else {
		return nil, errorx.NewDefaultError("账号不存在或者密码错误")
	}
}
