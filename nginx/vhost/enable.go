package vhost

import (
	"fmt"
	"nginxctl/helper"
	"nginxctl/nginx"
)

func EnableVHost(serverName string) {
	file := fmt.Sprintf("%s/%s", nginx.SitesAvailableDirPath, serverName)
	link := fmt.Sprintf("%s/%s", nginx.SitesEnabledDirPath, serverName)
	isLink, _ := helper.IsFileExist(link)

	// We cannot enabled the virtualhost if link file already exist
	if isLink {
		err := fmt.Errorf(
			"Error: can not enabled %s, the virtualhost is already enabled",
			serverName,
		)
		helper.ShowError(err)
	}

	// We cannot enabled the virtualhost if virtualhost is not exist
	isVHost, _ := helper.IsFileExist(file)
	if !isVHost {
		err := fmt.Errorf(
			"Error: can not enabled %s, the virtualhost is not exist",
			serverName,
		)
		helper.ShowError(err)
	}

	src := fmt.Sprintf("%s/%s", nginx.SitesAvailableDirPath, serverName)
	_, err := helper.WriteLink(src, nginx.SitesEnabledDirPath)
	if err != nil {
		helper.ShowError(err)
	}
	fmt.Println("Successfully enabled", serverName)
	return
}
