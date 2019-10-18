package xstruct

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

// Namespace is container of class
type Namespace struct {
	Map map[string]*Class
}

// Class is golang struct
type Class struct {
	UniqueName     string
	UserName       string
	SimpleName     string
	SimpleUserName string
	Attributes     map[string]string
	InnerClasses   []*Class
}

// DefineClassA is define unique class by array
func DefineClassA(namespace *Namespace, path []string) (*Class, error) {
	un := strings.Join(path, "")
	var buf bytes.Buffer
	for _, component := range path {
		buf.WriteString(toWord(component))
	}
	return DefineClass(namespace, un, buf.String(), path[len(path)-1])
}

// DefineClass is define unique class
func DefineClass(namespace *Namespace, uniqueName string, userName string, simpleName string) (*Class, error) {
	if val, ok := namespace.Map[uniqueName]; ok {
		if val.SimpleName != simpleName {
			return nil, errors.New("required class already defined")
		}
		return val, nil
	}
	ret := &Class{
		UniqueName:     uniqueName,
		UserName:       userName,
		SimpleName:     simpleName,
		SimpleUserName: toWord(simpleName),
		Attributes:     make(map[string]string),
	}
	namespace.Map[uniqueName] = ret
	return ret, nil
}

// DefineClassTree is define class tree from scope
func DefineClassTree(namespace *Namespace, scope *Scope) error {
	class, err := DefineClassA(namespace, scope.ToPath())
	if err != nil {
		return err
	}
	for k, v := range scope.Attributes {
		class.Attributes[k] = v
	}
	for _, child := range scope.Children {
		e, err := DefineClassA(namespace, child.ToPath())
		if err != nil {
			return err
		}
		found := false
		for _, ic := range class.InnerClasses {
			if ic == e {
				found = true
				break
			}
		}
		if !found {
			class.InnerClasses = append(class.InnerClasses, e)
		}
	}
	for _, child := range scope.Children {
		DefineClassTree(namespace, child)
	}
	return nil
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
