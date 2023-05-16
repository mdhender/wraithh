// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package clusters

import (
	"github.com/google/uuid"
	"github.com/mdhender/wraithh/generators/points"
	"github.com/mdhender/wraithh/models/cluster"
	"github.com/mdhender/wraithh/models/coordinates"
	"github.com/mdhender/wraithh/models/systems"
	"html/template"
	"log"
	"math"
	"os"
)

const (
	minSystemSeeds, maxSystemSeeds = 125, 1024
	minRadius, maxRadius           = 5.0, 45.0
	defaultRadius                  = 15.0
	sphereRatio                    = defaultRadius / maxRadius
)

// Generate creates a new cluster.
func Generate(options ...Option) (*cluster.Cluster, error) {
	cfg := config{
		initSystems:   128,
		pgen:          points.ClusteredPoint,
		radius:        15.0,
		sphereSize:    sphereRatio,
		templatesPath: "D:/wraith.dev/wraithh/templates/cluster.gohtml",
	}
	for _, opt := range options {
		if err := opt(&cfg); err != nil {
			return nil, err
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
	pp.SortByDistanceOrigin()

	type system struct {
		Id     int
		Coords coordinates.Coordinates
		NStars int
		Size   float64
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		Color template.JS
		Warps []coordinates.Point
	}

	// distribution of multi-star systems
	nstarslist := []int{4, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	//rand.Shuffle(len(nstarslist), func(i, j int) {
	//	nstarslist[i], nstarslist[j] = nstarslist[j], nstarslist[i]
	//})

	var set []*system
	for id, point := range pp.Points {
		var nstars int
		if len(nstarslist) == 0 {
			nstars = 1
		} else {
			nstars, nstarslist = nstarslist[0], nstarslist[1:]
		}
		scaled := point.Scale(cfg.radius)
		coords := coordinates.Coordinates{
			X: int(math.Round(scaled.X)),
			Y: int(math.Round(scaled.Y)),
			Z: int(math.Round(scaled.Z)),
		}
		ss := &system{
			Id:     id,
			Coords: coords,
			Size:   cfg.sphereSize,
			NStars: nstars,
		}
		switch nstars {
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		case 1:
			ss.Color = "Gray"
		case 2:
			ss.Color = "Blue"
		case 3:
			ss.Color = "Red"
		case 4:
			ss.Color = "Teal"
		default:
			ss.Color = "Random"
		}
		set = append(set, ss)
	}

	c := &cluster.Cluster{Radius: cfg.radius}
	for _, sys := range set {
		s := systems.System{
			Id:       uuid.New().String(),
			Location: sys.Coords,
		}
		c.Systems = append(c.Systems, s)
	}

	if cfg.mapFile != "" {
		ts, err := template.ParseFiles(cfg.templatesPath)
		if err != nil {
			return nil, err
		}
		w, err := os.OpenFile(cfg.mapFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		defer func(fp *os.File) {
			_ = fp.Close()
		}(w)
		err = ts.Execute(w, set)
		if err != nil {
			return nil, err
		}
		log.Printf("cluster: created %q\n", cfg.mapFile)
	}

	return c, nil
}
