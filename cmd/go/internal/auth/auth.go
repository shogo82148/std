// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package auth provides access to user-provided authentication credentials.
package auth

import (
	"github.com/shogo82148/std/net/http"
)

// AddCredentials populates the request header with the user's credentials
// as specified by the GOAUTH environment variable.
// It returns whether any matching credentials were found.
// req must use HTTPS or this function will panic.
func AddCredentials(client *http.Client, req *http.Request, prefix string) bool
