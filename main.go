package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/desktopgame/xstruct/xstruct"
	"github.com/desktopgame/xstruct/xstruct/detail"

	"github.com/beevik/etree"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough argument")
	}
	edoc := etree.NewDocument()
	err := edoc.ReadFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	namespace := &xstruct.Namespace{
		Map: make(map[string]*xstruct.Class),
	}
	sc := xstruct.XMLToScopeTree(edoc.Root())
	xstruct.DefineClassTree(namespace, sc)
	_, err = xstruct.DefineClassA(namespace, sc.ToPath())
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	for _, class := range namespace.Map {
		detail.WriteProgram(&buf, class)
	}
	fmt.Println(buf.String())
}
