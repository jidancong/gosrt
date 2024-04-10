package main

import (
	"fmt"

	"github.com/jidancong/gosrt/cmd"
)

func main() {
	if err := cmd.Cmd(); err != nil {
		fmt.Println("error:", err)
	}
}
