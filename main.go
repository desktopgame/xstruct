package main

import (
	"fmt"
	"log"

	"github.com/beevik/etree"
)

func main() {
	fmt.Println("hello")
	edoc := etree.NewDocument()
	err := edoc.ReadFromFile("./testdata/Planet.vcxproj")
	if err != nil {
		log.Fatal(err)
	}
	str, _ := edoc.WriteToString()
	fmt.Println(str)
}
