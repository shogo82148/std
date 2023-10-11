// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Hostnameは、カーネルが報告するホスト名を返します。
func Hostname() (name string, err error)
