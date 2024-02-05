package main

import (
	"fmt"
)

var (
	errFileNil      error = fmt.Errorf("file cannot be nil")
	errFileNotFound error = fmt.Errorf("file not found")
)

type file struct {
	name     string
	children []*file
	aliases  map[*file]*file
	isLink   bool
	target   *file
}

func newFile(name string) *file {
	return &file{
		name:     name,
		children: make([]*file, 0),
		aliases:  make(map[*file]*file),
		isLink:   false,
		target:   nil,
	}
}

func (f *file) addChild(child *file) {
	if f.isLink {
		fmt.Println("Cannot add a child to a soft link.")
		return
	}

	f.children = append(f.children, child)
}

func (f *file) addAlias(alias, target *file) {
	f.aliases[alias] = target
}

func (f *file) findFileByName(name string) (*file, bool) {
	if f == nil {
		return nil, true
	}

	if f.isLink {
		return f.target.findFileByName(name)
	}

	if f.name == name {
		return f, true
	}
	for alias, target := range f.aliases {
		if alias.name == name {
			return target, true
		}
	}
	for _, child := range f.children {
		foundFile, found := child.findFileByName(name)
		if found {
			return foundFile, true
		}
	}
	return nil, false
}

func (root *file) findParent(file1, file2 *file) (*file, error) {
	if root == nil || file1 == nil || file2 == nil {
		return nil, errFileNil
	}

	// Find the paths from the root to each file.
	pathToFile1, err := root.findPath(file1)
	if err != nil {
		return nil, err
	}
	pathToFile2, err := root.findPath(file2)
	if err != nil {
		return nil, err
	}

	// Find the closest common parent directory.
	var parent *file
	for i := 0; i < len(pathToFile1) && i < len(pathToFile2); i++ {
		if pathToFile1[i] == pathToFile2[i] {
			parent = pathToFile1[i]
		} else {
			break
		}
	}

	return parent, nil
}

// findPath is a helper function that finds the path from the root to the given file node.
func (root *file) findPath(_file *file) ([]*file, error) {
	if root == nil || _file == nil {
		return []*file{}, errFileNil
	}

	// If the root is the target file, return a path containing only the root.
	if root == _file {
		return []*file{root}, nil
	}

	// Recursively search for the target file in the children of the root.
	for _, child := range root.children {
		// Find the path from the child to the target file.
		path, err := child.findPath(_file)
		if err != nil && err != errFileNotFound {
			return nil, err
		}

		// If the target file is found in the subtree rooted at the child, append the child to the path and return it.
		if path != nil {
			return append([]*file{root}, path...), nil
		}
	}

	// If the target file is not found in the subtree rooted at the root, return nil.
	return nil, errFileNotFound
}

func main() {
	// Create a filesystem tree.
	root := newFile("root")
	a := newFile("a")
	b := newFile("b")
	c := newFile("c")
	d := newFile("d")

	// Add children to form the tree.
	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)

	// Add aliases
	_var := newFile("var")
	root.addAlias(_var, a)

	file1, found1 := root.findFileByName("b")
	if !found1 {
		fmt.Println("File b not found")
		return
	}

	file2, found2 := root.findFileByName("d")
	if !found2 {
		fmt.Println("Alias var not found")
		return
	}

	parent, err := root.findParent(file1, file2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The closest common parent directory is:", parent.name)
}
