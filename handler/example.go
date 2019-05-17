package handler

import (
	"context"
	"github.com/hxlb/cet46/proto/example"

	"github.com/micro/go-log"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *com_hxlb_srv_cet46.Request, rsp *com_hxlb_srv_cet46.Response) error {
	log.Log("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *com_hxlb_srv_cet46.StreamingRequest, stream com_hxlb_srv_cet46.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&com_hxlb_srv_cet46.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream com_hxlb_srv_cet46.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&com_hxlb_srv_cet46.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
