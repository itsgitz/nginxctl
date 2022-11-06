package nginx

import "testing"

func TestVerify(t *testing.T) {
	var want error = nil

	verify := Verify()
	if verify == want {
		t.Log("Error should be nil")
		t.Fail()
	}
}
