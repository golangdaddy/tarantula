package router

import (
	"time"
	"testing"
)

const (
	CONST_HOPEFULLYUNUSEDPORT = 3543
)

func TestServe(t *testing.T) {

	_, router := NewRouter("localhost")


	failChan := make(chan error)

	go func () {
		if err := router.Serve(CONST_HOPEFULLYUNUSEDPORT); err != nil {
			failChan <- err
		} else {
			failChan <- nil
		}
	}()

	time.Sleep(2 * time.Second)

	select {

	case hasError := <- failChan:

			t.Error(hasError)

		default:

	}
}
