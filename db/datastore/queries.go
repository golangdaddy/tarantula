package ds

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	//
	"github.com/golangdaddy/tarantula/web"
)

func (client *Client) RunQuery(req web.RequestInterface, q, s, dst interface{}) error {

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
