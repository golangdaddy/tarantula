package validation

import 	(
	"regexp"
	"reflect"
	"strings"
	"net/url"
	"encoding/hex"
	//
	"golang.org/x/crypto/sha3"
	//
	"github.com/golangdaddy/tarantula/web"
)

func IsAlpha(s string) (bool, string) {
	return regexp.MustCompile("[a-zA-Z]").MatchString(s), strings.ToLower(s)
}

func IsAlphanumeric(s string) (bool, string) {
	return regexp.MustCompile("[a-zA-Z0-9_]").MatchString(s), strings.ToLower(s)
}

func Hash256(input string) string {

	b := make([]byte, 64)

	sha3.ShakeSum256(b, []byte(input))

	return hex.EncodeToString(b)
}

type BodyValidationFunction func (web.RequestInterface, interface{}) (*web.ResponseStatus, interface{})
type PathValidationFunction func (web.RequestInterface, string) (*web.ResponseStatus, interface{})

type Spec struct {
	Type string
	Keys []string
}

type Config struct {
	Model interface{} `json:"model"`
	Type string `json:"type"`
	PathFunction PathValidationFunction `json:"-"`
	BodyFunction BodyValidationFunction `json:"-"`
	Keys []string `json:"-"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	SummaryValue string `json:"summary"`
	DefaultValue interface{} `json:"default"`
	DescriptionValue string `json:"description"`
}

// Adds a default value to the config
func (vc *Config) Default(x interface{}) *Config {
	if reflect.TypeOf(x).String() == vc.Type {
		vc.DefaultValue = x
	} else {
		panic("Wrong type '" + reflect.TypeOf(x).String() + "' for default value '" + vc.Type + "'")
	}
	return vc
}

// Adds a description to the config
func (vc *Config) Description(x string) *Config {
	vc.DescriptionValue = x
	return vc
}

// Adds a description to the config
func (vc *Config) Summary(x string) *Config {
	vc.SummaryValue = x
	return vc
}

func (vc *Config) Key() string {

	return vc.Keys[0]
}

func (vc *Config) KeyJoin(delim string) string {

	return strings.Join(vc.Keys, delim)
}

func (vc *Config) Expecting() string {

	return "expecting: " + vc.Type + " for keys: " + vc.KeyJoin(", ")
}

func NewConfig(validationType interface{}, pathFunction PathValidationFunction, bodyFunction BodyValidationFunction, ranges ...float64) *Config {

	cfg := &Config{
		Model: validationType,
		Type: reflect.TypeOf(validationType).String(),
		PathFunction: pathFunction,
		BodyFunction: bodyFunction,
	}

	switch len(ranges) {

		case 2:

			cfg.Min = ranges[0]
			cfg.Max = ranges[1]

	}

	return cfg
}

type JSON struct {}

// Returns a validation object which checks for (in)valid json
func Json() *Config {

	return NewConfig(
		JSON{},
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			return nil, param
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			return nil, s
		},
	)
}

type URL struct{}

// Returns a validation object which checks for valid url
func Url() *Config {

	return NewConfig(
		"",
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

      		param = strings.TrimSpace(param)

			_, err := url.ParseRequestURI(param); if err != nil { return req.Respond(400, ERR_PARSE_URL), "" }

			return nil, param
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

      		s = strings.TrimSpace(strings.ToLower(s))

			_, err := url.ParseRequestURI(s); if err != nil { return req.Respond(400, ERR_PARSE_URL), "" }

			return nil, s
		},
	)
}

// Returns a validation object for request body that checks a property to see if it's an array
func StringInterfaceArray() *Config {

	return NewConfig(
		[]string{},
		nil,
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			a, ok := param.([]interface{}); if !ok { return req.Respond(400, ERR_NOT_ARRAY), nil }

			list := make([]string, len(a))

			for i, x := range a {

				list[i], ok = x.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			}

			return nil, list
		},
	)
}
