package errhandler

import (
	"fmt"
	"os"
)

func Fatal(err error, print bool) {
	if err != nil {
		if print {
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
