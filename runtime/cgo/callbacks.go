// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//go:linkname _runtime_cgo_panic_internal runtime._cgo_panic_internal

//go:cgo_import_static x_cgo_init
//go:linkname x_cgo_init x_cgo_init
//go:linkname _cgo_init _cgo_init

//go:cgo_import_static x_cgo_malloc
//go:linkname x_cgo_malloc x_cgo_malloc
//go:linkname _cgo_malloc _cgo_malloc

//go:cgo_import_static x_cgo_free
//go:linkname x_cgo_free x_cgo_free
//go:linkname _cgo_free _cgo_free

//go:cgo_import_static x_cgo_thread_start
//go:linkname x_cgo_thread_start x_cgo_thread_start
//go:linkname _cgo_thread_start _cgo_thread_start

//go:cgo_import_static x_cgo_sys_thread_create
//go:linkname x_cgo_sys_thread_create x_cgo_sys_thread_create
//go:linkname _cgo_sys_thread_create _cgo_sys_thread_create

//go:cgo_import_static x_cgo_notify_runtime_init_done
//go:linkname x_cgo_notify_runtime_init_done x_cgo_notify_runtime_init_done
//go:linkname _cgo_notify_runtime_init_done _cgo_notify_runtime_init_done
