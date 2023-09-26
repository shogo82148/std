// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in server mode:
//   % go test -run "TestRunServer" -serve
// and then:
//   % gnutls-cli --insecure --debug 100 -p 10443 localhost > /tmp/log 2>&1
//   % python parse-gnutls-cli-debug-log.py < /tmp/log
