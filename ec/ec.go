// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package ec implements the logic for Empyrean Challenge
package ec

import "github.com/mdhender/wraithh/ec/types"

type Engine struct {
	Game struct {
		Id   string
		Name string
		Turn int
	}
	Players map[string]types.Player
	Orders  []*Orders
}

type Orders struct {
	Validated bool
	Handle    string
	Game      string
	Turn      int
	Secret    *Secret
	Orders    []types.Order
	Error     error
}
