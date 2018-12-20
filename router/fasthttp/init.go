package router

import	(
		"strings"
		"strconv"
		//
		"github.com/valyala/fasthttp"
		"github.com/golangdaddy/tarantula/log"
		//
		"github.com/golangdaddy/tarantula/router/common"
		"github.com/golangdaddy/tarantula/router/common/openapi"
		)

type FastHttpRouter func (ctx *fasthttp.RequestCtx)

func (router FastHttpRouter) Serve(port int) error {

	return fasthttp.ListenAndServe(":"+strconv.Itoa(port), fasthttp.RequestHandler(router))
}

func (router FastHttpRouter) ServeTLS(port int, crt, key string) error {

	return fasthttp.ListenAndServeTLS(":"+strconv.Itoa(port), crt, key, fasthttp.RequestHandler(router))
}

func NewRouter(logger logging.Logger, spec *openapi.APISpec) (*common.Node, FastHttpRouter) {

	root := common.Root()

	root.Config.Spec = spec
	root.Config.Log = logger

	return root, FastHttpRouter(
		func (ctx *fasthttp.RequestCtx) {

			fullPath := string(ctx.Path())

			node := common.Root()
			req := NewRequestObject(node, ctx)

			// check for subdomain routing

			subdomain := strings.Split(string(ctx.Host()), ".")[0]

			subNode := node.Config.SubdomainTrees[subdomain]
			if subNode != nil {

				subNode.MainHandler(req, fullPath)
				return
			}

			node.MainHandler(req, fullPath)

		},
	)
}
