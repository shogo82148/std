// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// AlertErrorはTLSアラートです。
//
// QUICトランスポートを使用する場合、QUICConnのメソッドは
// TLSアラートを送信する代わりにAlertErrorをラップしたエラーを返します。
type AlertError uint8

var _ error = AlertError(0)

func (e AlertError) Error() string
