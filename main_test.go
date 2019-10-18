package main

import (
	"fmt"
	"testing"

	"github.com/desktopgame/xstruct/xstruct/detail"
)

func TestConfig(t *testing.T) {
	buf := detail.CreateProgram("./testdata/Planet.vcxproj", detail.Option{
		Prefix: "",
		Suffix: "",
	})
	fmt.Println(buf.String())
}

func TestFilter(t *testing.T) {
	buf := detail.CreateProgram("./testdata/Planet.vcxproj.filters", detail.Option{
		Prefix: "",
		Suffix: "",
	})
	fmt.Println(buf.String())
}

func TestUser(t *testing.T) {
	buf := detail.CreateProgram("./testdata/Planet.vcxproj.user", detail.Option{
		Prefix: "",
		Suffix: "",
	})
	fmt.Println(buf.String())
}

func TestPackage(t *testing.T) {
	buf := detail.CreateProgram("./testdata/packages.config", detail.Option{
		Prefix: "Prefix",
		Suffix: "",
	})
	fmt.Println(buf.String())
}
