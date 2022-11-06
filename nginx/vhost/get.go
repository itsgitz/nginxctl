package vhost

import (
	"nginxctl/helper"
	"nginxctl/nginx"
)

// GetAllVHosts get all nginx virtual host configurations on the current machine
func GetAllVHosts() (map[string][]string, error) {
	// Check if `sites-available` directory is exists
	isAvailableDir, _ := helper.IsFileExist(nginx.SitesAvailableDirPath)

	// Check if `conf.d` directory is exists
	isConfdDir, _ := helper.IsFileExist(nginx.ConfdDirPath)

	// Check if `sites-enabled` directory is exists
	isEnabledDir, _ := helper.IsFileExist(nginx.SitesEnabledDirPath)

	vhosts := make(map[string][]string, 3)

	// List for vhosts in `sites-available` directory
	if isAvailableDir {
		files, err := helper.GetFiles(nginx.SitesAvailableDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.SitesAvailableDir] = files
	}

	// List for vhosts in `conf.d` directory
	if isConfdDir {
		files, err := helper.GetFiles(nginx.ConfdDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.ConfdDir] = files
	}

	// List for vhosts in `sites-enabled` directory
	if isEnabledDir {
		files, err := helper.GetFiles(nginx.SitesEnabledDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.SitesEnabledDir] = files
	}

	return vhosts, nil
}
