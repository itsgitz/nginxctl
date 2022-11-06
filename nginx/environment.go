package nginx

import "fmt"

const (
	RootDir           string = "/etc/nginx"
	SitesAvailableDir string = "sites-available"
	SitesEnabledDir   string = "sites-enabled"
	ConfdDir          string = "conf.d"

	StaticSite   string = "static_site"
	ReverseProxy string = "reverse_proxy"

	TestEmptyFile string = "/empty"
)

var (
	SitesAvailableDirPath string = fmt.Sprintf("%s/%s", RootDir, SitesAvailableDir)
	SitesEnabledDirPath   string = fmt.Sprintf("%s/%s", RootDir, SitesEnabledDir)
	ConfdDirPath          string = fmt.Sprintf("%s/%s", RootDir, ConfdDir)
)
