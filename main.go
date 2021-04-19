package main

// import "nowqs/frame/cli"

import (
	"fmt"
	"nowqs/frame/utils"
)

func main() {
	// cli.Run()
	a := utils.NewDefaultOptions()

	fmt.Printf("%v", a.New())
}
