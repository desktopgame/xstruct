package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/desktopgame/xstruct/xstruct/detail"
)

func main() {
	var (
		prefix = flag.String("prefix", "", "prefix")
		suffix = flag.String("suffix", "", "suffix")
	)
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("not enough argument")
	}
	buf := detail.CreateProgram(flag.Arg(0), detail.Option{
		Prefix: *prefix,
		Suffix: *suffix,
	})
	fmt.Println(buf.String())
}
