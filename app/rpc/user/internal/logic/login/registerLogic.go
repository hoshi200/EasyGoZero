package loginlogic

import (
	"EasyGoZero/app/rpc/user/internal/svc"
	"EasyGoZero/app/rpc/user/userRpcModel"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userRpcModel.ReqRegister) (*userRpcModel.ResRegister, error) {
	// todo: add your logic here and delete this line

	return &userRpcModel.ResRegister{}, nil
}
