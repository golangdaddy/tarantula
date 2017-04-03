package common

import (
    "fmt"
    "errors"
    "testing"
    //
    "github.com/golangdaddy/tarantula/web"
    "github.com/golangdaddy/tarantula/web/validation"
)

type Test struct {
    Method string
    Segments []Segment
    Handler func (req web.RequestInterface) *web.ResponseStatus
    Expected bool
}

func (test *Test) Path() string {
    path := ""
    for _, item := range test.Segments {
        path += "/" + item.Value
    }
    return path
}

type Segment struct {
    Value string
    ParamName string
    Validation *validation.Config
}

func TestNode(t *testing.T) {

    handler_default := func (req web.RequestInterface) *web.ResponseStatus {
        if req.FullPath() != req.Param("$fullpath").(string) {
            return req.Fail()
        }
        return nil
    }

    tests := []*Test{
        &Test{
            "GET",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
            },
            handler_default,
            true,
        },
        &Test{
            "POST",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
            },
            handler_default,
            true,
        },
        &Test{
            "GET",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
                Segment{
                    "blah",
                    "",
                    nil,
                },
            },
            handler_default,
            true,
        },
        &Test{
            "GET",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
                Segment{
                    "ale",
                    "username",
                    validation.Username(5, 16),
                },
            },
            func (req web.RequestInterface) *web.ResponseStatus {

                _, ok := req.Param("username").(string)
                if !ok {
                    return req.Fail()
                }

                return nil
            },
            false,
        },
        &Test{
            "GET",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
                Segment{
                    "alex",
                    "username",
                    validation.Username(4, 16),
                },
            },
            func (req web.RequestInterface) *web.ResponseStatus {

                _, ok := req.Param("username").(string)
                if !ok {
                    return req.Fail()
                }

                return nil
            },
            true,
        },
        &Test{
            "GET",
            []Segment{
                Segment{
                    "blah",
                    "",
                    nil,
                },
                Segment{
                    "alex",
                    "username",
                    validation.Username(4, 16),
                },
            },
            func (req web.RequestInterface) *web.ResponseStatus {

                _, ok := req.Param("username").(string)
                if !ok {
                    return req.Fail()
                }

                return nil
            },
            true,
        },
    }

    for _, test := range tests {

        req := web.NewTestInterface(test.Method, test.Path())

        fmt.Println("TESTING:", test, test.Path(), req.FullPath())

        root := newNode()
        node := root

        for _, segment := range test.Segments {
            if segment.Validation != nil {
                node = node.Param(segment.Validation, segment.ParamName)
            } else {
                node = node.Add(segment.Value)
            }
        }

        switch test.Method {
            case "GET":
                node.GET(test.Handler)
            case "POST":
                node.POST(test.Handler)
            case "DELETE":
                node.DELETE(test.Handler)
        }

        req.SetParam("$fullpath", test.Path())

        status := root.MainHandler(req, test.Path())
        if test.Expected {
            if status != nil {
                t.Error(status.Message)
                return
            }
        } else {
            if status == nil {
                t.Error(errors.New("TEST SUCCEEDED WHEN IT SHOULD HAVE FAILED!"))
                return
            } else {
                fmt.Println(status.Message)
            }
        }

    }

}
