package logic

import (
	"blog/errorx"
	"blog/service/user/api/internal/svc"
	"blog/service/user/api/internal/types"
	"blog/service/user/rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	res, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserReq{
		Username: req.Username,
		Password: req.Password,
		Tel:      req.Tel,
	})
	if err == nil && res.Code == 1 {
		return &types.RegisterRes{
			Id: int(res.Id),
		}, nil
	} else {
		if err == nil {
			return nil, errorx.NewDefaultError(res.Msg)
		} else {
			return nil, err
		}
	}
}
