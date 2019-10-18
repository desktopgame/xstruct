package detail

import (
	"bytes"
	"log"
	"unicode"

	"github.com/beevik/etree"
	"github.com/desktopgame/xstruct/xstruct"
)

type Option struct {
	Package string
	Prefix  string
	Suffix  string
}

func CreateProgram(path string, opt Option) bytes.Buffer {
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
	buf.WriteString("package ")
	buf.WriteString(opt.Package)
	buf.WriteString("\n")
	for _, class := range namespace.Map {
		WriteClassDef(opt, &buf, class)
	}
	root, err := xstruct.DefineClassA(namespace, sc.ToPath())
	if err != nil {
		log.Fatal(err)
	}
	WriteFuncDef(opt, &buf, root)
	return buf
}

func WriteFuncDef(opt Option, buf *bytes.Buffer, class *xstruct.Class) {
	// LoadFunc
	buf.WriteString("func Load")
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString("(path string) (*")
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString(", error) {")
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
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString(`
    xml.Unmarshal(xmlData, &data)
    return &data, nil`)
	buf.WriteString("\n}")
	buf.WriteString("\n")
	buf.WriteString("\n")
	// SaveFunc
	buf.WriteString("func Save")
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString("(path string, data *")
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString(", perm os.FileMode) error")
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

func WriteClassDef(opt Option, buf *bytes.Buffer, class *xstruct.Class) {
	buf.WriteString("type ")
	buf.WriteString(fixWord(opt, class.UserName))
	buf.WriteString(" struct {\n")
	buf.WriteString("    // define attribute\n")
	for k, _ := range class.Attributes {
		buf.WriteString("    ")
		buf.WriteString("Attr")
		buf.WriteString(k)
		buf.WriteString(" string `xml:\"")
		buf.WriteString(k)
		buf.WriteString(",attr\"`\n")
	}
	buf.WriteString("    // define subelement\n")
	for _, class := range class.InnerClasses {
		buf.WriteString("    ")
		buf.WriteString("Sub")
		buf.WriteString(toWord(class.SimpleUserName))
		buf.WriteString(" []*")
		buf.WriteString(fixWord(opt, class.UserName))
		buf.WriteString(" `xml:\"")
		buf.WriteString(class.SimpleName)
		buf.WriteString("\"`\n")
	}
	buf.WriteString("    // define content\n")
	buf.WriteString("    Content string `xml:\",chardata\"`\n")
	buf.WriteString("}\n")
}

func fixWord(opt Option, str string) string {
	return opt.Prefix + str + opt.Suffix
}

func toWord(str string) string {
	var buf bytes.Buffer
	for idx, rn := range str {
		if idx == 0 {
			buf.WriteRune(unicode.ToUpper(rn))
		} else {
			buf.WriteRune(rn)
		}
	}
	return buf.String()
}
