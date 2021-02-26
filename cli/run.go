package cli

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	// flag.v
	flag.Parse()
}

// Run is cli funciton
func Run() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "version":
		version()
	case "help":
		help()
	case "file":
		file()
	default:
		fmt.Printf("%s", descAll)
	}
}
