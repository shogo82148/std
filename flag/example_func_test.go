// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag_test

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/flag"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/os"
)

func ExampleFunc() {
	fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.Func("ip", "`IP address` to parse", func(s string) error {
		ip = net.ParseIP(s)
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	})
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// 256は有効なIPv4コンポーネントではありません。
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// Output:
	// {ip: 127.0.0.1, loopback: true}
	//
	// invalid value "256.0.0.1" for flag -ip: could not parse IP
	// Usage of ExampleFunc:
	//   -ip IP address
	//     	IP address to parse
	// {ip: <nil>, loopback: false}
}

func ExampleBoolFunc() {
	fs := flag.NewFlagSet("ExampleBoolFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	fs.BoolFunc("log", "logs a dummy message", func(s string) error {
		fmt.Println("dummy message:", s)
		return nil
	})
	fs.Parse([]string{"-log"})
	fs.Parse([]string{"-log=0"})

	// Output:
	// dummy message: true
	// dummy message: 0
}
