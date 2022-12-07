package logic

import (
	"blog/errorx"
	"blog/user/api/internal/svc"
	"blog/user/api/internal/types"
	"blog/user/models"
	"blog/user/rpc/types/user"
	"context"
	"github.com/golang-jwt/jwt/v4"

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
		token, err := l.GetJwtToken(auth.AccessSecret, auth.AccessExpire, models.User{Id: uint(userInfo.Id)})
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

func (l *LoginLogic) GetJwtToken(secretKey string, time int, userInfo models.User) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userInfo.Id,
		"username": userInfo.Username,
		"password": userInfo.Password,
	})

	token, err = claims.SignedString([]byte(secretKey))

	return token, err
}
