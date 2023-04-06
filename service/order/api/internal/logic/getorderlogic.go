package logic

import (
	"context"
	"errors"
	"github.com/linx93/microservice-demo/service/user/rpc/types/user"

	"github.com/linx93/microservice-demo/service/order/api/internal/svc"
	"github.com/linx93/microservice-demo/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	getUser, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	l.Debug("getUser=%v", getUser)
	if err != nil {
		return nil, err
	}

	if getUser.Name != "linx" {
		return nil, errors.New("用户不存在")
	}

	return &types.OrderReply{
		Id:   getUser.Id,
		Name: getUser.Name,
	}, nil

	return
}
