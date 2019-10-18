package main

import (
	"fmt"
	"log"
	"os"

	"github.com/desktopgame/xstruct/xstruct/detail"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough argument")
	}
	buf := detail.CreateProgram(os.Args[1], detail.Option{
		Prefix: "",
		Suffix: "",
	})
	fmt.Println(buf.String())
}
