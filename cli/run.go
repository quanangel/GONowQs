package cli

import (
	"flag"
	"fmt"
	"nowqs/frame/language"
	"os"
)

func init() {
	// flag.v
	flag.Parse()
}

var descAll = []string{
	"this is cli system",
	"desc test",
	language.GetMsg("empty"),
}

// Run is cli funciton
func Run() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		help()
	default:
		fmt.Printf("%s", descAll)
	}
}
