// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package ec implements the logic for Empyrean Challenge
package ec

import (
	"github.com/mdhender/wraithh/models/games"
	"github.com/mdhender/wraithh/models/player"
)

// Engine holds the state of a single game
type Engine struct {
	// Game holds information about the current game
	Game games.Game

	// Players holds every player that has ever been in this game.
	Players map[string]player.Player

	// Orders holds every player's set of orders for the current turn.
	Orders []*Orders
}
