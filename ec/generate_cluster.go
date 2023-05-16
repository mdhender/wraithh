// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

import (
	"errors"
	"fmt"
	"github.com/mdhender/wraithh/generators/points"
	"html/template"
	"io/fs"
	"log"
	"math"
	"os"
)

type GenerateClusterOption func(c *clusterGeneratorConfig) error
type clusterGeneratorConfig struct {
	initSystems int                  // number of systems to seed cluster with
	pgen        func() *points.Point // points generator
	radius      float64
	sphereSize  float64
}

func SetKind(kind string) (func(config *clusterGeneratorConfig) error, error) {
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
	return func(config *clusterGeneratorConfig) error {
		config.pgen = pgen
		return nil
	}, nil
}
func SetRadius(r float64) (func(config *clusterGeneratorConfig) error, error) {
	if r < 8.0 || r > 45.0 {
		return nil, fmt.Errorf("radius must be between 8 and 45")
	}
	return func(config *clusterGeneratorConfig) error {
		config.radius = r
		config.sphereSize = config.radius / 45.0

		return nil
	}, nil
}
func SetSystems(n int) (func(config *clusterGeneratorConfig) error, error) {
	if n < 125 || n > 1024 {
		return nil, fmt.Errorf("init systems must be between 125 and 1024")
	}
	return func(config *clusterGeneratorConfig) error {
		config.initSystems = n
		return nil
	}, nil
}

func (e *Engine) GenerateCluster(options ...GenerateClusterOption) error {
	cfg := clusterGeneratorConfig{
		initSystems: 128,
		pgen:        points.ClusteredPoint,
		radius:      15.0,
		sphereSize:  15.0 / 45.0,
	}
	for _, opt := range options {
		if err := opt(&cfg); err != nil {
			return err
		}
	}

	// have we generated the cluster?
	var doGenerateCluster bool
	if _, err := os.Stat("cluster.json"); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		doGenerateCluster = true
	} else { // cluster exists
		doGenerateCluster = false
	}
	log.Printf("generate: cluster %v\n", doGenerateCluster)

	pp := points.NewPoints(cfg.initSystems*2, cfg.pgen)
	log.Println(pp.MinAvgMax())

	cp := pp.CullByCompanions(6)
	cpmin, cpavg, cpmax := cp.MinAvgMax()
	for cp.Length() > cfg.initSystems {
		cp = cp.CullByCompanions(6)
		cpmin, cpavg, cpmax = cp.MinAvgMax()
	}
	log.Printf("len %8d min %10.7f avg %10.7f max %10.7f\n", cp.Length(), cpmin, cpavg, cpmax)
	pp = cp

	type system struct {
		Id     int
		Coords Coordinates
		Size   float64
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		Color template.JS
		Warps []Point
	}

	var systems []*system
	for id, point := range pp.Points {
		scaled := point.Scale(cfg.radius)
		coords := Coordinates{
			X: int(math.Round(scaled.X)),
			Y: int(math.Round(scaled.Y)),
			Z: int(math.Round(scaled.Z)),
		}
		systems = append(systems, &system{
			Id:     id,
			Coords: coords,
			Size:   cfg.sphereSize,
			Color:  "Random",
		})
	}

	var tmplFile = "D:/wraith.dev/wraithh/templates/cluster.gohtml"
	ts, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}
	w, err := os.OpenFile("cluster.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	err = ts.Execute(w, systems)
	if err != nil {
		return err
	}
	_ = w.Close()
	return fmt.Errorf("!implemented")
}
