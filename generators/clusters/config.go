// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package clusters

import "github.com/mdhender/wraithh/generators/points"

type config struct {
	initSystems   int                  // number of systems to seed cluster with
	mapFile       string               // if set, create a map
	pgen          func() *points.Point // points generator
	radius        float64
	sphereSize    float64
	templatesPath string // path to template files
}
