package nginx

import "testing"

func TestSitesAvailableDirPath(t *testing.T) {
	var expected string = "/etc/nginx/sites-available"
	if expected != SitesAvailableDirPath {
		t.Log("SitesAvailableDirPath should be /etc/nginx/sites-available")
		t.Fail()
	}
}

func TestSitesEnabledDitPath(t *testing.T) {
	var expected string = "/etc/nginx/sites-enabled"
	if expected != SitesEnabledDirPath {
		t.Log("SitesEnabledDirPath should be /etc/nginx/sites-enabled")
		t.Fail()
	}
}

func TestConfdDirPath(t *testing.T) {
	var expected string = "/etc/nginx/conf.d"
	if expected != ConfdDirPath {
		t.Log("ConfdDirPath should be /etc/nginx/conf.d")
		t.Fail()
	}
}
