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

// ErrBodyReadAfterCloseは、ボディが閉じられた後に [Request] または [Response] のボディを読み取る場合に返されます。
// これは通常、HTTP [Handler] が [ResponseWriter] のWriteHeaderまたはWriteを呼び出した後にボディが読み取られた場合に発生します。
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
