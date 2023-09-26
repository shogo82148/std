// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// hairpinConn is a net.Conn that makes a “hairpin” call when closed, back into
// the tls.Conn which is calling it.
