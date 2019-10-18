package xstruct

import (
	"errors"
	"strings"
)

// Namespace is container of class
type Namespace struct {
	Map map[string]*Class
}

// Class is golang struct
type Class struct {
	UniqueName   string
	SimpleName   string
	Attributes   map[string]string
	InnerClasses []*Class
}

// DefineClassA is define unique class by array
func DefineClassA(namespace *Namespace, path []string) (*Class, error) {
	un := strings.Join(path[:], "")
	return DefineClass(namespace, un, path[len(path)-1])
}

// DefineClass is define unique class
func DefineClass(namespace *Namespace, uniqueName string, simpleName string) (*Class, error) {
	if val, ok := namespace.Map[uniqueName]; ok {
		if val.SimpleName != simpleName {
			return nil, errors.New("required class already defined")
		}
		return val, nil
	}
	ret := &Class{
		UniqueName: uniqueName,
		SimpleName: simpleName,
		Attributes: make(map[string]string),
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
		class.InnerClasses = append(class.InnerClasses, e)
	}
	for _, child := range scope.Children {
		DefineClassTree(namespace, child)
	}
	return nil
}
