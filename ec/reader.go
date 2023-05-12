// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

type GameJS struct {
	Id   string
	Name string
	Turn int
}

type PlayerJS struct {
	Handle string
	Nation string
}

func LoadGame(path string) (*Engine, error) {
	var e Engine
	e.Players = make(map[string]Player)

	var game GameJS
	if err := fromjson(path, "game", &game); err != nil {
		return nil, err
	}
	e.Game.Id = game.Id
	e.Game.Name = game.Name
	e.Game.Turn = game.Turn

	players := make(map[string]PlayerJS)
	if err := fromjson(path, "players", &players); err != nil {
		return nil, err
	}
	for k, player := range players {
		e.Players[k] = Player{
			Handle: player.Handle,
			Nation: player.Nation,
		}
	}

	return &e, nil
}