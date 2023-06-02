package main

import (
	"Orbital_Hebao/kitex_servers/kitex_gen/sum/sumsvc"
	handler "Orbital_Hebao/kitex_servers/kitex_handler"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
)

func main() {
	r2, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}

	svr2 := sumsvc.NewServer(
		new(handler.SumImpl),
		server.WithRegistry(r2),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "sum"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8890}),
	)

	err2 := svr2.Run()
	if err2 != nil {
		log.Println(err2.Error())
	}
}