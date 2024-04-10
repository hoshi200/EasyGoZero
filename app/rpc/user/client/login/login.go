// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package login

import (
	"context"

	"EasyGoZero/app/rpc/user/userRpcModel"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqLogin    = userRpcModel.ReqLogin
	ReqRegister = userRpcModel.ReqRegister
	ResLogin    = userRpcModel.ResLogin
	ResRegister = userRpcModel.ResRegister

	Login interface {
		Register(ctx context.Context, in *ReqRegister, opts ...grpc.CallOption) (*ResRegister, error)
		Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ResLogin, error)
	}

	defaultLogin struct {
		cli zrpc.Client
	}
)

func NewLogin(cli zrpc.Client) Login {
	return &defaultLogin{
		cli: cli,
	}
}

func (m *defaultLogin) Register(ctx context.Context, in *ReqRegister, opts ...grpc.CallOption) (*ResRegister, error) {
	client := userRpcModel.NewLoginClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultLogin) Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ResLogin, error) {
	client := userRpcModel.NewLoginClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
