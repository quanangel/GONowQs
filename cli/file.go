package cli

import (
	"flag"
	"fmt"
	"os"
)

func file() {
	fileCli := flag.NewFlagSet("file", flag.ExitOnError)
	fileSearch := fileCli.Bool("search", true, "search file")
	fmt.Println(*fileSearch)
	fmt.Println(os.Args[2])
}
