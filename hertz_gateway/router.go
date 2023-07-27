// Code generated by hertz generator.

package main

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func customizedRegister(r *server.Hertz) {
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "the api gateway is running")
	})
}
