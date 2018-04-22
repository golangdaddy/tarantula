package testing

import (
	"github.com/golangdaddy/tarantula/router/testing/manifest"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/httpclient"
)

type App struct {
	List []*common.HandlerSpec
	Manifest *manifest.Manifest
	*httpclient.Client
}

func NewApp() *App {
	return &App{
		[]*common.HandlerSpec{},
		nil,
		httpclient.NewClient(nil),
	}
}
