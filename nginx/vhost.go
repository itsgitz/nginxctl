package nginx

import (
	"io/ioutil"
	"os"
)

// GetAllVHosts get all nginx virtual host configurations on the current machine
func GetAllVHosts() (map[string][]string, error) {
	// Check if `sites-available` directory is exists
	isAvailableDir, _ := isDir(SitesAvailableDirPath)

	// Check if `conf.d` directory is exists
	isConfdDir, _ := isDir(ConfdDirPath)

	vhosts := make(map[string][]string, 2)
	if isAvailableDir {
		available, err := ioutil.ReadDir(SitesAvailableDirPath)
		if err != nil {
			return nil, err
		}

		for _, file := range available {
			vhosts[SitesAvailableDir] = append(vhosts[SitesAvailableDir], file.Name())
		}
	}

	if isConfdDir {
		confd, err := ioutil.ReadDir(ConfdDirPath)
		if err != nil {
			return nil, err
		}

		for _, file := range confd {
			vhosts[ConfdDir] = append(vhosts[ConfdDir], file.Name())
		}
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
