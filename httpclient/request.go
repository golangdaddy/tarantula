package httpclient

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"io/ioutil"
)

type Request struct {
	*Client
	*http.Request
	Retries int
	Success string
}

func (req *Request) WithRetry(retry int) *Request {
	req.Retries += retry
	return req
}

func (req *Request) WithStatus(successStatusCodes string) *Request {
	req.Success = strings.ToLower(successStatusCodes)
	return req
}

func (req *Request) Do() ([]byte, error) {

	var err error

	for x := 0; x < req.Retries; x++ {

		var b []byte

		b, err = req.actuallyDo()
		if err == nil {
			return b, err
		}

	}

	return nil, err
}

func (req *Request) actuallyDo() ([]byte, error) {

	resp, err := req.httpClient.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf(
			"HTTP %s: FAILED: %s",
			req.Method,
			err,
		)
	}

	successCode, err := strconv.Atoi(req.Success)
	if err != nil {

		successPattern, err := strconv.Atoi(string(req.Success[0]))
		if err != nil {
			return nil, fmt.Errorf(
				"HTTP %s: FAILED TO INTERPRET YOUR SUCCESS PATTERN: %s",
				req.Method,
				req.Success,
			)
		}

		// check if 404 falls within 4XX pattern
		if resp.StatusCode < successPattern * 100 && resp.StatusCode >= (successPattern + 1) * 100 {
			return nil, fmt.Errorf(
				"HTTP %s: RESULT %d != EXPECTED %dXX - %s",
				req.Method,
				resp.StatusCode,
				successPattern,
				resp.Status,
			)
		}

	}

	if resp.StatusCode != successCode {
		return nil, fmt.Errorf(
			"HTTP %s: RESULT %d != EXPECTED %d - %s",
			req.Method,
			resp.StatusCode,
			successCode,
			resp.Status,
		)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("GET: RESPONSE BODY IS NIL")
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return b, nil
}
