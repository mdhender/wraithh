// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import "github.com/mdhender/wraithh/internal/tree"

// Tree is copied from https://github.com/shivamMg/rd
// which is released under the MIT license and is
// copyright (c) 2018 Shivam Mamgain.

// Tree is a parse tree node.
// Symbol can either be a terminal (Token) or a non-terminal (see Builder's Enter method).
// Tokens matched using Builder's Match method or added using Builder's Add method, can be retrieved by type asserting Symbol.
// Subtrees are child nodes of the current node.
type Tree struct {
	Symbol   interface{}
	Subtrees []*Tree
}

func NewTree(symbol interface{}, subtrees ...*Tree) *Tree {
	t := Tree{Symbol: symbol}
	for _, subtree := range subtrees {
		if subtree != nil {
			t.Subtrees = append(t.Subtrees, subtree)
		}
	}
	return &t
}

func (t *Tree) Data() interface{} {
	if t == nil {
		return ""
	}
	return t.Symbol
}

func (t *Tree) Children() (c []tree.Node) {
	for _, subtree := range t.Subtrees {
		c = append(c, subtree)
	}
	return
}

// Add adds a subtree as a child to t.
func (t *Tree) Add(subtree *Tree) {
	t.Subtrees = append(t.Subtrees, subtree)
}

// Detach removes a subtree as a child of t.
func (t *Tree) Detach(subtree *Tree) {
	for i, st := range t.Subtrees {
		if st == subtree {
			t.Subtrees = append(t.Subtrees[:i], t.Subtrees[i+1:]...)
			break
		}
	}
}

func (t *Tree) String() string {
	return tree.SprintHrn(t)
}
