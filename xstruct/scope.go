package xstruct

import "github.com/beevik/etree"

// Scope is class scope
type Scope struct {
	Parent     *Scope
	Children   []*Scope
	Attributes map[string]string
	Name       string
	Content    string
}

// XMLToScopeTree is convert etree.Element to Scope
func XMLToScopeTree(elem *etree.Element) *Scope {
	scope := &Scope{
		Name:       elem.Tag,
		Content:    elem.Text(),
		Attributes: make(map[string]string),
	}
	for _, child := range elem.ChildElements() {
		childScope := XMLToScopeTree(child)
		scope.Children = append(scope.Children, childScope)
		childScope.Parent = scope
	}
	for _, attr := range elem.Attr {
		scope.Attributes[attr.Key] = attr.Value
	}
	return scope
}

// ToPath is hierarchy convert to string array
func (scope *Scope) ToPath() []string {
	var buf []string
	iter := scope
	for iter != nil {
		buf = append(buf, scope.Name)
		iter = iter.Parent
	}
	//see:https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return buf
}
