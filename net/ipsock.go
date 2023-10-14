// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// SplitHostPortは、"host:port"、"host％zone:port"、"[host]:port"、または "[host%zone]:port" のネットワークアドレスをhostまたはhost％zoneとポートに分割します。
//
// ホストポート内のリテラルIPv6アドレスは、"[::1]:80"、"[::1％lo0]:80"のように角括弧で囲む必要があります。
//
// hostportパラメータ、およびhostとportの結果の詳細については、func Dialを参照してください。
func SplitHostPort(hostport string) (host, port string, err error)

// JoinHostPort はホストとポートを "ホスト:ポート" のネットワークアドレスに結合します。
// ホストがコロンを含んでいる場合、リテラルIPv6アドレスで見つかるように、JoinHostPortは "[host]:port" を返します。
//
// ホストとポートパラメータの説明については、func Dial を参照してください。
func JoinHostPort(host, port string) string
