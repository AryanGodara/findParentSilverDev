package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindParent(t *testing.T) {
	// Create a simple file tree.
	root := &file{name: "root"}
	a := &file{name: "a"}
	b := &file{name: "b"}
	c := &file{name: "c"}
	d := &file{name: "d"}

	// Add children to form the tree.
	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)

	// Test cases.
	cases := []struct {
		desc   string
		root   *file
		file1  *file
		file2  *file
		parent *file
		err    error
	}{}

	for _, tc := range cases {
		par, err := tc.root.findParent(tc.file1, tc.file2)
		assert.Equal(t, err, tc.err, fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))

		if par != nil {
			assert.Equal(t, par, tc.parent, fmt.Sprintf("%s: expected %v got %v\n", tc.desc, tc.parent, par))
		}
	}
}
