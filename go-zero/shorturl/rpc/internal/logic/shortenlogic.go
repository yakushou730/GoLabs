package logic

import (
	"context"
	"gozerolabs/shorturl/rpc/model"

	"github.com/zeromicro/go-zero/core/hash"

	"gozerolabs/shorturl/rpc/internal/svc"
	"gozerolabs/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	key := hash.Md5Hex([]byte(in.GetUrl()))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		Shorten: key,
		Url:     in.GetUrl(),
	})
	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{Shorten: key}, nil
}
