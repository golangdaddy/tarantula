package ds

import (
    "testing"
    //
    "cloud.google.com/go/datastore"
    //
    "github.com/golangdaddy/tarantula/web"
    //
    "github.com/blokhubio/models"
)

func TestRunQuery(t *testing.T) {

    client := &Client{}

    req := web.NewTestInterface()

    q := datastore.NewQuery(models.CONST_DS_ENTITY_ACCOUNT_NAMESPACE)

    dst := []*models.NamespaceAccount{}

    _, err := client.RunQuery(req, q, nil, &dst)
    if err != nil {
        t.Error(err)
        return
    }

    req.Log().DebugJSON(dst)
}
