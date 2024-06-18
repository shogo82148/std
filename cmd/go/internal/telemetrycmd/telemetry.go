// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package telemetrycmd implements the "go telemetry" command.
package telemetrycmd

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdTelemetry = &base.Command{
	UsageLine: "go telemetry [off|local|on]",
	Short:     "manage telemetry data and settings",
	Long: `Telemetry is used to manage Go telemetry data and settings.

Telemetry can be in one of three modes: off, local, or on.

When telemetry is in local mode, counter data is written to the local file
system, but will not be uploaded to remote servers.

When telemetry is off, local counter data is neither collected nor uploaded.

When telemetry is on, telemetry data is written to the local file system
and periodically sent to https://telemetry.go.dev/. Uploaded data is used to
help improve the Go toolchain and related tools, and it will be published as
part of a public dataset.

For more details, see https://telemetry.go.dev/privacy.
This data is collected in accordance with the Google Privacy Policy
(https://policies.google.com/privacy).

To view the current telemetry mode, run "go telemetry".
To disable telemetry uploading, but keep local data collection, run
"go telemetry local".
To enable both collection and uploading, run “go telemetry on”.
To disable both collection and uploading, run "go telemetry off".

See https://go.dev/doc/telemetry for more information on telemetry.
`,
	Run: runTelemetry,
}
