package ds

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	//
	"github.com/golangdaddy/tarantula/web"
)

func (client *Client) RunKeysQuery(req web.RequestInterface, query *datastore.Query) ([]*datastore.Key, error) {

	keys, err := client.GetAll(context.Background(), query, nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (client *Client) RunKeysQueryAE(req web.RequestInterface, query *datastoreAE.Query) ([]*datastoreAE.Key, error) {

	ctx := appengine.NewContext(req.R())
	keys, err := query.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (client *Client) RunQuery(req web.RequestInterface, q, dst interface{}) error {

	_, ok := req.(*web.TestInterface)
	if ok {
		return nil
	}

	switch query := q.(type) {

		case *datastore.Query:

			_, err := client.GetAll(context.Background(), query, dst)
			if err != nil {
				return err
			}

		case *datastoreAE.Query:

			ctx := appengine.NewContext(req.R())
			_, err := query.GetAll(ctx, dst)
			if err != nil {
				return err
			}

	}

	return nil
}
