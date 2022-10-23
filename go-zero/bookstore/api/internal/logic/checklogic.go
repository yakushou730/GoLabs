package logic

import (
	"context"
	"gozerolabs/bookstore/rpc/check/checker"

	"gozerolabs/bookstore/api/internal/svc"
	"gozerolabs/bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (*types.CheckResp, error) {
	r, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{Book: req.Book})
	if err != nil {
		return nil, err
	}

	return &types.CheckResp{
		Found: r.GetFound(),
		Price: r.GetPrice(),
	}, nil
}
