// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9

package syslog_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/log/syslog"
)

func ExampleDial() {
	sysLog, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
	sysLog.Emerg("And this is a daemon emergency with demotag.")
}
