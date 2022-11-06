package vhost

import (
	"nginxctl/helper"
	"nginxctl/nginx"
	"testing"
)

func TestDirIsEmpty(t *testing.T) {
	var want bool = false
	exist, _ := helper.IsFileExist(nginx.SitesAvailableDirPath)

	if exist != want {
		t.Log("The directory should be empty")
		t.Fail()
	}
}

func TestGetFilesEmpty(t *testing.T) {
	var want []string

	files, _ := helper.GetFiles(nginx.SitesAvailableDirPath)

	if len(want) != len(files) {
		t.Log("The files should be don't exist")
		t.Fail()
	}
}

func TestVHostEmpty(t *testing.T) {
	want := make(map[string][]string, 2)
	want[nginx.SitesAvailableDir] = []string{}
	want[nginx.ConfdDir] = []string{}

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
