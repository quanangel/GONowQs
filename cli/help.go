package cli

import (
	"fmt"
	"nowqs/frame/language"
)

var descAll = []string{
	"this is NowQs cli system",
	"desc test",
	"Usage:",
	"	<run file name> <command> [arguments]",
	"The commoand are:",
	"	version		print the system version",
	"	help		print the desc",
	language.GetMsg("empty"),
}

func help() {
	for _, v := range descAll {
		fmt.Println(v)
	}
}
