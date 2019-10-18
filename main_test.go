package main

import (
	"fmt"
	"testing"

	"github.com/desktopgame/xstruct/xstruct/detail"
)

func TestConfig(t *testing.T) {
	buf := detail.CreateProgram("./testdata/Planet.vcxproj")
	fmt.Println(buf.String())
}
