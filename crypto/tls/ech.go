// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// ECHRejectionErrorは、リモートサーバーによってECHが拒否された場合に返される
// エラー型です。サーバーが再試行用のECHConfigListを提示した場合、
// RetryConfigListフィールドにそのリストが含まれます。
//
// クライアントは、RetryConfigsが空のECHRejectionErrorを、
// サーバーからの安全なシグナルとして扱うことができます。
type ECHRejectionError struct {
	RetryConfigList []byte
}

func (e *ECHRejectionError) Error() string
