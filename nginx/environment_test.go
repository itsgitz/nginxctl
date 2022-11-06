package nginx

import "testing"

func TestSitesAvailableDirPath(t *testing.T) {
	var want string = "/etc/nginx/sites-available"
	if want != SitesAvailableDirPath {
		t.Log("SitesAvailableDirPath should be /etc/nginx/sites-available")
		t.Fail()
	}
}

func TestSitesEnabledDitPath(t *testing.T) {
	var want string = "/etc/nginx/sites-enabled"
	if want != SitesEnabledDirPath {
		t.Log("SitesEnabledDirPath should be /etc/nginx/sites-enabled")
		t.Fail()
	}
}

func TestConfdDirPath(t *testing.T) {
	var want string = "/etc/nginx/conf.d"
	if want != ConfdDirPath {
		t.Log("ConfdDirPath should be /etc/nginx/conf.d")
		t.Fail()
	}
}
