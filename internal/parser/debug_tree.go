// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import "github.com/mdhender/wraithh/internal/tree"

// DebugTree is copied from https://github.com/shivamMg/rd
// which is released under the MIT license and is
// copyright (c) 2018 Shivam Mamgain.

// DebugTree is a debug tree node.
// Can be printed to help tracing the parsing flow.
type DebugTree struct {
	data     string
	subtrees []*DebugTree
}

func newDebugTree(data string) *DebugTree {
	return &DebugTree{
		data:     data,
		subtrees: []*DebugTree{},
	}
}

func (dt *DebugTree) add(subtree *DebugTree) {
	dt.subtrees = append(dt.subtrees, subtree)
}

func (dt *DebugTree) Data() interface{} {
	return dt.data
}

func (dt *DebugTree) Children() (c []tree.Node) {
	for _, child := range dt.subtrees {
		c = append(c, child)
	}
	return
}

func (dt *DebugTree) String() string {
	return tree.SprintHrn(dt)
}
