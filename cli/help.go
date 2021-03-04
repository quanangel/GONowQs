package cli

import (
	"fmt"
	"nowqs/frame/language"
)

var descAll = []string{
	language.GetMsg("this is NowQs cli"),
	language.GetMsg("usage") + ":",
	"	<" + language.GetMsg("run file name") + "> <" + language.GetMsg("command") + "> [" + language.GetMsg("arguments") + "]",
	language.GetMsg("the command list") + ":",
	"	file		" + language.GetMsg("show about file"),
	"	web		    " + language.GetMsg("run about web"),
	"	version		" + language.GetMsg("show the system version"),
	"	help		" + language.GetMsg("show the desc"),
}

func help() {
	for _, v := range descAll {
		fmt.Println(v)
	}
}
