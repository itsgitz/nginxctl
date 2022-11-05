package vhost

import (
	"io/ioutil"
	"nginxctl/nginx"
	"os"
)

// GetAllVHosts get all nginx virtual host configurations on the current machine
func GetAllVHosts() (map[string][]string, error) {
	// Check if `sites-available` directory is exists
	isAvailableDir, _ := isDir(nginx.SitesAvailableDirPath)

	// Check if `conf.d` directory is exists
	isConfdDir, _ := isDir(nginx.ConfdDirPath)

	// Check if `sites-enabled` directory is exists
	isEnabledDir, _ := isDir(nginx.SitesEnabledDirPath)

	vhosts := make(map[string][]string, 3)

	// List for vhosts in `sites-available` directory
	if isAvailableDir {
		files, err := getFiles(nginx.SitesAvailableDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.SitesAvailableDir] = files
	}

	// List for vhosts in `conf.d` directory
	if isConfdDir {
		files, err := getFiles(nginx.ConfdDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.ConfdDir] = files
	}

	// List for vhosts in `sites-enabled` directory
	if isEnabledDir {
		files, err := getFiles(nginx.SitesEnabledDirPath)
		if err != nil {
			return nil, err
		}
		vhosts[nginx.SitesEnabledDir] = files
	}

	return vhosts, nil
}

// isDir will check if the directory is exists
func isDir(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// getFiles get filenames as slice
func getFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var list []string
	for _, f := range files {
		list = append(list, f.Name())
	}

	return list, nil
}
