package router

import	(
		"regexp"
		"strings"
		"net/http"
		//
		"github.com/golangdaddy/tarantula/router/common"
		)

type WildcardRouter struct {
	handler http.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {}

func (router *WildcardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) { router.handler.ServeHTTP(w, r) }

// create a new router for an app
func NewRouter(projectName, host string) (*common.Node, *WildcardRouter) {

	root := common.Root()

	root.Config.ProjectName = projectName
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