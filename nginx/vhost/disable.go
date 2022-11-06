package vhost

import (
	"fmt"
	"nginxctl/helper"
	"nginxctl/nginx"
)

func DisableVHost(serverName string) {
	link := fmt.Sprintf("%s/%s", nginx.SitesEnabledDirPath, serverName)
	isLink, _ := helper.IsFileExist(link)
	if !isLink {
		err := fmt.Errorf(
			"Error: can not disabled %s, link file not found",
			serverName,
		)
		helper.ShowError(err)
	}
	err := helper.DeleteFile(link)
	if err != nil {
		helper.ShowError(err)
	}
	fmt.Println("Successfully disable", serverName)
	return
}
