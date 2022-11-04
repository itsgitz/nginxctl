package nginx

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func Verify() error {
	// This application cannot run on non Linux machine
	if runtime.GOOS != "linux" {
		return errors.New("Error: nginxctl currently can only run on Linux like machine")
	}

	// Make sure that nginx application is installed
	_, err := exec.Command("nginx", "-v").Output()
	if err != nil {
		return fmt.Errorf("Nginx is not installed: %v", err)
	}

	return nil
}
