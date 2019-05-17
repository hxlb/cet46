package main

import (
	"github.com/hxlb/cet46/handler"
	"github.com/hxlb/cet46/proto/example"
	"github.com/hxlb/cet46/subscriber"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

func main() {
	//
	reg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// New Service
	service := micro.NewService(
		micro.Name("com.hxlb.srv.cet46"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// Initialise service
	service.Init()

	// Register Handler
	com_hxlb_srv_cet46.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.hxlb.srv.cet46", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.hxlb.srv.cet46", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
