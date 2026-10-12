// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// FcntlFlockは、[F_GETLK]、[F_SETLK]、[F_SETLKW] コマンドに対してfcntlシステムコールを実行します。
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error
