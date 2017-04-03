package validation

import (
	"fmt"
	"strings"
	"strconv"
	//
	"github.com/golangdaddy/tarantula/web"
)

type IPv4 [4]int

func (ipv4 IPv4) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ipv4[0], ipv4[1], ipv4[2], ipv4[3])
}

// Returns a validation object that checks for a string with a length within optional range
func IPv4Address() *Config {

	min := 7.0
	max := 15.0

	return NewConfig(
		"127.0.0.1",
		func (req web.RequestInterface, ipAddress string) (*web.ResponseStatus, interface{}) {

            ipv4 := IPv4{}
			parts := strings.Split(ipAddress, ".")
			if len(parts) != 4 {
				return req.Respond(400, ERR_RANGE_EXCEED), nil
			}
            for x, s := range parts {
                i, err := strconv.Atoi(
                    strings.TrimSpace(s),
                )
                if err != nil {
                    return req.Respond(400, ERR_NOT_INT + ": " + err.Error()), nil
                }
                ipv4[x] = i
            }

			return nil, ipv4
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			ipAddress, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }


			ipv4 := IPv4{}
			parts := strings.Split(ipAddress, ".")
			if len(parts) != 4 {
				return req.Respond(400, ERR_RANGE_EXCEED), nil
			}
            for x, s := range parts {
                i, err := strconv.Atoi(
                    strings.TrimSpace(s),
                )
                if err != nil {
                    return req.Respond(400, ERR_NOT_INT + ": " + err.Error()), nil
                }
                ipv4[x] = i
            }

			return nil, ipv4
		},
		min,
		max,
	)
}
