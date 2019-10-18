package main

import (
	"fmt"
	"log"
	"os"

	"github.com/desktopgame/xstruct/xstruct"

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
	for _, v := range namespace.Map {
		fmt.Println(v.UniqueName)
	}
}
