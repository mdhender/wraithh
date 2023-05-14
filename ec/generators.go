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

func (e *Engine) GenerateCluster(kind string, radius float64, n int) error {
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

	pgen := points.ClusteredPoint
	if kind == "sphere" {
		pgen = points.SpherePoint
	} else if kind == "uniform" {
		pgen = points.UniformPoint
	}

	if n < 125 {
		n = 125
	} else if n > 1024 {
		n = 1024
	}

	pp := points.NewPoints(n*2, pgen)
	log.Println(pp.MinAvgMax())

	cp := pp.CullByCompanions(6)
	cpmin, cpavg, cpmax := cp.MinAvgMax()
	for cp.Length() > n {
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
	sphereSize := radius / 45.0
	for id, point := range pp.Points {
		scaled := point.Scale(radius)
		coords := Coordinates{
			X: int(math.Round(scaled.X)),
			Y: int(math.Round(scaled.Y)),
			Z: int(math.Round(scaled.Z)),
		}
		systems = append(systems, &system{
			Id:     id,
			Coords: coords,
			Size:   sphereSize,
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
