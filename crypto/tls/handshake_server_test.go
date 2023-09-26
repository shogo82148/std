// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// recordingConn is a net.Conn that records the traffic that passes through it.
// WriteTo can be used to produce Go code that contains the recorded traffic.

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in server mode:
//   % go test -test.run "TestRunServer" -serve
// The recorded bytes are written to stdout.

// Generated using:
// $ go test -test.run TestRunServer -serve -ciphersuites=0xc00a
// $ openssl s_client -host 127.0.0.1 -port 10443 -cipher ECDHE-ECDSA-AES256-SHA

// $ go test -run TestRunServer -serve -clientauth 1 \
//     -ciphersuites=0xc011 -minversion=0x0303 -maxversion=0x0303

/* corresponding key for cert is:
-----BEGIN EC PARAMETERS-----
BgUrgQQAIw==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBkJN9X4IqZIguiEVKMqeBUP5xtRsEv4HJEtOpOGLELwO53SD78Ew8
k+wLWoqizS3NpQyMtrU8JFdWfj+C57UNkOugBwYFK4EEACOhgYkDgYYABACVjJF1
FMBexFe01MNvja5oHt1vzobhfm6ySD6B5U7ixohLZNz1MLvT/2XMW/TdtWo+PtAd
3kfDdq0Z9kUsjLzYHQFMH3CQRnZIi4+DzEpcj0B22uCJ7B0rxE4wdihBsmKo+1vx
+U56jb0JuK7qixgnTy5w/hOWusPTQBbNZU6sER7m8Q==
-----END EC PRIVATE KEY-----
*/
