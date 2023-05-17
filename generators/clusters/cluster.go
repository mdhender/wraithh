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
		clustered:     true,
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
		Id     string
		Coords coordinates.Coordinates
		NStars int
		Size   float64
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		Color template.JS
		Warps []coordinates.Point
	}

	// distribution of multi-star systems
	nstarslist := []int{4, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	//rand.Shuffle(len(nstarslist), func(i, j int) {
	//	nstarslist[i], nstarslist[j] = nstarslist[j], nstarslist[i]
	//})
	log.Printf("cluster: nstars %d\n", len(nstarslist))

	var set []*system
	locations := make(map[string]*system)
	for n, point := range pp.Points {
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
		id := coords.String()
		if locations[id] != nil {
			locations[id].NStars += nstars
			log.Printf("cluster: %d collided!\n", n)
		} else {
			set = append(set, &system{
				Id:     id,
				Coords: coords,
				Size:   cfg.sphereSize,
				NStars: nstars,
			})
		}
	}
	for _, ss := range set {
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		switch ss.NStars {
		case 1:
			ss.Color = "Black"
		case 2:
			ss.Color = "Blue"
		case 3:
			ss.Color = "Gray"
		case 4:
			ss.Color = "Green"
		case 5:
			ss.Color = "Magenta"
		case 6:
			ss.Color = "Purple"
		case 7:
			ss.Color = "Red"
		case 8:
			ss.Color = "Teal"
		case 9:
			ss.Color = "White"
		case 10:
			ss.Color = "Yellow"
		default:
			ss.Color = "Random"
		}
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
