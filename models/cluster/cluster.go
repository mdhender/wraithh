// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cluster

type Cluster struct {
	Radius  float64
	Systems []string // id for every system in the cluster
	Stars   []string // id for every star in the cluster
}
