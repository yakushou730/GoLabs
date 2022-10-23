package svc

import (
	"golabs/go-zero/bookstore/api/internal/config"
	"golabs/go-zero/bookstore/rpc/add/adder"
	"golabs/go-zero/bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
