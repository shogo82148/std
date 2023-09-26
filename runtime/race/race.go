// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (race && linux && amd64) || (race && darwin && amd64) || (race && windows && amd64)
// +build race,linux,amd64 race,darwin,amd64 race,windows,amd64

package race

/*
void __tsan_init(void **racectx);
void __tsan_fini(void);
void __tsan_map_shadow(void *addr, void *size);
void __tsan_go_start(void *racectx, void **chracectx, void *pc);
void __tsan_go_end(void *racectx);
void __tsan_read(void *racectx, void *addr, void *pc);
void __tsan_write(void *racectx, void *addr, void *pc);
void __tsan_read_range(void *racectx, void *addr, long sz, long step, void *pc);
void __tsan_write_range(void *racectx, void *addr, long sz, long step, void *pc);
void __tsan_func_enter(void *racectx, void *pc);
void __tsan_func_exit(void *racectx);
void __tsan_malloc(void *racectx, void *p, long sz, void *pc);
void __tsan_free(void *p);
void __tsan_acquire(void *racectx, void *addr);
void __tsan_release(void *racectx, void *addr);
void __tsan_release_merge(void *racectx, void *addr);
void __tsan_finalizer_goroutine(void *racectx);
*/

func Initialize(racectx *uintptr)

func Finalize()

func MapShadow(addr, size uintptr)

func FinalizerGoroutine(racectx uintptr)

func Read(racectx uintptr, addr, pc uintptr)

func Write(racectx uintptr, addr, pc uintptr)

func ReadRange(racectx uintptr, addr, sz, step, pc uintptr)

func WriteRange(racectx uintptr, addr, sz, step, pc uintptr)

func FuncEnter(racectx uintptr, pc uintptr)

func FuncExit(racectx uintptr)

func Malloc(racectx uintptr, p, sz, pc uintptr)

func Free(p uintptr)

func GoStart(racectx uintptr, chracectx *uintptr, pc uintptr)

func GoEnd(racectx uintptr)

func Acquire(racectx uintptr, addr uintptr)

func Release(racectx uintptr, addr uintptr)

func ReleaseMerge(racectx uintptr, addr uintptr)
