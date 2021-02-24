package cli

import (
	"fmt"
	"nowqs/frame/config"
)

func version() {
	fmt.Println(config.Version)
}
