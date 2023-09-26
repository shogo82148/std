// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

const InfCallstackSource = `
package main
import "C"
import "time"

func loop() {
        for i := 0; i < 1000; i++ {
                time.Sleep(time.Millisecond*5)
        }
}

func main() {
        go loop()
        time.Sleep(time.Second * 1)
}
`
