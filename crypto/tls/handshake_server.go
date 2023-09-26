// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// serverHandshakeState contains details of a server handshake in progress.
// It's discarded once the handshake has completed.

// suppVersArray is the backing array of ClientHelloInfo.SupportedVersions
