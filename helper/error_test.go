package helper

import (
	"errors"
	"testing"
)

func TestShowError(t *testing.T) {
	var want error = errors.New("This is error")
	if want == nil {
		t.Log("This error should be not nil")
		t.Fail()
	}
}
