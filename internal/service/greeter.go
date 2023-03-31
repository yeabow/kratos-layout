package service

import (
	helloworldv1 "buf.build/gen/go/fyy/demo/protocolbuffers/go/helloworld/v1"
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	data *conf.Data
	log  *log.Helper
}

// SayHello implements helloworldv1.GreeterServer.
func (g GreeterService) SayHello(ctx context.Context, c *connect.Request[helloworldv1.HelloRequest]) (*connect.Response[helloworldv1.HelloReply], error) {
	return connect.NewResponse[helloworldv1.HelloReply](&helloworldv1.HelloReply{Message: "Hello " + c.Msg.Name}), nil
}

// NewGreeterService new a greeter service.
func NewGreeterService(data *conf.Data, logger log.Logger) *GreeterService {
	return &GreeterService{data: data, log: log.NewHelper(logger)}
}
