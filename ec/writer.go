// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

import (
	"path/filepath"
)

func (e *Engine) SaveGame(path string) error {
	path = filepath.Join(path, "out")
	game := GameJS{
		Id:   e.Game.Id,
		Name: e.Game.Name,
		Turn: e.Game.Turn,
	}
	if err := tojson(path, "game", game); err != nil {
		return err
	}

	players := make(map[string]PlayerJS)
	for k, player := range e.Players {
		players[k] = PlayerJS{
			Handle: player.Handle,
			Nation: player.Nation,
		}
	}
	if err := tojson(path, "players", players); err != nil {
		return err
	}

	return nil
}
