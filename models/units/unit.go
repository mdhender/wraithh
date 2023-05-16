// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package units

import "fmt"

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
