package httpclient

import (
	"net/http"
	"encoding/json"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	//
	"github.com/dghubble/sling"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		&http.Client{},
	}
}

func NewUrlfetchClient(ctx context.Context) *Client {

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

	b, err := client.New(httpRequest).Do()
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

	httpRequest, err := sling.New().Get(url).BodyJSON(src).Request()
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		for k, v := range header {
			httpRequest.Header.Set(k, v)
		}
	}

	b, err := client.New(httpRequest).Do()
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

func (client *Client) Delete(url string, src interface{}, headers ...map[string]string) ([]byte, error) {

	httpRequest, err := sling.New().Get(url).BodyJSON(src).Request()
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		for k, v := range header {
			httpRequest.Header.Set(k, v)
		}
	}

	b, err := client.New(httpRequest).Do()
	if err != nil {
		return nil, err
	}

	return b, nil
}
