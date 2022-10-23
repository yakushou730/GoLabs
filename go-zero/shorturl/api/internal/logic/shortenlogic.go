package logic

import (
	"context"
	"golabs/go-zero/shorturl/rpc/transform"

	"golabs/go-zero/shorturl/api/internal/svc"
	"golabs/go-zero/shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenReq{Url: req.Url})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{Shorten: resp.GetShorten()}, nil
}
