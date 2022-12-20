package logic

import (
	"blog/service/user/models"
	"blog/service/user/rpc/internal/svc"
	"blog/service/user/rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoRes, error) {
	var userInfo models.User
	res := l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ? or username = ?", in.Id, in.Username).
		First(&models.User{}).
		Scan(&userInfo)
	if res.RowsAffected == 0 {
		return &user.UserInfoRes{
			Code: 0,
			Msg:  "不存在该用户",
		}, nil
	} else {
		return &user.UserInfoRes{
			Code:      1,
			Msg:       "获取用户信息成功",
			Id:        int32(userInfo.Id),
			Username:  userInfo.Username,
			Password:  userInfo.Password,
			Tel:       *userInfo.Tel,
			Email:     *userInfo.Email,
			Age:       int32(userInfo.Age),
			Gender:    int32(userInfo.Gender),
			AvatarUrl: userInfo.AvatarUrl,
		}, nil
	}
}
