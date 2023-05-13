// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
)

func (e *Engine) GenerateCluster() error {
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
	type system struct {
		Id     int
		Coords Point
		Delta  float64
		Size   float64
		// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
		Color template.JS
		Warps []Point
	}
	var points []Point
	const radius = 19
	const minDistance = radius / 4.0
	n := 0
	//pgen := newUniformPoint
	pgen := newSurfacePoint
	for len(points) < 125 {
		n++
		p := pgen().Scale(19.0)
		nearest := math.MaxFloat64
		for _, nbr := range points {
			if delta := p.DistanceTo(nbr); delta < nearest {
				nearest = delta
			}
		}
		if minDistance < nearest {
			points = append(points, p)
		}
	}
	log.Printf("generate: %d tries to create %d points\n", n, len(points))

	var systems []*system
	for id, p := range points {
		systems = append(systems, &system{
			Id:     id,
			Coords: p,
			Delta:  p.DistanceTo(Point{}),
			Size:   0.75,
			Color:  "Random",
		})
	}

	// calculate three closest neighbors
	for n := range systems {
		origin := systems[n].Coords
		type companion struct {
			point Point
			delta float64 // distance to system
		}
		var companions []companion
		for id, p := range points {
			if id == systems[n].Id {
				continue
			}
			companions = append(companions, companion{point: p, delta: origin.DistanceTo(p)})
		}
		sort.Slice(companions, func(i, j int) bool {
			return companions[i].delta < companions[j].delta
		})
		switch rand.Intn(2) + rand.Intn(2) + rand.Intn(2) {
		case 3:
			//systems[n].Warps = append(systems[n].Warps, companions[9].point)
			systems[n].Warps = append(systems[n].Warps, companions[3].point)
			fallthrough
		case 2:
			//systems[n].Warps = append(systems[n].Warps, companions[6].point)
			systems[n].Warps = append(systems[n].Warps, companions[2].point)
			fallthrough
		case 1:
			//systems[n].Warps = append(systems[n].Warps, companions[3].point)
			systems[n].Warps = append(systems[n].Warps, companions[1].point)
			fallthrough
		default:
			systems[n].Warps = append(systems[n].Warps, companions[0].point)
		}
	}

	sort.Slice(systems, func(i, j int) bool {
		return systems[i].Delta < systems[j].Delta
	})

	for id, p := range systems {
		p.Id = id
	}

	//for _, point := range systems {
	//	// Black, Blue, Gray, Green, Magenta, Purple, Random, Red, Teal, White, Yellow
	//	if point.Delta < 2 {
	//		point.Color = "Random"
	//	} else if point.Delta < 3 {
	//		point.Color = "Black"
	//	} else if point.Delta < 4 {
	//		point.Color = "Blue"
	//	} else if point.Delta < 6 {
	//		point.Color = "Gray"
	//	} else if point.Delta < 8 {
	//		point.Color = "Green"
	//	} else if point.Delta < 10 {
	//		point.Color = "Magenta"
	//	} else if point.Delta < 12 {
	//		point.Color = "Purple"
	//	} else if point.Delta < 14 {
	//		point.Color = "Red"
	//	} else if point.Delta < 16 {
	//		point.Color = "Teal"
	//	} else if point.Delta < 18 {
	//		point.Color = "White"
	//	} else {
	//		point.Color = "Yellow"
	//	}
	//}

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
