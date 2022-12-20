package logic

import (
	"blog/service/user/models"
	"blog/service/user/rpc/internal/svc"
	"blog/service/user/rpc/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserRes, error) {
	res := l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", in.Id).
		First(&models.User{})

	if res.RowsAffected == 0 {
		return &user.UpdateUserRes{
			Code: 0,
			Msg:  "不存在此用户",
		}, nil
	}

	err := l.svcCtx.DB.
		Model(&models.User{}).
		Select("Username", "Gender", "Age", "Email", "AvatarUrl", "Tel").
		Where("id = ?", in.Id).
		Updates(&models.User{
			Username:  in.Username,
			Gender:    int(in.Gender),
			Age:       int(in.Age),
			Email:     &in.Email,
			AvatarUrl: in.AvatarUrl,
			Tel:       &in.Tel,
		}).
		Error
	if err == nil {
		return &user.UpdateUserRes{
			Code: 1,
			Msg:  "用户信息更新成功",
		}, nil
	} else {
		return nil, err
	}
}
