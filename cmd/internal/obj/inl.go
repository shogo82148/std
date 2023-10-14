// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import "github.com/shogo82148/std/cmd/internal/src"

// InlTree is a collection of inlined calls. The Parent field of an
// InlinedCall is the index of another InlinedCall in InlTree.
//
// The compiler maintains a global inlining tree and adds a node to it
// every time a function is inlined. For example, suppose f() calls g()
// and g has two calls to h(), and that f, g, and h are inlineable:
//
//	 1 func main() {
//	 2     f()
//	 3 }
//	 4 func f() {
//	 5     g()
//	 6 }
//	 7 func g() {
//	 8     h()
//	 9     h()
//	10 }
//	11 func h() {
//	12     println("H")
//	13 }
//
// Assuming the global tree starts empty, inlining will produce the
// following tree:
//
//	[]InlinedCall{
//	  {Parent: -1, Func: "f", Pos: <line 2>},
//	  {Parent:  0, Func: "g", Pos: <line 5>},
//	  {Parent:  1, Func: "h", Pos: <line 8>},
//	  {Parent:  1, Func: "h", Pos: <line 9>},
//	}
//
// The nodes of h inlined into main will have inlining indexes 2 and 3.
//
// Eventually, the compiler extracts a per-function inlining tree from
// the global inlining tree (see pcln.go).
type InlTree struct {
	nodes []InlinedCall
}

// InlinedCall is a node in an InlTree.
type InlinedCall struct {
	Parent   int
	Pos      src.XPos
	Func     *LSym
	Name     string
	ParentPC int32
}

// Add adds a new call to the tree, returning its index.
func (tree *InlTree) Add(parent int, pos src.XPos, func_ *LSym, name string) int

// AllParents invokes do on each InlinedCall in the inlining call
// stack, from outermost to innermost.
//
// That is, if inlIndex corresponds to f inlining g inlining h,
// AllParents invokes do with the call for inlining g into f, and then
// inlining h into g.
func (tree *InlTree) AllParents(inlIndex int, do func(InlinedCall))

func (tree *InlTree) Parent(inlIndex int) int

func (tree *InlTree) InlinedFunction(inlIndex int) *LSym

func (tree *InlTree) CallPos(inlIndex int) src.XPos

// OutermostPos returns the outermost position corresponding to xpos,
// which is where xpos was ultimately inlined to. In the example for
// InlTree, main() contains inlined AST nodes from h(), but the
// outermost position for those nodes is line 2.
func (ctxt *Link) OutermostPos(xpos src.XPos) src.Pos

// InnermostPos returns the innermost position corresponding to xpos,
// that is, the code that is inlined and that inlines nothing else.
// In the example for InlTree above, the code for println within h
// would have an innermost position with line number 12, whether
// h was not inlined, inlined into g, g-then-f, or g-then-f-then-main.
// This corresponds to what someone debugging main, f, g, or h might
// expect to see while single-stepping.
func (ctxt *Link) InnermostPos(xpos src.XPos) src.Pos

// AllPos invokes do with every position in the inlining call stack for xpos,
// from outermost to innermost. That is, xpos corresponds to f inlining g inlining h,
// AllPos invokes do with the position in f, then the position in g, then the position in h.
func (ctxt *Link) AllPos(xpos src.XPos, do func(src.Pos))
