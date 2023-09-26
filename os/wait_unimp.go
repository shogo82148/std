// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// aix, darwin, js/wasm, openbsd and solaris don't implement
// waitid/wait6. netbsd implements wait6, but that is causing test
// failures, see issue #48789.

//go:build aix || darwin || (js && wasm) || openbsd || solaris

package os
