package cli

import (
	"flag"
	"fmt"
	"nowqs/frame/utils"
	"os"
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
	fmt.Println(os.Args[3:])
	searchType := utils.UTypeFolder
	folder, _ := os.Getwd()
	name := ""
	for _, val := range os.Args[3:] {
		// TODO: select
	}
	// searchType := flagSet.String("t", "folder", language.GetMsg("search type"))
	// dir := flagSet.String("d", "", language.GetMsg("search dir"))
	// name := flagSet.String("n", "", language.GetMsg("search name"))
	// flagSet.Parse(os.Args[3:])
	// if "" == *dir || !isExit(*dir) {
	// 	rootDir, _ := os.Getwd()
	// 	*dir = rootDir + "/"
	// }
	// fmt.Println(*searchType)
	// fmt.Println(*dir)
	// fmt.Println(*name)
	// result := utils.Search(*searchType, *dir, *name)

	// fmt.Println(language.GetMsg("search count") + ": " + strconv.Itoa(result.Num))
	// fmt.Println(language.GetMsg("search take time") + ": " + strconv.Itoa(int(result.TakeTime)))

}
