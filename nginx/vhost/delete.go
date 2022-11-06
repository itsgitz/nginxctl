package vhost

import (
	"fmt"
	"nginxctl/helper"
	"nginxctl/nginx"
)

func DeleteVHost(serverName string) {
	file := fmt.Sprintf("%s/%s", nginx.SitesAvailableDirPath, serverName)
	link := fmt.Sprintf("%s/%s", nginx.SitesEnabledDirPath, serverName)

	// Don't delete if file is not exist
	isFile, _ := helper.IsFileExist(file)
	if !isFile {
		err := fmt.Errorf(
			"Error: unable to delete virtualhost configuration on %s, file not found",
			file,
		)
		helper.ShowError(err)
	}

	// Check if link is exist too
	isLink, _ := helper.IsFileExist(link)
	// If link found, then remove the link first
	if isLink {
		err := helper.DeleteFile(link)
		if err != nil {
			helper.ShowError(err)
		}
		fmt.Println("Successfully delete configuration on", link)
	}

	// if not found, only delete configuration on /etc/nginx/sites-available
	err := helper.DeleteFile(file)
	if err != nil {
		helper.ShowError(err)
	}
	fmt.Println("Successfully deleted configuration on", file)
	return
}
