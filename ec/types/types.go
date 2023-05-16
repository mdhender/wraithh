// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package types

import "fmt"

type Player struct {
	Id     string
	Handle string
	Secret string
	Nation string
}

type Cluster struct {
	Systems []System
}

type Coordinates struct { // location being set up
	X, Y, Z int
	System  string // suffix for multi-star system, A...Z
	Orbit   int
}

type Point struct {
	X, Y, Z float64
}

type System struct {
	Location Coordinates
}

func (c Coordinates) String() string {
	if c.Orbit == 0 {
		return fmt.Sprintf("(%d,%d,%d%s)", c.X, c.Y, c.Z, c.System)
	}
	return fmt.Sprintf("(%d,%d,%d%s, %d)", c.X, c.Y, c.Z, c.System, c.Orbit)
}

type Order interface {
	Execute() error
}

type TransferDetail struct {
	Unit     Unit
	Quantity int
}

func (td *TransferDetail) String() string {
	return fmt.Sprintf("{%d %s}", td.Quantity, td.Unit)
}

type Unit struct {
	Name      string // name
	TechLevel int    // optional tech level
}

func (u Unit) String() string {
	if u.TechLevel == 0 {
		return u.Name
	}
	return fmt.Sprintf("%s-%d", u.Name, u.TechLevel)
}
