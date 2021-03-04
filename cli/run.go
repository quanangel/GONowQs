package cli

import (
	"flag"
	"fmt"
	"os"
)

// init is initiation function
func init() {
	flag.Parse()
}

// Run is cli funciton
func Run() {
	if len(os.Args) == 1 {
		runExit(1)
	}
	switch os.Args[1] {
	case "version":
		version()
	case "help":
		help()
	case "file":
		file()
	case "web":
		web()
	default:
		fmt.Printf("%s", descAll)
	}
}

// runExit is exit function
func runExit(code int) {
	help()
	os.Exit(1)
}

func isExit(path string) bool {
	_, err := os.Stat(path)
	return nil == err || os.IsExist(err)
}
