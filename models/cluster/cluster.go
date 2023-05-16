// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cluster

import "github.com/mdhender/wraithh/models/systems"

type Cluster struct {
	Radius  float64
	Systems []*systems.System
}
