package main

import (
	"fmt"
)

var (
	errFileNil error = fmt.Errorf("file cannot be nil")
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

func (f *file) getChildren(_file *file) []*file {
	return _file.children
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

	// If either path is empty, return nil.
	if len(pathToFile1) == 0 || len(pathToFile2) == 0 {
		return nil, fmt.Errorf("path not found")
	}

	return nil, nil
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
		path := child.findPath(_file)

		// If the path contains the target file, add the root to the path and return it.
		if len(path) > 0 {
			return append([]*file{root}, path...)
		}
	}

	// If the target file is not found in the subtree rooted at the root, return nil.
	return nil, nil
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
	parent := root.findParent(c, d)
	if parent != nil {
		fmt.Println("The closest common parent directory is:", parent.name)
	} else {
		fmt.Println("No common parent directory found.")
	}
}
