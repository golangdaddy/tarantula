package openapi

type APISpec struct {
	Swagger string `json:"swagger"`
	Info *Info `json:"info"`
	Host string `json:"host"`
	BasePath string `json:"basePath"`
	Schemes []string `json:"schemes"`
	Consumes []string `json:"consumes"`
	Produces []string `json:"produces"`
	Paths map[string]*Path `json:"paths"`
	Definitions map[string]*Definition `json:"definitions"`
	SecurityDefinitions map[string]*SecurityDefinition `json:"securityDefinitions"`
}

type SecurityDefinition struct {
	Type string `json:"type"`
}

type SecuritySchemeObject struct {
	Type []string `json:"type"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	In string `json:"in,omitempty"`
	Scheme []string `json:"scheme,omitempty"`
	BearerFormat []string `json:"bearerFormat,omitempty"`
	// oauth
	Flow string `json:"flow,omitempty"` // "implicit", "password", "application" or "accessCode"
	AuthorizationUrl string `json:"authorizationUrl,omitempty"` // "implicit", "password", "application" or "accessCode"
	TokenUrl string `json:"tokenUrl,omitempty"` // "implicit", "password", "application" or "accessCode"
	Scopes map[string]string `json:"scopes,omitempty"`
}

type Info struct {
	Version string `json:"version"`
	Title string `json:"title"`
	Description string `json:"description"`
	TermsOfService string `json:"termsOfService"`
	Contact struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		URL   string `json:"url"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

type Parameter struct {
	// required 'fixed' fields
	Name string `json:"name,omitempty"`
	// options: header, formData, query, path
	In string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required interface{} `json:"required,omitempty"`
	// if body
	Schema *Schema `json:"schema,omitempty"`
	// else all of the below
	Type string `json:"type,omitempty"`
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
	Format string `json:"format,omitempty"`
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`
	Items map[string]string `json:"items,omitempty"`
	CollectionFormat string `json:"collectionFormat,omitempty"`
//
	Default interface{} `json:"default,omitempty"`
	// String validations
	MaxLength *int64 `json:"maxLength,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	// "string", "number", "integer", "boolean", "array" or "file". If type is "file" see docs
}

// taken from https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#schemaObject
// bits stolen from https://github.com/go-swagger/go-swagger/blob/master/generator/structs.go
type Schema struct {
	Ref string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
	// misc
	Format string `json:"format,omitempty"`
	Required bool `json:"required,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	// String validations
	MinLength *int64 `json:"minLength,omitempty"`
	MaxLength *int64 `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	Default interface{} `json:"default,omitempty"`
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	// "string", "number", "integer", "boolean", or "array" etc
	Items *Items `json:"items,omitempty"`
}

type Items struct {
	// misc
	// "string", "number", "integer", "boolean", or "array" etc
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	// String validations
	MaxLength *int64 `json:"maxLength,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	Default interface{} `json:"default,omitempty"`
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	Items *Items `json:"items,omitempty"`
}

type Path struct {
	GET *PathMethod `json:"get,omitempty"`
	PUT *PathMethod `json:"put,omitempty"`
	POST *PathMethod `json:"post,omitempty"`
	PATCH *PathMethod `json:"patch,omitempty"`
	DELETE *PathMethod `json:"delete,omitempty"`
	HEAD *PathMethod `json:"head,omitempty"`
	OPTIONS *PathMethod `json:"options,omitempty"`
}

type PathMethod struct {
	Description string `json:"description,omitempty"`
	OperationID string `json:"operationId,omitempty"`
	Parameters []*Parameter `json:"parameters,omitempty"`
	Responses Responses `json:"responses,omitempty"`
	Produces []string `json:"produces,omitempty"`
	Security []*SecuritySchemeObject `json:"security,omitempty"`
}

type Definition struct {
	Type  string `json:"type"`
	Ref string   `json:"$ref,omitempty"`
	Required []string `json:"required,omitempty"`
	Properties map[string]Parameter `json:"properties,omitempty"`
}

type StatusSchema struct {
	Type string `json:"type"`
	Items map[string]string `json:"items,omitempty"`
}

type StatusCode struct {
	Description string `json:"description"`
	Schema StatusSchema `json:"schema"`
}

// Comments refer to the line above the comment :)
type Responses struct {
	Code100 *StatusCode `json:"100,omitempty"`
	// Continue
	Code101 *StatusCode `json:"101,omitempty"`
	// Switching Protocols
	Code102 *StatusCode `json:"102,omitempty"`
	// Processing (WebDAV)
	Code200 *StatusCode `json:"200,omitempty"`
	// OK
	Code201 *StatusCode `json:"201,omitempty"`
	// Created
	Code202 *StatusCode `json:"202,omitempty"`
	// Accepted
	Code203 *StatusCode `json:"203,omitempty"`
	// Non-Authoritative Information
	Code204 *StatusCode `json:"204,omitempty"`
	// No Content
	Code205 *StatusCode `json:"205,omitempty"`
	// Reset Content
	Code206 *StatusCode `json:"206,omitempty"`
	// Partial Content
	Code207 *StatusCode `json:"207,omitempty"`
	// Multi-Status (WebDAV)
	Code208 *StatusCode `json:"208,omitempty"`
	// Already Reported (WebDAV)
	Code226 *StatusCode `json:"226,omitempty"`
	// IM Used
	Code300 *StatusCode `json:"300,omitempty"`
	// Multiple Choices
	Code301 *StatusCode `json:"301,omitempty"`
	// Moved Permanently
	Code302 *StatusCode `json:"302,omitempty"`
	// Found
	Code303 *StatusCode `json:"303,omitempty"`
	// See Other
	Code304 *StatusCode `json:"304,omitempty"`
	// Not Modified
	Code305 *StatusCode `json:"305,omitempty"`
	// Use Proxy
	Code306 *StatusCode `json:"306,omitempty"`
	// (Unused)
	Code307 *StatusCode `json:"307,omitempty"`
	// Temporary Redirect
	Code308 *StatusCode `json:"308,omitempty"`
	// Permanent Redirect (experimental)
	Code400 *StatusCode `json:"400,omitempty"`
	// Bad Request
	Code401 *StatusCode `json:"401,omitempty"`
	// Unauthorized
	Code402 *StatusCode `json:"402,omitempty"`
	// Payment Required
	Code403 *StatusCode `json:"403,omitempty"`
	// Forbidden
	Code404 *StatusCode `json:"404,omitempty"`
	// Not Found
	Code405 *StatusCode `json:"405,omitempty"`
	// Method Not Allowed
	Code406 *StatusCode `json:"406,omitempty"`
	// Not Acceptable
	Code407 *StatusCode `json:"407,omitempty"`
	// Proxy Authentication Required
	Code408 *StatusCode `json:"408,omitempty"`
	// Request Timeout
	Code409 *StatusCode `json:"409,omitempty"`
	// Conflict
	Code410 *StatusCode `json:"410,omitempty"`
	// Gone
	Code411 *StatusCode `json:"411,omitempty"`
	// Length Required
	Code412 *StatusCode `json:"412,omitempty"`
	// Precondition Failed
	Code413 *StatusCode `json:"413,omitempty"`
	// Request Entity Too Large
	Code414 *StatusCode `json:"414,omitempty"`
	// Request-URI Too Long
	Code415 *StatusCode `json:"415,omitempty"`
	// Unsupported Media Type
	Code416 *StatusCode `json:"416,omitempty"`
	// Requested Range Not Satisfiable
	Code417 *StatusCode `json:"417,omitempty"`
	// Expectation Failed
	Code418 *StatusCode `json:"418,omitempty"`
	// I'm a teapot (RFC 2324)
	Code420 *StatusCode `json:"420,omitempty"`
	// Enhance Your Calm (Twitter)
	Code422 *StatusCode `json:"422,omitempty"`
	// Unprocessable Entity (WebDAV)
	Code423 *StatusCode `json:"423,omitempty"`
	// Locked (WebDAV)
	Code424 *StatusCode `json:"424,omitempty"`
	// Failed Dependency (WebDAV)
	Code425 *StatusCode `json:"425,omitempty"`
	// Reserved for WebDAV
	Code426 *StatusCode `json:"426,omitempty"`
	// Upgrade Required
	Code428 *StatusCode `json:"428,omitempty"`
	// Precondition Required
	Code429 *StatusCode `json:"429,omitempty"`
	// Too Many Requests
	Code431 *StatusCode `json:"431,omitempty"`
	// Request Header Fields Too Large
	Code444 *StatusCode `json:"444,omitempty"`
	// No Response (Nginx)
	Code449 *StatusCode `json:"449,omitempty"`
	// Retry With (Microsoft)
	Code450 *StatusCode `json:"450,omitempty"`
	// Blocked by Windows Parental Controls (Microsoft)
	Code451 *StatusCode `json:"451,omitempty"`
	// Unavailable For Legal Reasons
	Code499 *StatusCode `json:"499,omitempty"`
	// Client Closed Request (Nginx)
	Code500 *StatusCode `json:"500,omitempty"`
	// Internal Server Error
	Code501 *StatusCode `json:"501,omitempty"`
	// Not Implemented
	Code502 *StatusCode `json:"502,omitempty"`
	// Bad Gateway
	Code503 *StatusCode `json:"503,omitempty"`
	// Service Unavailable
	Code504 *StatusCode `json:"504,omitempty"`
	// Gateway Timeout
	Code505 *StatusCode `json:"505,omitempty"`
	// HTTP Version Not Supported
	Code506 *StatusCode `json:"506,omitempty"`
	// Variant Also Negotiates (Experimental)
	Code507 *StatusCode `json:"507,omitempty"`
	// Insufficient Storage (WebDAV)
	Code508 *StatusCode `json:"508,omitempty"`
	// Loop Detected (WebDAV)
	Code509 *StatusCode `json:"509,omitempty"`
	//  Bandwidth Limit Exceeded (Apache)
	Code510 *StatusCode `json:"510,omitempty"`
	// Not Extended
	Code511 *StatusCode `json:"511,omitempty"`
	// Network Authentication Required
	Code598 *StatusCode `json:"598,omitempty"`
	// Network read timeout error
}
