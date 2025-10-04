// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// This program converts CSV calibration data printed by
//
//	go test -run=Calibrate/Name -calibrate >file.csv
//
// into an SVG file. Invoke as:
//
//	go run calibrate_graph.go file.csv >file.svg
//
// See calibrate.md for more details.

package main

// A Point is an X, Y coordinate in the data being plotted.
type Point struct {
	X, Y float64
}

// A Graph is a graph to draw as SVG.
type Graph struct {
	Title   string
	Geomean []Point
	Lines   [][]Point
	XAxis   string
	YAxis   string
	Min     Point
	Max     Point
}

// Layout constants for drawing graph
const (
	DX   = 600
	DY   = 150
	ML   = 80
	MT   = 30
	MR   = 10
	MB   = 50
	PS   = 14
	W    = ML + DX + MR
	H    = MT + DY + MB
	Tick = 5
)

// An SVGPoint is a point in the SVG image, in pixel units,
// with Y increasing down the page.
type SVGPoint struct {
	X, Y int
}

func (p SVGPoint) String() string

// SVG returns the SVG text for the graph.
func (g *Graph) SVG() []byte
