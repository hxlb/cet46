package subscriber

import (
	"context"
	"github.com/hxlb/cet46/proto/example"
	"github.com/micro/go-log"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *com_hxlb_srv_cet46.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *com_hxlb_srv_cet46.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
