// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package weak provides weak pointers with the goal of memory efficiency.
The primary use-cases for weak pointers are for implementing caches,
canonicalization maps (like the unique package), and for tying together
the lifetimes of separate values (for example, through a map with weak
keys).

# Advice

This package is intended to target niche use-cases like the unique
package, and the structures inside are not intended to be general
replacements for regular Go pointers, maps, etc.
Misuse of the structures in this package may generate unexpected and
hard-to-reproduce bugs.
Using the facilities in this package to try and resolve out-of-memory
issues requires careful consideration, and even so, will likely be the
wrong answer if the solution does not fall into one of the listed
use-cases above.

The structures in this package are intended to be an implementation
detail of the package they are used by (again, see the unique package).
If you're writing a package intended to be used by others, as a rule of
thumb, avoid exposing the behavior of any weak structures in your package's
API.
Doing so will almost certainly make your package more difficult to use
correctly.
*/
package weak
