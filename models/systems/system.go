// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package systems

import "github.com/mdhender/wraithh/models/coordinates"

type System struct {
	Id       int
	Location coordinates.Coordinates
}
