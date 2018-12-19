package main

import (
	"fmt"
	"errors"
	"testing"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/web/validation"
	"github.com/golangdaddy/tarantula/log/testing"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/router/common/openapi"
	"github.com/golangdaddy/tarantula/router/standard"
)

const (
	CONST_SPEC_HOST = "dummyhost"
	CONST_SPEC_TITLE = "my title"
	CONST_SPEC_URL = "https://example.com"
	CONST_SPEC_EMAIL = "address@example.com"
	CONST_SPEC_LICENSE = "MIT"
	CONST_SPEC_BASEPATH = "/api"
)

type TestObject struct {
	Hello string
	World int
}

func TestMain(t *testing.T) {

	logClient := logs.NewClient().NewLogger()

	s := openapi.NewSpec(CONST_SPEC_HOST, CONST_SPEC_TITLE)
	s.BasePath = CONST_SPEC_BASEPATH
	s.Info.Contact.URL = CONST_SPEC_URL
	s.Info.Contact.Email = CONST_SPEC_EMAIL
	s.Info.License.URL = CONST_SPEC_URL

	root, _ := router.NewRouter(logClient, s)

	api := root.Add(CONST_SPEC_BASEPATH)

	api.GET(
		dummyHandler,
	).Describe(
		"This is a GET endpoint!",
	).Response(
		TestObject{},
	)

	api.POST(
		dummyHandler,
	).Describe(
		"This is a POST endpoint!",
	).Body(
		&common.Payload{
			"hello": validation.String(10, 20).Description("The hellos!").Default("Helloy"),
			"world": validation.Int().Description("The worlds!").Default(2),
		},
	).Response(
		TestObject{},
	)

	apiResource := api.Add("/resource").Param(
		validation.String(1, 64).Description("The id of the user."),
		"id",
	)

		apiResource.GET(
			dummyHandler,
		).Describe(
			"Handles access to the resource",
		).Response(
			TestObject{},
		)

	req := web.NewTestInterface("", "")
	spec := root.Config.BuildOpenAPISpec(req)

	t.Run(
		"Test the spec",
		func (t *testing.T) {

			if spec.Host != CONST_SPEC_HOST {
				t.Error(errors.New("SPEC HAS INVALID HOST!"))
			}

			if spec.Info.Title != CONST_SPEC_TITLE {
				t.Error(errors.New("SPEC HAS INVALID TITLE!"))
			}

			if spec.Info.Contact.URL != CONST_SPEC_URL {
				t.Error(errors.New("SPEC HAS INVALID CONTACT URL!"))
			}

			if len(spec.Paths) != 2 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUM OF PATHS! %v", len(spec.Paths)))
			}

			if spec.Paths["/resource/:id"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			if len(spec.Paths["/resource/:id"].GET.Parameters) != 1 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUMBER OF PARAMETERS! %v", len(spec.Paths)))
			}

			if len(spec.Definitions) != 1 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUMBER OF DEFINITIONS! %v", len(spec.Definitions)))
			}

		},
	)

	b, _ := json.Marshal(spec)
	fmt.Println(string(b))
}

func dummyHandler(req web.RequestInterface) *web.ResponseStatus {

	return nil
}
