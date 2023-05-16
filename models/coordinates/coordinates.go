// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package coordinates

import "fmt"

type Coordinates struct { // location being set up
	X, Y, Z int
	System  string // suffix for multi-star system, A...Z
	Orbit   int
}

func (c Coordinates) String() string {
	if c.Orbit == 0 {
		return fmt.Sprintf("(%d,%d,%d%s)", c.X, c.Y, c.Z, c.System)
	}
	return fmt.Sprintf("(%d,%d,%d%s, %d)", c.X, c.Y, c.Z, c.System, c.Orbit)
}
