// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"EasyGoZero/app/rpc/user/internal/logic/login"
	"EasyGoZero/app/rpc/user/internal/svc"
	"EasyGoZero/app/rpc/user/userRpcModel"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	userRpcModel.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServer) Register(ctx context.Context, in *userRpcModel.ReqRegister) (*userRpcModel.ResRegister, error) {
	l := loginlogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *LoginServer) Login(ctx context.Context, in *userRpcModel.ReqLogin) (*userRpcModel.ResLogin, error) {
	l := loginlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
