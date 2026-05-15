// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for sanitizers. See runtime/cgo/sigaction.go.

<<<<<<< HEAD
//go:build (linux && (amd64 || arm64 || loong64 || ppc64le)) || (freebsd && amd64)
=======
//go:build (linux && (386 || amd64 || arm64 || loong64 || ppc64 || ppc64le)) || (freebsd && amd64)
>>>>>>> af3262e38f7b8a0f3d5b763c800723aeb5af8082

package runtime
