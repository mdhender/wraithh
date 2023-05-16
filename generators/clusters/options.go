// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package clusters

import (
	"fmt"
	"github.com/mdhender/wraithh/generators/points"
	"path/filepath"
)

type Option func(c *config) error

func CreateHtmlMap(name string) (func(*config) error, error) {
	return func(config *config) error {
		config.mapFile = name
		if config.mapFile != "" {
			config.mapFile = filepath.Clean(config.mapFile)
		}
		return nil
	}, nil
}

func SetKind(kind string) (func(*config) error, error) {
	var pgen func() *points.Point
	switch kind {
	case "cluster":
		pgen = points.ClusteredPoint
	case "sphere": // okay
		pgen = points.SpherePoint
	case "uniform": // okay
		pgen = points.UniformPoint
	default:
		return nil, fmt.Errorf("kind must be uniform, cluster, or sphere")
	}
	return func(config *config) error {
		config.pgen = pgen
		return nil
	}, nil
}

func SetRadius(r float64) (func(*config) error, error) {
	if r < minRadius || r > maxRadius {
		return nil, fmt.Errorf("radius must be between %3.1f and %3.1f", minRadius, maxRadius)
	}
	return func(config *config) error {
		config.radius = r
		config.sphereSize = config.radius * sphereRatio

		return nil
	}, nil
}

func SetSystems(n int) (func(*config) error, error) {
	if n < minSystemSeeds || n > maxSystemSeeds {
		return nil, fmt.Errorf("init systems must be between %d and %d", minSystemSeeds, maxSystemSeeds)
	}
	return func(config *config) error {
		config.initSystems = n
		return nil
	}, nil
}
