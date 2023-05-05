// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

// Element is copied from https://github.com/shivamMg/rd
// which is released under the MIT license and is
// copyright (c) 2018 Shivam Mamgain.

type element struct {
	index   int
	nonTerm *Tree
}
