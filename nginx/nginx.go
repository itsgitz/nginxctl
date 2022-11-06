package nginx

import "os/exec"

func NginxTestConfiguration() (string, error) {
	nginx, err := exec.Command("nginx", "-t").Output()
	if err != nil {
		return "", err
	}
	return string(nginx), nil
}
