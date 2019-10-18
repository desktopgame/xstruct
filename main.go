package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/desktopgame/xstruct/xstruct"

	"github.com/beevik/etree"
)

func writeProgram(buf *bytes.Buffer, class *xstruct.Class) {
	buf.WriteString("type ")
	buf.WriteString(class.UniqueName)
	buf.WriteString(" {\n")
	buf.WriteString("    // define attribute\n")
	for k, _ := range class.Attributes {
		buf.WriteString("    ")
		buf.WriteString("Attr")
		buf.WriteString(k)
		buf.WriteString(" string `xml:\"")
		buf.WriteString(k)
		buf.WriteString(",attr\"\n")
	}
	buf.WriteString("    // define subelement\n")
	for _, class := range class.InnerClasses {
		buf.WriteString("    ")
		buf.WriteString("Sub")
		buf.WriteString(class.SimpleName)
		buf.WriteString(" []*")
		buf.WriteString(class.UniqueName)
		buf.WriteString(" `xml:\"")
		buf.WriteString(class.SimpleName)
		buf.WriteString("\"\n")
	}
	buf.WriteString("    // define content\n")
	buf.WriteString("    Content string `xml:\",chardata\"`\n")
	buf.WriteString("}\n")
}

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
		writeProgram(&buf, class)
	}
	fmt.Println(buf.String())
}
