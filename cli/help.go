package cli

import (
	"flag"
	"fmt"
)

func help() {
	// helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	fmt.Printf("args=%s\r\n", flag.Args())
}
