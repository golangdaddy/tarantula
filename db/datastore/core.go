package ds

import (
	"golang.org/x/net/context"
	datastore "cloud.google.com/go/datastore"
	datastoreAE "google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	//
	"github.com/golangdaddy/tarantula/web"
)

func (client *Client) GetIncomplete(req web.RequestInterface, entityType string, intId int64, dst interface{}) (bool, error) {
	return client.GetStruct(req, entityType, "", intId, nil, dst)
}

func (client *Client) GetStruct(req web.RequestInterface, entityType, keyName string, intId int64, ancestor, dst interface{}) (bool, error) {

	var key interface{}

	if client.appEngine {
		ctx := appengine.NewContext(req.R())

		parent, _ := ancestor.(*datastoreAE.Key)
		key = datastoreAE.NewKey(ctx, entityType, keyName, intId, parent)
	} else {
		parent, _ := ancestor.(*datastore.Key)
		key = datastore.NameKey(entityType, keyName, parent)
	}
	return client.GetKey(req, key, dst)
}

func (client *Client) GetKey(req web.RequestInterface, key interface{}, dst interface{}) (bool, error) {

	var err error
	var notFound string

	if client.appEngine {
		notFound = datastoreAE.ErrNoSuchEntity.Error()
		err = datastoreAE.Get(
			appengine.NewContext(req.R()),
			key.(*datastoreAE.Key),
			dst,
		)
	} else {
		notFound = datastore.ErrNoSuchEntity.Error()
		err = client.Get(context.Background(), key.(*datastore.Key), dst)
	}

	if err != nil {
		if err.Error() == notFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (client *Client) PutIncomplete(req web.RequestInterface, entityType string, ancestor, src interface{}) error {

	return client.PutStruct(req, entityType, "", 0, ancestor, src)
}

func (client *Client) PutStruct(req web.RequestInterface, entityType, keyName string, intId int64, ancestor, src interface{}) error {

	var key interface{}

	if client.appEngine {
		ctx := appengine.NewContext(req.R())
		parent, _ := ancestor.(*datastoreAE.Key)
		key = datastoreAE.NewKey(ctx, entityType, keyName, intId, parent)
	} else {
		parent, _ := ancestor.(*datastore.Key)
		key = datastore.NameKey(entityType, keyName, parent)
	}

	return client.PutKey(req, key, src)
}

func (client *Client) DeleteStruct(req web.RequestInterface, entityType, keyName string, intId int64, ancestor interface{}) error {

	var key interface{}

	if client.appEngine {
		ctx := appengine.NewContext(req.R())
		parent, _ := ancestor.(*datastoreAE.Key)
		key = datastoreAE.NewKey(ctx, entityType, keyName, intId, parent)
	} else {
		parent, _ := ancestor.(*datastore.Key)
		key = datastore.NameKey(entityType, keyName, parent)
	}

	return client.DeleteKey(req, key)
}

func (client *Client) PutKey(req web.RequestInterface, key interface{}, src interface{}) error {

	var err error

	_, ok := req.(*web.TestInterface)
	if ok {
		return nil
	}

	if client.appEngine {
		_, err = datastoreAE.Put(
			appengine.NewContext(req.R()),
			key.(*datastoreAE.Key),
			src,
		)
	} else {
		_, err = client.Put(
			context.Background(),
			key.(*datastore.Key),
			src,
		)
	}

	return err
}

func (client *Client) DeleteKey(req web.RequestInterface, key interface{}) error {

	var err error

	if client.appEngine {
		err = datastoreAE.Delete(
			appengine.NewContext(req.R()),
			key.(*datastoreAE.Key),
		)
	} else {
		err = client.Delete(
			context.Background(),
			key.(*datastore.Key),
		)
	}

	return err
}
