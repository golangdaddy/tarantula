package manifest

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type Manifest struct {
	Host string `json:"host"`
	Spec string `json:"spec"`
	Headers map[string]string `json:"headers"`
	Label string `json:"label"`
	Variables map[string]interface{} `json:"variables"`
	Endpoints []*Endpoint `json:"endpoints"`
}

func (manifest *Manifest) SetHeader(k, v string) {
	manifest.Headers[k] = v
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
