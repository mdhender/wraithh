// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package systems

import "github.com/mdhender/wraithh/models/coordinates"

// System is a stellar system containing one or more stars.
type System struct {
	Id       string                  // unique identifier for the system
	Location coordinates.Coordinates // location of the system
	Stars    []*Star                 // stars in the system
}
