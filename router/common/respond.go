package common

import (
	"errors"
	"strconv"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/markup"
)

func Respond(req web.RequestInterface, status *web.ResponseStatus) error {

	// return with no action if handler returns nil
	if status == nil { return nil }

	switch v := status.Value.(type) {

		case nil:

		case string:

			req.Write([]byte(v))

		case []byte:

			req.Write(v)

		case [][]byte:

			for _, b := range v {
				req.Write(b)
			}

		case *g.ELEMENT:

			b, err := v.Render()
			if err != nil {
				return err
			}
			req.Write(b)

		default:

			req.SetHeader("Content-Type", "application/json")
			b, err := json.Marshal(status.Value)
			if err != nil {
				return err
			}
			req.Write(b)

	}

	if status.Code >= 200 && status.Code < 300 { return nil }

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.MessageString()

	req.HttpError(statusMessage, status.Code)

	return errors.New(statusMessage)
}
