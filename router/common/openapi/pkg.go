package openapi

func NewSpec(host, serviceName string) *APISpec {
	return &APISpec{
		Swagger: "2.0",
		Info: &Info{
			Title: serviceName,
		},
		Host: host,
		Paths: map[string]*Path{},
		Schemes: []string{"http"},
		Consumes: []string{"application/json"},
		Produces: []string{"application/json"},
		Definitions: map[string]*Definition{},
	}
}
