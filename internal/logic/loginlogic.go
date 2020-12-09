package logic

import (
	"context"

	"github.com/nizonglonggit/usercenter/internal/svc"
	"github.com/nizonglonggit/usercenter/usercenter"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &usercenter.LoginResp{}, nil
}
