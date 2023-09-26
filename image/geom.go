// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point struct {
	X, Y int
}

// String returns a string representation of p like "(3,4)".
func (p Point) String() string

// Add returns the vector p+q.
func (p Point) Add(q Point) Point

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point

// Mul returns the vector p*k.
func (p Point) Mul(k int) Point

// Div returns the vector p/k.
func (p Point) Div(k int) Point

// In returns whether p is in r.
func (p Point) In(r Rectangle) bool

// Mod returns the point q in r such that p.X-q.X is a multiple of r's width
// and p.Y-q.Y is a multiple of r's height.
func (p Point) Mod(r Rectangle) Point

// Eq returns whether p and q are equal.
func (p Point) Eq(q Point) bool

// ZP is the zero Point.
var ZP Point

// Pt is shorthand for Point{X, Y}.
func Pt(X, Y int) Point

// A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
// It is well-formed if Min.X <= Max.X and likewise for Y. Points are always
// well-formed. A rectangle's methods always return well-formed outputs for
// well-formed inputs.
type Rectangle struct {
	Min, Max Point
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rectangle) String() string

// Dx returns r's width.
func (r Rectangle) Dx() int

// Dy returns r's height.
func (r Rectangle) Dy() int

// Size returns r's width and height.
func (r Rectangle) Size() Point

// Add returns the rectangle r translated by p.
func (r Rectangle) Add(p Point) Rectangle

// Sub returns the rectangle r translated by -p.
func (r Rectangle) Sub(p Point) Rectangle

// Inset returns the rectangle r inset by n, which may be negative. If either
// of r's dimensions is less than 2*n then an empty rectangle near the center
// of r will be returned.
func (r Rectangle) Inset(n int) Rectangle

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rectangle) Intersect(s Rectangle) Rectangle

// Union returns the smallest rectangle that contains both r and s.
func (r Rectangle) Union(s Rectangle) Rectangle

// Empty returns whether the rectangle contains no points.
func (r Rectangle) Empty() bool

// Eq returns whether r and s are equal.
func (r Rectangle) Eq(s Rectangle) bool

// Overlaps returns whether r and s have a non-empty intersection.
func (r Rectangle) Overlaps(s Rectangle) bool

// In returns whether every point in r is in s.
func (r Rectangle) In(s Rectangle) bool

// Canon returns the canonical version of r. The returned rectangle has minimum
// and maximum coordinates swapped if necessary so that it is well-formed.
func (r Rectangle) Canon() Rectangle

// ZR is the zero Rectangle.
var ZR Rectangle

// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}.
func Rect(x0, y0, x1, y1 int) Rectangle
