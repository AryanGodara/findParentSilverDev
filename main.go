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
}

func newFile(name string) *file {
	return &file{name: name}
}

func (f *file) addChild(child *file) {
	f.children = append(f.children, child)
}

func (root *file) findParent(file1, file2 *file) (*file, error) {
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

	// Find the closest common parent directory.
	parent, err := root.findParent(b, d)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parent.name)
}
