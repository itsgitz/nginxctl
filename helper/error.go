package helper

import (
	"fmt"
	"os"
)

// ShowError help the application show the error message without `panic`
func ShowError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
