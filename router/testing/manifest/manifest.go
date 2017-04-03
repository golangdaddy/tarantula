package manifest

import (
    "fmt"
    "strings"
    "io/ioutil"
    "encoding/json"
    //
    "github.com/golangdaddy/tarantula/router/common"
)

type Manifest struct {
    Host string `json:"host"`
    Spec string `json:"spec"`
    Token string `json:"token"`
    Authorization string `json:"authorization"`
    Label string `json:"label"`
    Variables map[string]interface{} `json:"variables"`
    Endpoints []*Endpoint `json:"endpoints"`
}

func (manifest *Manifest) AddEndpoints(endpoints ...*Endpoint) {

    manifest.Endpoints = append(
        manifest.Endpoints,
        endpoints...,
    )
}

func (manifest *Manifest) NewEndpoint(label, method, path string) *Endpoint {

    return &Endpoint{
        manifest: manifest,
        Label: label,
        Method: method,
        Endpoint: path,
    }
}

type Endpoint struct {
    manifest *Manifest `json:"-"`
    Label string `json:label`
    Method string `json:"method"`
    Endpoint string `json:"endpoint"`
    PathArgs map[string]string `json:"pathArgs"`
    BodyArgs map[string]string `json:"bodyArgs"`
    BodyLiterals map[string]interface{} `json:"bodyLiterals"`
    Use map[string]string `json:"use"`
    Spec *common.HandlerSpec `json:"spec"`
    DelayBefore int `json:"delayBefore"`
    DelayAfter int `json:"delayAfter"`
}

func (endpoint *Endpoint) SecDelayBefore(x int) *Endpoint {
    endpoint.DelayBefore = x
    return endpoint
}

func (endpoint *Endpoint) SecDelayAfter(x int) *Endpoint {
    endpoint.DelayAfter = x
    return endpoint
}

func (endpoint *Endpoint) NewPathArgs(args map[string]string) *Endpoint {
    endpoint.PathArgs = args
    return endpoint
}

func (endpoint *Endpoint) NewBodyArgs(args map[string]string) *Endpoint {
    endpoint.BodyArgs = args
    return endpoint
}

func (endpoint *Endpoint) NewBodyLiterals(args map[string]interface{}) *Endpoint {
    endpoint.BodyLiterals = args
    return endpoint
}

func (endpoint *Endpoint) NewUsage(args map[string]string) *Endpoint {
    endpoint.Use = args
    return endpoint
}

func LoadManifest(path string) (*Manifest, error) {

        fmt.Println("LOADING MANIFEST", path)

        b, err := ioutil.ReadFile(path)
        if err != nil {
            return nil, err
        }

        s := strings.Replace(string(b), " ", "", -1)
        s = strings.Replace(s, "\n", "", -1)

        return ParseManifest([]byte(s))
}

func ParseManifest(b []byte) (*Manifest, error) {

        manifest := &Manifest{}
        err := json.Unmarshal(b, &manifest)
        if err != nil {
            return nil, err
        }

        return manifest, nil
}
