package helper

import (
	"nginxctl/nginx"
	"testing"
)

func TestIsFileNotExist(t *testing.T) {
	var want bool = false

	// for local test
	fileExist, err := IsFileExist(nginx.TestEmptyFile)
	if err == nil {
		t.Log("This should be error because the file is not exist")
		t.Fail()
	}

	if fileExist != want {
		t.Log("This should be false because the file is not exist")
		t.Fail()
	}
}

func TestGetFilesEmpty(t *testing.T) {
	var want []string

	files, err := GetFiles(nginx.TestEmptyFile)
	if err == nil {
		t.Log("This should be error because file is not exist")
		t.Fail()
	}

	if len(files) != len(want) {
		t.Log("This should be error because file is not exist")
		t.Fail()
	}
}
