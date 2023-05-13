// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

import (
	"encoding/json"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

type Point struct { // location being set up
	X, Y, Z float64
}

func (p Point) DistanceTo(b Point) float64 {
	dx, dy, dz := p.X-b.X, p.Y-b.Y, p.Z-b.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p Point) Scale(n float64) Point {
	return Point{X: p.X * n, Y: p.Y * n, Z: p.Z * n}
}

func fromjson(path, name string, data any) error {
	buffer, err := os.ReadFile(filepath.Clean(filepath.Join(path, name+".json")))
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, data)
}

func newClusteredPoint() Point {
	var u = rand.Float64()
	var v = rand.Float64()
	var theta = u * 2.0 * math.Pi
	var phi = math.Acos(2.0*v - 1.0)
	var sinTheta = math.Sin(theta)
	var cosTheta = math.Cos(theta)
	var sinPhi = math.Sin(phi)
	var cosPhi = math.Cos(phi)
	var r = rand.Float64()
	return Point{
		X: r * sinPhi * cosTheta,
		Y: r * sinPhi * sinTheta,
		Z: r * cosPhi,
	}
}

func newSurfacePoint() Point {
	var u = rand.Float64()
	var v = rand.Float64()
	var theta = u * 2.0 * math.Pi
	var phi = math.Acos(2.0*v - 1.0)
	var sinTheta = math.Sin(theta)
	var cosTheta = math.Cos(theta)
	var sinPhi = math.Sin(phi)
	var cosPhi = math.Cos(phi)
	return Point{
		X: sinPhi * cosTheta,
		Y: sinPhi * sinTheta,
		Z: cosPhi,
	}
}

func newUniformPoint() Point {
	var u = rand.Float64()
	var v = rand.Float64()
	var theta = u * 2.0 * math.Pi
	var phi = math.Acos(2.0*v - 1.0)
	var sinTheta = math.Sin(theta)
	var cosTheta = math.Cos(theta)
	var sinPhi = math.Sin(phi)
	var cosPhi = math.Cos(phi)
	var r = math.Cbrt(rand.Float64())
	return Point{
		X: r * sinPhi * cosTheta,
		Y: r * sinPhi * sinTheta,
		Z: r * cosPhi,
	}
}

func tojson(path, name string, data any) error {
	buffer, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Clean(filepath.Join(path, name+".json")), buffer, 0644)
}
