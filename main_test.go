package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindParent(t *testing.T) {
	// Create a more complex file tree.
	root := newFile("root")
	a := newFile("a")
	b := newFile("b")
	c := newFile("c")
	d := newFile("d")
	e := newFile("e")
	f := newFile("f")
	g := newFile("g")

	// Add children to form the tree.
	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)
	b.addChild(e)
	d.addChild(f)
	c.addChild(g)

	// Add aliases.
	x := newFile("x")
	root.addAlias(x, d) // Alias "x" for d
	y := newFile("y")
	c.addAlias(y, f) // Alias "y" for f
	z := newFile("z")
	e.addAlias(z, a) // Alias "z" for a

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
			desc:   "Alias x pointing to d",
			root:   root,
			file1:  x,
			file2:  f,
			parent: d,
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
