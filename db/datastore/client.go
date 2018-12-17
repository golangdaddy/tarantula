package ds

import (
	"golang.org/x/net/context"
	//
	"cloud.google.com/go/datastore"
)

func NewClient(ctx *context.Context, projectId string) Client {

	if ctx != nil {
		ds, err := datastore.NewClient(*ctx, projectId)
		if err != nil {
			panic(err)
		}
		return Client{
			false,
			ds,
		}
	}
	return Client{
		appEngine: true,
	}
}

type Client struct {
	appEngine bool
	*datastore.Client
}

func (client *Client) IsAppEngine() bool {
	return client.appEngine
}
