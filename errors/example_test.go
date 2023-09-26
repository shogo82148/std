// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func Example() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}
