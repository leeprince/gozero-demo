// Code generated by goctl. DO NOT EDIT.
// Source: demogrpc.proto

package demogrpcclient

import (
	"context"

	"gozero-demo/demogrpc/demogrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = demogrpc.Request
	Response = demogrpc.Response

	Demogrpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDemogrpc struct {
		cli zrpc.Client
	}
)

func NewDemogrpc(cli zrpc.Client) Demogrpc {
	return &defaultDemogrpc{
		cli: cli,
	}
}

func (m *defaultDemogrpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := demogrpc.NewDemogrpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
