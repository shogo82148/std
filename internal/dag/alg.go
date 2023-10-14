// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dag

// Transpose reverses all edges in g.
func (g *Graph) Transpose()

// Topo returns a topological sort of g. This function is deterministic.
func (g *Graph) Topo() []string

// TransitiveReduction removes edges from g that are transitively
// reachable. g must be transitively closed.
func (g *Graph) TransitiveReduction()
