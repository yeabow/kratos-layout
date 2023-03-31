package server

import (
	"buf.build/gen/go/fyy/demo/bufbuild/connect-go/helloworld/v1/helloworldv1connect"
	helloworldv1 "buf.build/gen/go/fyy/demo/protocolbuffers/go/helloworld/v1"
	"context"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"
	"testing"
)

func TestConnect(t *testing.T) {
	client := helloworldv1connect.NewGreeterClient(http.DefaultClient,
		"http://localhost:8000")
	res, err := client.SayHello(
		context.Background(),
		connect.NewRequest(&helloworldv1.HelloRequest{Name: "fyy"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Message)
}
