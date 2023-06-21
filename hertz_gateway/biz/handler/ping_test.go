// Code generated by hertz generator.

package handler

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *app.RequestContext
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Ping",
			args: args{
				ctx: context.TODO(),
				c:   &app.RequestContext{},
			},
		},
		{
			name: "Test Ping with Context",
			args: args{
				ctx: context.WithValue(context.Background(), struct{ key string }{key: "demokey"}, struct{ value string }{value: "demovalue"}),
				c:   &app.RequestContext{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ping(tt.args.ctx, tt.args.c)

			expectedStatus := consts.StatusOK
			expectedResponse := `{"message": "pong"}`

			assert.Equal(t, expectedStatus, tt.args.c.Response.StatusCode())
			assert.JSONEq(t, expectedResponse, string(tt.args.c.Response.Body()))
		})
	}
}
