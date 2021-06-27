package lib

import (
	"fmt"
	"os"
)

func Abort(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	os.Exit(1)
}
