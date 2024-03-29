// Code generated by Kitex v0.5.2. DO NOT EDIT.
package sumsvc

import (
	sum "github.com/Linda-ui/orbital_HeBao/kitex_services/sum/kitex_gen/sum"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler sum.SumSvc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
