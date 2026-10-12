// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// browserbridge runs the crypto/internal/fips140test WebAssembly module in a
// web browser, driven from the host. It relays the BoringSSL acvptool module
// wrapper protocol for ACVP algorithm testing, and runs the FIPS 140-3
// functional tests, against the browser's WebAssembly engine.
//
// The bridge runs as a long-lived server on the host:
//
//	browserbridge -serve -wasm fips140test.test -wasmexec wasm_exec.js
//
// It prints a URL to open (once) in the browser under test. The page loads the
// WebAssembly module and waits for sessions. Each session instantiates a fresh
// module instance, equivalent to one process execution, with the argv and
// environment provided by a client.
//
// The browser may run on a different machine: the page is served at -addr and
// works over plain HTTP, the URL contains an unguessable token. Clients
// connect over a Unix socket, as acvptool runs on the same host as the
// bridge.
//
// Invoked with no flags, browserbridge acts as a module wrapper for acvptool:
//
//	ACVP_WRAPPER=1 acvptool -json vectors.json -wrapper ./browserbridge
//
// It relays its standard input and output to the module running in the browser,
// forwards the relevant environment variables, and exits with the module's
// exit code. It connects to the server over a fixed-name Unix socket in the
// user cache directory.
//
// The -run flag instead executes the module with the given arguments, to run
// the functional tests in the browser:
//
//	browserbridge -run -- -test.run 'TestIntegrityCheck|TestFIPS140' -test.v
//
// Some functional tests re-exec the test binary (TestCASTPasses,
// TestCASTFailures, TestIntegrityCheckFailure), which can't be done on
// js/wasm. Those run on the host with GOBROWSERBRIDGE pointing at this binary;
// they exec it (browserbridge -run) in place of themselves, so the re-exec'd
// module instance runs in the browser and the host test checks the relayed
// output. See crypto/internal/fips140test.
//
// TestIntegrityCheckFailure additionally passes -corrupt, which makes the
// bridge serve, for that one session, a module whose go:fipsinfo checksum has
// been overwritten, so the in-browser integrity check fails as it would for a
// tampered binary.
package main
