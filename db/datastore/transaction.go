package ds

import (
	"net/http"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine"
//	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	//
	"github.com/golangdaddy/tarantula/web"
)

func (client *Client) RunInTransaction(req web.RequestInterface, f func (context.Context) error) error {

	ctx := appengine.NewContext(req.R().(*http.Request))

	if client.appEngine {
		return datastoreAE.RunInTransaction(
			ctx,
			f,
			nil,
		)
	}

	return f(ctx)
}
