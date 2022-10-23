package logic

import (
	"context"
	"golabs/go-zero/bookstore/rpc/model"

	"golabs/go-zero/bookstore/rpc/add/add"
	"golabs/go-zero/bookstore/rpc/add/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Book{
		Book:  in.GetBook(),
		Price: in.GetPrice(),
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{Ok: true}, nil
}
