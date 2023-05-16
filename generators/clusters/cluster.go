// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package clusters

import (
	"fmt"
	"github.com/mdhender/wraithh/ec/types"
	"github.com/mdhender/wraithh/generators/points"
	"html/template"
	"log"
	"math"
	"os"
)

func Generate(options ...Option) error {
	cfg := config{
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
		Coords types.Coordinates
		Size   float64
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		Color template.JS
		Warps []types.Point
	}

	var systems []*system
	for id, point := range pp.Points {
		scaled := point.Scale(cfg.radius)
		coords := types.Coordinates{
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
