package helper

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// isFileDir will check if the directory is exists
func IsFileExist(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// getFiles get filenames as slice
func GetFiles(path string) ([]string, error) {
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

func WriteConfig(path string, config []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(config)
	if err != nil {
		return err
	}

	return nil
}

func WriteLink(src, dst string) (string, error) {
	ln, err := exec.Command("ln", "-s", src, dst).Output()
	if err != nil {
		return "", err
	}
	return string(ln), err
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
