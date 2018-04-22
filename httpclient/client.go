package httpclient

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	//
	"github.com/dghubble/sling"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) *Client {
	if c == nil {
		c = &http.Client{}
	}
	return &Client{c}
}

func NewUrlfetchClient(ctx context.Context, seconds ...int) *Client {

	if len(seconds) > 0 {
		ctx, _ = context.WithDeadline(
			ctx,
			time.Now().Add(time.Duration(1000000000 * seconds[0]) * time.Second),
		)
	}

	return &Client{
		urlfetch.Client(ctx),
	}
}

func (client *Client) New(httpRequest *http.Request) *Request {
	return &Request{
		client,
		httpRequest,
		1,
		"200",
	}
}

func (client *Client) Get(url string, dst interface{}, headers ...map[string]string) ([]byte, error) {

	httpRequest, err := sling.New().Get(url).Request()
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		for k, v := range header {
			httpRequest.Header.Set(k, v)
		}
	}

	b, err := client.New(httpRequest).DoReq()
	if err != nil {
		return nil, err
	}

	if dst != nil {
		err = json.Unmarshal(b, dst)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	return b, nil
}

func (client *Client) Post(url string, src, dst interface{}, headers ...map[string]string) ([]byte, error) {

	httpRequest, err := sling.New().Post(url).BodyJSON(src).Request()
	if err != nil {
		return nil, fmt.Errorf("HTTP POST FAILED: %s", err)
	}

	for _, header := range headers {
		for k, v := range header {
			httpRequest.Header.Set(k, v)
		}
	}

	b, err := client.New(httpRequest).DoReq()
	if err != nil {
		return nil, fmt.Errorf("HTTP POST FAILED: %s", err)
	}

	if dst != nil {
		err = json.Unmarshal(b, dst)
		if err != nil {
			return nil, fmt.Errorf("HTTP POST FAILED - UNMARSHAL - %s: %s", err, string(b))
		}
		return nil, nil
	}

	return b, nil
}

func (client *Client) Delete(url string, src interface{}, headers ...map[string]string) ([]byte, error) {

	httpRequest, err := sling.New().Delete(url).BodyJSON(src).Request()
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		for k, v := range header {
			httpRequest.Header.Set(k, v)
		}
	}

	b, err := client.New(httpRequest).DoReq()
	if err != nil {
		return nil, err
	}

	return b, nil
}
