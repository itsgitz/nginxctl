package vhost

import (
	"nginxctl/nginx"
	"testing"
)

func TestDirIsEmpty(t *testing.T) {
	var expected bool = false
	exist, _ := isDir(nginx.SitesAvailableDirPath)

	if exist != expected {
		t.Log("The directory should be empty")
		t.Fail()
	}
}

func TestGetFilesEmpty(t *testing.T) {
	var expected []string

	files, _ := getFiles(nginx.SitesAvailableDirPath)

	if len(expected) != len(files) {
		t.Log("The files should be don't exist")
		t.Fail()
	}
}

func TestVHostEmpty(t *testing.T) {
	expected := make(map[string][]string, 2)
	expected[nginx.SitesAvailableDir] = []string{}
	expected[nginx.ConfdDir] = []string{}

	vhosts, _ := GetAllVHosts()

	if len(vhosts[nginx.SitesAvailableDir]) > 0 {
		t.Log(nginx.SitesAvailableDirPath, "configurations should be nil")
		t.Fail()
	}

	if len(vhosts[nginx.ConfdDir]) > 0 {
		t.Log(nginx.ConfdDirPath, "configurations should be nil")
		t.Fail()
	}
}
