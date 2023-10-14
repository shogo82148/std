// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (js && wasm) || plan9

// On plan9, per http://9p.io/magic/man2html/2/access: “Since file permissions
// are checked by the server and group information is not known to the client,
// access must open the file to check permissions.”
//
// js,wasm is similar, in that it does not define syscall.Access.

package modload
