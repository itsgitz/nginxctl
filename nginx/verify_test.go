package nginx

import "testing"

func TestVerify(t *testing.T) {
	var expected error = nil

	verify := Verify()
	if verify == expected {
		t.Log("Error should be not nil")
		t.Fail()
	}
}
