package vhost

import (
	"fmt"
	"nginxctl/nginx"
	"nginxctl/nginx/vhost"
)

func List(onlyArgs map[string]bool) {
	vhosts, err := vhost.GetAllVHosts()
	if err != nil {
		fmt.Println(err)
		return
	}

	if onlyArgs[nginx.SitesAvailableDir] {
		// Only --sites-available
		show(vhosts[nginx.SitesAvailableDir], nginx.SitesAvailableDirPath)
	} else if onlyArgs[nginx.SitesEnabledDir] {
		// Only --sites-enabled
		show(vhosts[nginx.SitesEnabledDir], nginx.SitesEnabledDirPath)
	} else if onlyArgs[nginx.ConfdDir] {
		// Only --confd
		show(vhosts[nginx.ConfdDir], nginx.ConfdDirPath)
	} else {
		// All, means no argument specified
		show(vhosts[nginx.SitesAvailableDir], nginx.SitesAvailableDirPath)
		show(vhosts[nginx.SitesEnabledDir], nginx.SitesEnabledDirPath)
		show(vhosts[nginx.ConfdDir], nginx.ConfdDirPath)
	}
}

func show(vhosts []string, path string) {
	num := len(vhosts)
	fmt.Printf("Configurations on %s (total %d):\n", path, num)
	if num == 0 {
		fmt.Printf("* No config file\n\n")
		return
	}

	for _, v := range vhosts {
		fmt.Println("*", v)
	}
	fmt.Println()
	return
}
