package cli

import (
	"flag"
	"os"
)

func init() {
	// flag.v
	flag.Parse()
}

// Run is cli funciton
func Run() {
	switch os.Args[1] {
	case "help":
		help()
	}
}
