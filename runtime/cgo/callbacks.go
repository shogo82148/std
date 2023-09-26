// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//go:cgo_import_static x_cgo_init
//go:linkname x_cgo_init x_cgo_init
//go:linkname _cgo_init _cgo_init

//go:cgo_import_static x_cgo_thread_start
//go:linkname x_cgo_thread_start x_cgo_thread_start
//go:linkname _cgo_thread_start _cgo_thread_start

//go:cgo_import_static x_cgo_sys_thread_create
//go:linkname x_cgo_sys_thread_create x_cgo_sys_thread_create
//go:linkname _cgo_sys_thread_create _cgo_sys_thread_create

//go:cgo_import_static x_cgo_pthread_key_created
//go:linkname x_cgo_pthread_key_created x_cgo_pthread_key_created
//go:linkname _cgo_pthread_key_created _cgo_pthread_key_created

//go:cgo_import_static x_crosscall2_ptr
//go:linkname x_crosscall2_ptr x_crosscall2_ptr
//go:linkname _crosscall2_ptr _crosscall2_ptr

//go:linkname _set_crosscall2 runtime.set_crosscall2

//go:cgo_import_static x_cgo_bindm
//go:linkname x_cgo_bindm x_cgo_bindm
//go:linkname _cgo_bindm _cgo_bindm

//go:cgo_import_static x_cgo_notify_runtime_init_done
//go:linkname x_cgo_notify_runtime_init_done x_cgo_notify_runtime_init_done
//go:linkname _cgo_notify_runtime_init_done _cgo_notify_runtime_init_done

//go:cgo_import_static x_cgo_set_context_function
//go:linkname x_cgo_set_context_function x_cgo_set_context_function
//go:linkname _cgo_set_context_function _cgo_set_context_function

//go:cgo_import_static _cgo_yield
//go:linkname _cgo_yield _cgo_yield

//go:cgo_import_static x_cgo_getstackbound
//go:linkname x_cgo_getstackbound x_cgo_getstackbound
//go:linkname _cgo_getstackbound _cgo_getstackbound
