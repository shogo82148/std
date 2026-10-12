// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Spectasks renders SPEC-TRANSITION.md's task dependencies as a Graphviz
// graph, and checks the document's copies of those dependencies against each
// other.
//
//	go run ./cmd/spectasks -w        # write SPEC-TASKS.dot
//	go run ./cmd/spectasks           # check only; write the graph to stdout
//
// SPEC-TRANSITION.md states its dependencies three times: once in Part 2's
// table, which this tool treats as the source of truth; once per task in
// Part 3's "Needs:"/"Prefers:" header lines; and a third time, reversed, in
// Part 3's "Blocks:" lines. go test ./cmd/spectasks fails if they disagree, and
// also fails if SPEC-TASKS.dot has not been regenerated since. Adding the graph
// as a fourth hand-maintained copy would have made that worse, so the graph is
// generated.
//
// Two things the graph shows that the table cannot. The critical path is
// computed, not asserted: a task is on it when delaying it delays the whole
// project, which is the set of tasks whose longest-chain-in plus
// longest-chain-out spans the graph. And milestones -- the points worth
// announcing -- are drawn as nodes, so it is visible which tasks feed them.
// Milestone state is derived from the task checkboxes and is never written
// down separately.
package main
