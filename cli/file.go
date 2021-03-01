package cli

import (
	"flag"
	"fmt"
	"nowqs/frame/language"
)

func file() {
	fileCli := flag.NewFlagSet("file", flag.ExitOnError)

	switch flag.Args()[1] {
	case "-s":
		searchCliShort(fileCli)
	case "-search":
		fmt.Println(2)
	default:
		fmt.Printf("%s\r\n", "help")
	}

	fileCli.Parse(flag.Args()[1:])
}

func searchCliShort(flagSet *flag.FlagSet) {
	dir := flagSet.String("d", "", language.GetMsg("lack search dir"))
	println(dir)
	fmt.Println(flagSet.Args())
	// TODO: not finish
}
