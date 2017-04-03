// Package router implements the tarantula http router for net/http projects
package router

import	(
		"regexp"
		"strings"
		"strconv"
		"net/http"
		//
		"github.com/golangdaddy/tarantula/log"
		"github.com/golangdaddy/tarantula/router/common"
		)

type WildcardRouter struct {
	handler http.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {}

func (router *WildcardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) { router.handler.ServeHTTP(w, r) }

func (router *WildcardRouter) Serve(port int) error {

  return http.ListenAndServe(":"+strconv.Itoa(port), router)
}

// Creates a new router.
func NewRouter(log logging.Logger, host string) (*common.Node, *WildcardRouter) {

	root := common.Root()

	root.Config.Log = log
	root.Config.Host = host

	f := func (res http.ResponseWriter, r *http.Request) {

		node := common.Root()

		req := NewRequestObject(node, res, r)

		// check for subdomain routing

		subdomain := strings.Split(r.URL.Host, ".")[0]

		subNode := node.Config.SubdomainTrees[subdomain]
		if subNode != nil {

			subNode.MainHandler(req, r.URL.Path)
			return
		}

		node.MainHandler(req, r.URL.Path)

	}

	return root, &WildcardRouter{http.HandlerFunc(f)}
}
