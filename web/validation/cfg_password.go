package validation

import (
	"unicode"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which checks for password
func PasswordWeak() *Config {

	min := PASSWORD_MIN_LENGTH
	max := STRING_MAX_LENGTH

	return NewConfig(
		"myPassword1",
		func (req web.RequestInterface, s string) (status *web.ResponseStatus, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				s,
			)
			if status != nil {
				return status, nil
			}

			if !verifyWeakPassword(s) { return req.Respond(400, "YOUR PASSWORD IS TOO WEAK"), nil }

			return nil, Hash256(s)
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, "PARAM IS NOT A STRING"), nil }

			status, s = checkString(
				req,
				min,
				max,
				s,
			)
			if status != nil {
				return status, nil
			}

            if !verifyWeakPassword(s) { return req.Respond(400, "YOUR PASSWORD IS TOO WEAK"), nil }

			return nil, Hash256(s)
		},
	)
}

// Returns a validation object which checks for password
func PasswordHard() *Config {

	min := float64(PASSWORD_MIN_LENGTH + 2)
	max := STRING_MAX_LENGTH

    return NewConfig(
        "^RT8t y8yt8T*^* $$%£^022",
        func (req web.RequestInterface, s string) (status *web.ResponseStatus, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				s,
			)
			if status != nil {
				return status, nil
			}

            if !verifyHardPassword(s) { return req.Respond(400, "YOUR PASSWORD IS TOO WEAK"), nil }

            return nil, Hash256(s)
        },
        func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

            s, ok := param.(string); if !ok { return req.Respond(400, "PARAM IS NOT A STRING"), nil }

			status, s = checkString(
				req,
				min,
				max,
				s,
			)
			if status != nil {
				return status, nil
			}

            if !verifyHardPassword(s) { return req.Respond(400, "YOUR PASSWORD IS TOO WEAK"), nil }

            return nil, Hash256(s)
        },
		min,
		max,
    )
}

func verifyHardPassword(input string) bool {

    var letters int
    var numbers int
    var number bool
    var special bool
    var upper bool

    for _, s := range input {

      switch {

        case unicode.IsNumber(s):
            number = true
            numbers++

        case unicode.IsUpper(s):
            upper = true
            letters++

        case unicode.IsPunct(s) || unicode.IsSymbol(s):
            special = true

        case unicode.IsLetter(s) || s == ' ':
            letters++

      }

    }

    return ((letters + numbers) >= PASSWORD_MIN_LENGTH) && number && special && upper
}

func verifyWeakPassword(input string) bool {

    var letters int
    var numbers int
    var other int

    for _, s := range input {

      switch {

        case unicode.IsNumber(s):
            numbers++

        case unicode.IsLetter(s) || s == ' ':
            letters++

        case unicode.IsPunct(s) || unicode.IsSymbol(s):
            other++

      }

    }

    return (letters + numbers + other) >= PASSWORD_MIN_LENGTH
}
