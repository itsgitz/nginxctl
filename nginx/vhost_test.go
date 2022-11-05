package nginx

import "testing"

func TestDirIsEmpty(t *testing.T) {
	var expected bool = false
	exist, _ := isDir(SitesAvailableDirPath)

	if exist != expected {
		t.Log("The directory should be empty")
		t.Fail()
	}
}

func TestGetFilesEmpty(t *testing.T) {
	var expected []string

	files, _ := getFiles(SitesAvailableDirPath)

	if len(expected) != len(files) {
		t.Log("The files should be don't exist")
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
