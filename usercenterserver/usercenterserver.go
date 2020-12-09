// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./usercenterserver_mock.go -package usercenterserver -source $GOFILE

package usercenterserver

import (
	"context"

	"github.com/nizonglonggit/usercenter/usercenter"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	LoginResp       = usercenter.LoginResp
	BaseResp        = usercenter.BaseResp
	RegisterUserReq = usercenter.RegisterUserReq
	LoginReq        = usercenter.LoginReq

	UserCenterServer interface {
		RegisterUser(ctx context.Context, in *RegisterUserReq) (*BaseResp, error)
		Login(ctx context.Context, in *LoginReq) (*LoginResp, error)
	}

	defaultUserCenterServer struct {
		cli zrpc.Client
	}
)

func NewUserCenterServer(cli zrpc.Client) UserCenterServer {
	return &defaultUserCenterServer{
		cli: cli,
	}
}

func (m *defaultUserCenterServer) RegisterUser(ctx context.Context, in *RegisterUserReq) (*BaseResp, error) {
	client := usercenter.NewUserCenterServerClient(m.cli.Conn())
	return client.RegisterUser(ctx, in)
}

func (m *defaultUserCenterServer) Login(ctx context.Context, in *LoginReq) (*LoginResp, error) {
	client := usercenter.NewUserCenterServerClient(m.cli.Conn())
	return client.Login(ctx, in)
}