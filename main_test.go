package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindParent(t *testing.T) {
	root := newFile("root")
	a := newFile("a")
	b := newFile("b")
	c := newFile("c")
	d := newFile("d")
	e := newFile("e")
	f := newFile("f")
	g := newFile("g")

	// Add aliases.
	x := newFile("x")
	root.addAlias(x, d) // Alias "x" for d
	y := newFile("y")
	c.addAlias(y, f) // Alias "y" for f
	z := newFile("z")
	e.addAlias(z, a) // Alias "z" for a

	// Adding more files for soft link demonstration
	h := newFile("h") // New directory
	i := newFile("i") // New file in 'h'

	// Add children to form the tree.
	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)
	b.addChild(e)
	d.addChild(f)
	c.addChild(g)
	h.addChild(i) // Add 'i' to 'h'

	// Incorporating soft links (previously aliases).
	root.addAlias(h, d) // Soft link "h" to "d", simulating a directory 'h' pointing to 'd'
	c.addAlias(i, f)    // Soft link "i" to "f", simulating a file 'i' pointing to 'f'

	// Test cases including nil files, and aliases.
	cases := []struct {
		desc   string
		root   *file
		file1  *file
		file2  *file
		parent *file
		err    error
	}{
		{
			desc:   "Common parent of c and d",
			root:   root,
			file1:  c,
			file2:  d,
			parent: a,
			err:    nil,
		},
		{
			desc:   "Common parent of e and g",
			root:   root,
			file1:  e,
			file2:  g,
			parent: root,
			err:    nil,
		},
		{
			desc:   "Nil file1",
			root:   root,
			file1:  nil,
			file2:  b,
			parent: nil,
			err:    errFileNil,
		},
		{
			desc:   "Nil file2",
			root:   root,
			file1:  a,
			file2:  nil,
			parent: nil,
			err:    errFileNil,
		},
		{
			desc:   "Alias x pointing to d",
			root:   root,
			file1:  x,
			file2:  f,
			parent: d,
			err:    nil,
		},
		{
			desc:   "Alias y and g",
			root:   root,
			file1:  y,
			file2:  g,
			parent: a,
			err:    nil,
		},
		{
			desc:   "Alias z (for a) and b",
			root:   root,
			file1:  z,
			file2:  b,
			parent: root,
			err:    nil,
		},
		{
			desc:   "Common parent of c and d, considering soft links",
			root:   root,
			file1:  c,
			file2:  d,
			parent: a,
			err:    nil,
		},
		{
			desc:   "Soft link 'h' pointing to 'd' and its child 'f'",
			root:   root,
			file1:  h, // Now represents a soft link to 'd'
			file2:  a,
			parent: a, // Expected common parent is 'a', since 'h' is a soft link to 'd'
			err:    nil,
		},
		{
			desc:   "Soft link 'i' within 'h' pointing to 'f'",
			root:   root,
			file1:  i, // Represents a soft link/file in 'h' pointing to 'f'
			file2:  g,
			parent: a, // 'g' is under 'c', and 'i' is a soft link pointing to 'f' under 'd', both under 'a'
			err:    nil,
		},
	}

	for _, tc := range cases {
		if tc.file1 != nil {
			tc.file1, _ = tc.root.findFileByName(tc.file1.name)
		}
		if tc.file2 != nil {
			tc.file2, _ = tc.root.findFileByName(tc.file2.name)
		}
		par, err := tc.root.findParent(tc.file1, tc.file2)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected %v got %v\n", tc.desc, tc.err, err))

		if err == nil {
			assert.Equal(t, tc.parent, par, fmt.Sprintf("%s: expected %v got %v\n", tc.desc, tc.parent.name, par.name))
		}
	}
}
