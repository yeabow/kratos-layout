package server

import (
	"buf.build/gen/go/fyy/demo/bufbuild/connect-go/helloworld/v1/helloworldv1connect"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/connect"
)

// NewConnectServer new a gRPC-Connect server.
func NewConnectServer(c *conf.Server, service *service.GreeterService, logger log.Logger) *connect.Server {
	var opts = []connect.ServerOption{
		connect.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Connect.Network != "" {
		opts = append(opts, connect.Network(c.Connect.Network))
	}
	if c.Connect.Addr != "" {
		opts = append(opts, connect.Address(c.Connect.Addr))
	}
	if c.Connect.Timeout != nil {
		opts = append(opts, connect.Timeout(c.Connect.Timeout.AsDuration()))
	}

	srv := connect.NewServer(opts...)
	srv.Handle(helloworldv1connect.NewGreeterHandler(service))
	return srv
}
