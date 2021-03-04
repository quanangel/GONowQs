package cli

import (
	"flag"
	"fmt"
	"nowqs/frame/language"
	"nowqs/frame/utils"
	"os"
	"strconv"
)

func file() {
	if 1 >= len(flag.Args()) {
		runExit(1)
	}
	fileCli := flag.NewFlagSet("file", flag.ExitOnError)
	switch flag.Args()[1] {
	case "-s":
		searchCliShort(fileCli)
	case "-search":
		fmt.Println(2)
	default:
		fmt.Printf("%s\r\n", "help")
	}

}

func searchCliShort(flagSet *flag.FlagSet) {
	dir := flagSet.String("d", "", language.GetMsg("search dir"))
	name := flagSet.String("n", "", language.GetMsg("search name"))
	flagSet.Parse(os.Args[3:])
	fmt.Println(*name)
	if *name == "" {
		fmt.Println(language.GetMsg("lack search name"))
		os.Exit(1)
	}

	if "" == *dir || !isExit(*dir) {
		rootDir, _ := os.Getwd()
		*dir = rootDir + string(os.PathSeparator)
	}
	result := utils.Search(*dir, *name)

	fmt.Println(language.GetMsg("search count") + ": " + strconv.Itoa(result.Num))
	fmt.Print(language.GetMsg("search take time") + ":")
	fmt.Println(result.TakeTime)
	if 0 < result.Num {
		for _, val := range result.Data {
			fmt.Println(val)
		}
	}
}
