package logic

import (
	"context"
	"golabs/go-zero/shorturl/rpc/transformer"

	"golabs/go-zero/shorturl/api/internal/svc"
	"golabs/go-zero/shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (*types.ExpandResp, error) {
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{Shorten: req.Shorten})
	if err != nil {
		return nil, err
	}

	return &types.ExpandResp{Url: resp.GetUrl()}, nil
}
