// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ships

import "github.com/mdhender/wraithh/models/coordinates"

// Ship is either a ship or a colony(?!!?).
type Ship struct {
	Id       string // unique identifier for ship or colony
	Kind     Kind
	Location coordinates.Coordinates
	// attributes like hull, cargo, bridge, engines
}
