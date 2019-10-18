package detail

import (
	"bytes"
	"log"

	"github.com/beevik/etree"
	"github.com/desktopgame/xstruct/xstruct"
)

func CreateProgram(path string) bytes.Buffer {
	edoc := etree.NewDocument()
	err := edoc.ReadFromFile(path)
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
		WriteClassDef(&buf, class)
	}
	root, err := xstruct.DefineClassA(namespace, sc.ToPath())
	if err != nil {
		log.Fatal(err)
	}
	WriteFuncDef(&buf, root)
	return buf
}

func WriteFuncDef(buf *bytes.Buffer, class *xstruct.Class) {
	// LoadFunc
	buf.WriteString("func Load")
	buf.WriteString(class.UniqueName)
	buf.WriteString("(path string) ")
	buf.WriteString(class.UniqueName)
	buf.WriteString(" {")
	buf.WriteString(`
    xmlFile, err := os.Open(path)
    if err != nil {
    	return nil, err
    }
    defer xmlFile.Close()
    xmlData, err := ioutil.ReadAll(xmlFile)
    if err != nil {
    	return nil, err
    }`)
	buf.WriteString("\n")
	buf.WriteString("    var data ")
	buf.WriteString(class.UniqueName)
	buf.WriteString(`
    xml.Unmarshal(xmlData, &data)
    return &data, nil`)
	buf.WriteString("\n}")
	buf.WriteString("\n")
	buf.WriteString("\n")
	// SaveFunc
	buf.WriteString("func Save")
	buf.WriteString(class.UniqueName)
	buf.WriteString("(path string, data *")
	buf.WriteString(class.UniqueName)
	buf.WriteString(", perm uint32) error")
	buf.WriteString(" {")
	buf.WriteString(`
    buf, err := xml.MarshalIndent(data, "", "    ")
    if err != nil {
        return err
    }`)
	buf.WriteString(`
    err = ioutil.WriteFile(path, buf, perm)
    if err != nil {
    	return err
    }`)
	buf.WriteString("\n")
	buf.WriteString("    return nil\n")
	buf.WriteString("}")
}

func WriteClassDef(buf *bytes.Buffer, class *xstruct.Class) {
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
