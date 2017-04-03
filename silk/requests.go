package silk

import 	(
		"time"
		"bytes"
		"strings"
		"strconv"
		"net/http"
		"io/ioutil"
		"encoding/json"
		//
		"github.com/golangdaddy/tarantula/web"
		)

// SendRequests makes get requests and returns success/fail and a byte array on success
func SendRequests(req web.RequestInterface, fn, method, url string, retries int, data interface{}) (*web.ResponseStatus, []byte) {
	fn += "/SendRequests/"+method

  var err error
	b := []byte{}

	method = strings.ToUpper(method)

	if method == "POST" && data != nil {
	  
	  b, err = json.Marshal(data); if err != nil { return req.Respond(500, err.Error()), nil }
	  
	}

	status, responseBytes := DoRequest(req, fn, method, url, bytes.NewBuffer(b)); if status == nil { return nil, responseBytes }

	// retry until max retry limit
	for i := 0; i < retries; i++ {

		time.Sleep(time.Second / 4)

		status, responseBytes = DoRequest(req, fn, method, url, bytes.NewBuffer(b)); if status != nil { continue }
	
	  break
	}

  return status, responseBytes
}

func DoRequest(req web.RequestInterface, fn, method, url string, payload *bytes.Buffer) (*web.ResponseStatus, []byte) {
	fn += "/DoRequest"

	request, err := http.NewRequest(method, url, payload); if err != nil { return req.Respond(500, err.Error()), nil }
	request.Header.Set("Content-Type", "application/json")

  c := &http.Client{}

	resp, err := c.Do(request); if err != nil { return req.Respond(500, err.Error()), nil }

	if resp == nil || resp.Body == nil { return req.Respond(500, "BAD RESPONSE BODY"), nil }

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil { return req.Respond(500, err.Error()), nil }

	if resp.StatusCode != 200 {
	  
	  req.Log().NewError(method + ": " + url)
	  
	  return req.Respond(500, "OUTBOUND "+method+" REQUEST FAILED WITH: "+strconv.Itoa(resp.StatusCode)+" - "+resp.Status), nil
	  
	}

	return nil, b
}
