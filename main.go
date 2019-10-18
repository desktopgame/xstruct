package main

import (
	"fmt"
	"log"

	"github.com/desktopgame/xstruct/xstruct"

	"github.com/beevik/etree"
)

func main() {
	fmt.Println("hello")
	edoc := etree.NewDocument()
	err := edoc.ReadFromFile("./testdata/Planet.vcxproj")
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
