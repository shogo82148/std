// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag_test

import (
	"github.com/shogo82148/std/flag"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/os"
)

func ExampleTextVar() {
	fs := flag.NewFlagSet("ExampleTextVar", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.TextVar(&ip, "ip", net.IPv4(192, 168, 0, 100), "`IP address` to parse")
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)

	// 256は有効なIPv4コンポーネントではありません
	ip = nil
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)

	// Output:
	// {ip: 127.0.0.1}
	//
	// invalid value "256.0.0.1" for flag -ip: invalid IP address: 256.0.0.1
	// Usage of ExampleTextVar:
	//   -ip IP address
	//     	IP address to parse (default 192.168.0.100)
	// {ip: <nil>}
}
