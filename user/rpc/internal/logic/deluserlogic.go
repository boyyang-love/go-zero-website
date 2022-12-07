package logic

import (
	"blog/user/models"
	"context"

	"blog/user/rpc/internal/svc"
	"blog/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *user.DelUserReq) (*user.DelUserRes, error) {
	var userInfo models.User
	res := l.svcCtx.DB.
		Model(&userInfo).
		Where("id = ?", in.Id).
		First(&userInfo)
	if res.RowsAffected == 0 {
		return &user.DelUserRes{
			Code: 0,
			Msg:  "不存在该用户",
		}, nil
	} else {
		err := l.svcCtx.DB.Model(&userInfo).Where("id = ?", in.Id).Delete(&userInfo).Error
		if err == nil {
			return &user.DelUserRes{
				Code: 1,
				Msg:  "删除用户成功",
			}, nil
		} else {
			return nil, err
		}
	}
}
