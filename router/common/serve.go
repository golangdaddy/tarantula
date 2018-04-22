package common

import (
		"strings"
		//
		"github.com/golangdaddy/tarantula/web"
		)

const	(
		ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
		)

// main handler
func (node *Node) MainHandler(req web.RequestInterface, fullPath string) (status *web.ResponseStatus) {

	// enforce https-only if required
	if node.Config.forceTLS {
		if !req.IsTLS() {
			status = req.Respond(502, "PLEASE UPGRADE TO HTTPS")
			Respond(req, status)
			return
		}
	}

	// set CORS headers
	for k, v := range node.Config.headers {
		req.SetHeader(k, v)
	}

	// return if preflight request
	if req.Method() == "OPTIONS" { return }

	switch fullPath {

		case "/swagger.json":

			// render the OpenAPI handler documentation

			spec := node.Config.BuildOpenAPISpec(req)

			Respond(req, req.Respond(spec))
			return

		case "/_.json":

			// render the handler documentation

			spec := node.Config.buildSpec(req)

			Respond(req, req.Respond(spec))
			return

		case "/robots.txt":

			req.Write([]byte(ROBOTS_TXT))
			return

		default:

			rootFunc := node.Config.GetRootFunction(fullPath)

			if rootFunc != nil {

				status = node.RunModule(req); if status != nil { Respond(req, status); return }
				status = node.RunModules(req); if status != nil { Respond(req, status); return }
				Respond(req, rootFunc(req))
				return

			}

	}

	segments := strings.Split(fullPath, "/")[1:]

	next := node

	for _, segment := range segments {

		if len(segment) == 0 { break }

		var n *Node
		n, status = next.Next(req, segment)
		if status != nil {
			Respond(req, status)
			return
		}

		if n != nil {
			for k, v := range n.requestParams { req.SetParam(k, v) }
			next = n
			continue
		}

		req.HttpError("NO ROUTE FOUND AT " + next.FullPath() + "/" + segment, 404)
		return
	}

	// resolve handler

	handler := next.handler(req)
	if handler == nil {
		req.HttpError("NO CONTROLLER FOUND AT " + next.FullPath(), 500)
		return
	}

	// read the request body and unmarshal into specified schema
	status = handler.ReadPayload(req)
	if status != nil {
		Respond(req, status)
		return
	}

	// execute modules
	status = handler.Node.RunModules(req)
	if status != nil {
		Respond(req, status)
		return
	}

	if handler.IsFile {

		status = handler.DetectContentType(req, handler.filePath)
		if status != nil {
			Respond(req, status)
			return
		}

		req.SetHeader("Content-Type", handler.fileType)

		status = req.Respond(handler.fileCache)
		Respond(req, status)
		return
	}

	if handler.function == nil {
		req.Log().Panic("FAILED TO GET FUNCTION TO SERVE HANDLE OPERATION!")
	}

	Respond(
		req,
		handler.function(req),
	)
	return
}
