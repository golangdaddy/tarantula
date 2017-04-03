package girlfriend

import (
  "testing"
  //
  "github.com/golangdaddy/girlfriend/common"
)

func TestDuplicateParamPanic(t *testing.T) {
  defer func() {
      if r := recover(); r != nil {
          t.Log("Duplicate param panic works!")
      }
  }()

  root, _ := NewRouter("host")

  root.Config.SetHandlerRegistry(dup_param_panic_reg())

  foo := root.Add("foo")
    foo.Param(common.CountryISO2(), "countryCode").GET("1")
    foo.Param(common.String(), "id").GET("2")

  t.Log("Duplicate param panic doesn't work.")
  t.Fail()
}

func dup_param_panic_reg() common.Registry {
  return common.Registry{
    "1": func (req common.RequestInterface) *common.ResponseStatus {
      return common.Respond([]byte("foo"))
    },
    "2": func (req common.RequestInterface) *common.ResponseStatus {
      return common.Respond([]byte("bar"))
    },
  }
}
