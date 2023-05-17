// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package systems

import (
	"github.com/mdhender/wraithh/models/coordinates"
	"github.com/mdhender/wraithh/models/orbits"
)

// Star is a single star system containing one or more Orbit(s)
type Star struct {
	Id       string // unique identifier for the star system
	Location coordinates.Coordinates
	Orbits   [11]orbits.Orbit
}
