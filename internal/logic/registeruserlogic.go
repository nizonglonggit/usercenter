package logic

import (
	"context"

	"github.com/nizonglonggit/usercenter/internal/svc"
	"github.com/nizonglonggit/usercenter/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterUserReq) (*user.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &user.BaseResp{}, nil
}
