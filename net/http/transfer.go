// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/net/http/internal"
)

// ErrLineTooLongは、不正なチャンクエンコーディングでリクエストまたはレスポンスボディを読み取る場合に返されます。
var ErrLineTooLong = internal.ErrLineTooLong

<<<<<<< HEAD
// ErrBodyReadAfterCloseは、ボディが閉じられた後にRequestまたはResponseのボディを読み取る場合に返されます。
// これは通常、HTTPハンドラがResponseWriterのWriteHeaderまたはWriteを呼び出した後にボディが読み取られた場合に発生します。
=======
// ErrBodyReadAfterClose is returned when reading a [Request] or [Response]
// Body after the body has been closed. This typically happens when the body is
// read after an HTTP [Handler] calls WriteHeader or Write on its
// [ResponseWriter].
>>>>>>> upstream/release-branch.go1.22
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
