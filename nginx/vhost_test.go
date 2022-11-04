package nginx

import "testing"

func TestDirIsEmpty(t *testing.T) {
	var expected bool = false
	exist, _ := isDir("/whatever")

	if exist != expected {
		t.Log("The non-existing directory should be", expected)
		t.Fail()
	}
}

func TestVHostEmpty(t *testing.T) {
	expected := make(map[string][]string, 2)
	expected[SitesAvailableDir] = []string{}
	expected[ConfdDir] = []string{}

	vhosts, _ := GetAllVHosts()

	if len(vhosts[SitesAvailableDir]) > 0 {
		t.Log(SitesAvailableDirPath, "configurations should be nil")
		t.Fail()
	}

	if len(vhosts[ConfdDir]) > 0 {
		t.Log(ConfdDirPath, "configurations should be nil")
		t.Fail()
	}
}
