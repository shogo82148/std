// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Only run where builders (build.golang.org) have
// access to compiled packages for import.
//
//go:build !android && !ios && !js && !wasip1

package types2_test

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/regexp"
	"github.com/shogo82148/std/sort"
	"github.com/shogo82148/std/strings"
)

// ExampleScope prints the tree of Scopes of a package created from a
// set of parsed files.
func ExampleScope() {
	// Parse the source files for a package.
	var files []*syntax.File
	for _, src := range []string{
		`package main
import "fmt"
func main() {
	freezing := FToC(-18)
	fmt.Println(freezing, Boiling) }
`,
		`package main
import "fmt"
type Celsius float64
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func FToC(f float64) Celsius { return Celsius(f - 32 / 9 * 5) }
const Boiling Celsius = 100
func Unused() { {}; {{ var x int; _ = x }} } // make sure empty block scopes get printed
`,
	} {
		files = append(files, mustParse(src))
	}

	// Type-check a package consisting of these files.
	// Type information for the imported "fmt" package
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types2.Config{Importer: defaultImporter()}
	pkg, err := conf.Check("temperature", files, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tree of scopes.
	// For determinism, we redact addresses.
	var buf strings.Builder
	pkg.Scope().WriteTo(&buf, 0, true)
	rx := regexp.MustCompile(` 0x[a-fA-F0-9]*`)
	fmt.Println(rx.ReplaceAllString(buf.String(), ""))

	// Output:
	// package "temperature" scope {
	// .  const temperature.Boiling temperature.Celsius
	// .  type temperature.Celsius float64
	// .  func temperature.FToC(f float64) temperature.Celsius
	// .  func temperature.Unused()
	// .  func temperature.main()
	// .  main scope {
	// .  .  package fmt
	// .  .  function scope {
	// .  .  .  var freezing temperature.Celsius
	// .  .  }
	// .  }
	// .  main scope {
	// .  .  package fmt
	// .  .  function scope {
	// .  .  .  var c temperature.Celsius
	// .  .  }
	// .  .  function scope {
	// .  .  .  var f float64
	// .  .  }
	// .  .  function scope {
	// .  .  .  block scope {
	// .  .  .  }
	// .  .  .  block scope {
	// .  .  .  .  block scope {
	// .  .  .  .  .  var x int
	// .  .  .  .  }
	// .  .  .  }
	// .  .  }
	// .  }
	// }
}

// ExampleInfo prints various facts recorded by the type checker in a
// types2.Info struct: definitions of and references to each named object,
// and the type, value, and mode of every expression in the package.
func ExampleInfo() {
	// Parse a single source file.
	const input = `
package fib

type S string

var a, b, c = len(b), S(c), "hello"

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) - fib(x-2)
}`
	// Type-check the package.
	// We create an empty map for each kind of input
	// we're interested in, and Check populates them.
	info := types2.Info{
		Types: make(map[syntax.Expr]types2.TypeAndValue),
		Defs:  make(map[*syntax.Name]types2.Object),
		Uses:  make(map[*syntax.Name]types2.Object),
	}
	pkg := mustTypecheck(input, nil, &info)

	// Print package-level variables in initialization order.
	fmt.Printf("InitOrder: %v\n\n", info.InitOrder)

	// For each named object, print the line and
	// column of its definition and each of its uses.
	fmt.Println("Defs and Uses of each named object:")
	usesByObj := make(map[types2.Object][]string)
	for id, obj := range info.Uses {
		posn := id.Pos()
		lineCol := fmt.Sprintf("%d:%d", posn.Line(), posn.Col())
		usesByObj[obj] = append(usesByObj[obj], lineCol)
	}
	var items []string
	for obj, uses := range usesByObj {
		sort.Strings(uses)
		item := fmt.Sprintf("%s:\n  defined at %s\n  used at %s",
			types2.ObjectString(obj, types2.RelativeTo(pkg)),
			obj.Pos(),
			strings.Join(uses, ", "))
		items = append(items, item)
	}
	sort.Strings(items) // sort by line:col, in effect
	fmt.Println(strings.Join(items, "\n"))
	fmt.Println()

	// TODO(gri) Enable once positions are updated/verified
	// fmt.Println("Types and Values of each expression:")
	// items = nil
	// for expr, tv := range info.Types {
	// 	var buf strings.Builder
	// 	posn := expr.Pos()
	// 	tvstr := tv.Type.String()
	// 	if tv.Value != nil {
	// 		tvstr += " = " + tv.Value.String()
	// 	}
	// 	// line:col | expr | mode : type = value
	// 	fmt.Fprintf(&buf, "%2d:%2d | %-19s | %-7s : %s",
	// 		posn.Line(), posn.Col(), types2.ExprString(expr),
	// 		mode(tv), tvstr)
	// 	items = append(items, buf.String())
	// }
	// sort.Strings(items)
	// fmt.Println(strings.Join(items, "\n"))

	// Output:
	// InitOrder: [c = "hello" b = S(c) a = len(b)]
	//
	// Defs and Uses of each named object:
	// builtin len:
	//   defined at <unknown position>
	//   used at 6:15
	// func fib(x int) int:
	//   defined at fib:8:6
	//   used at 12:20, 12:9
	// type S string:
	//   defined at fib:4:6
	//   used at 6:23
	// type int:
	//   defined at <unknown position>
	//   used at 8:12, 8:17
	// type string:
	//   defined at <unknown position>
	//   used at 4:8
	// var b S:
	//   defined at fib:6:8
	//   used at 6:19
	// var c string:
	//   defined at fib:6:11
	//   used at 6:25
	// var x int:
	//   defined at fib:8:10
	//   used at 10:10, 12:13, 12:24, 9:5
}
