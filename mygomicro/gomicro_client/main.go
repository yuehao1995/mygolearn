package main

import (
	"context"
	"fmt"

	proto "mygolearn/mygomicro/proto" //这里写你的proto文件放置路劲

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcd"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	fmt.Println(rsp)
	if err != nil {
		fmt.Println(err)
	}

	reg := etcd.NewRegistry(registry.Addrs("http://192.168.1.204:2379", "http://192.168.1.205:2379", "http://192.168.1.206:2379"))
	service1 := micro.NewService(
		// Set service name
		micro.Name("greeter.client"),
		micro.Registry(reg),
	)
	greeter = proto.NewGreeterService("greeter.client", service1.Client())
	rsp1, err1 := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	fmt.Println(rsp1)
	if err != nil {
		fmt.Println(err1)
	}

}
