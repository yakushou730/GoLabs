package logic

import (
	"context"
	"golabs/go-zero/shorturl/rpc/internal/svc"
	"golabs/go-zero/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.GetShorten())
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{Url: res.Url}, nil
}
