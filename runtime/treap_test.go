// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

// Wrap the Treap one more time because go:notinheap doesn't
// actually follow a structure across package boundaries.
//
//go:notinheap
