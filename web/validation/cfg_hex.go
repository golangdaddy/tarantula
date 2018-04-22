package validation

import (
	"encoding/hex"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a hex string with a length within optional range
func Hex(min, max float64) *Config {

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(param)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, param
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(s)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, s
		},
	)

	config.Min = min
	config.Max = max

	return config
}

// Returns a validation object that decodes a hex string with a length within optional range
func HexDecode(min, max float64) *Config {

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			b, err := hex.DecodeString(param)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, b
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			b, err := hex.DecodeString(s)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, b
		},
	)

	config.Min = min
	config.Max = max

	return config
}

// Returns a validation object that checks for a hex string with fixed range
func Hex256() *Config {

	max := float64(64)
	min := max

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(param)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, Sanitize(param)
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(s)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, Sanitize(s)
		},
	)

	config.Min = min
	config.Max = max

	return config
}
