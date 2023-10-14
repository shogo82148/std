// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuzz

import (
	"github.com/shogo82148/std/context"
)

// RunFuzzWorker is called in a worker process to communicate with the
// coordinator process in order to fuzz random inputs. RunFuzzWorker loops
// until the coordinator tells it to stop.
//
// fn is a wrapper on the fuzz function. It may return an error to indicate
// a given input "crashed". The coordinator will also record a crasher if
// the function times out or terminates the process.
//
// RunFuzzWorker returns an error if it could not communicate with the
// coordinator process.
func RunFuzzWorker(ctx context.Context, fn func(CorpusEntry) error) error
