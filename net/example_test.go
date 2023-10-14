// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net_test

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/time"
)

func ExampleListener() {
	// ローカルシステムの利用可能なすべてのユニキャストおよび
	// IPv4マルチキャストアドレス上のTCPポート2000でリッスンします。
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 新しいgoroutineで接続を処理する。
		// ループはその後、受け付けを再開するため、
		// 複数の接続を同時に処理することが可能です。
		go func(c net.Conn) {
			// すべての受信データをエコーします。
			io.Copy(c, c)
			// 接続をシャットダウンします。
			c.Close()
		}(conn)
	}
}

func ExampleDialer() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}
}

func ExampleDialer_unix() {
	// DialUnixはcontext.Contextパラメータを受け取りません。この例は、
	// Contextを使用してUnixソケットにダイヤルする方法を示しています。ただし、
	// Contextはダイヤル操作にのみ適用されます。接続が確立された後には適用されません。
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	d.LocalAddr = nil // もしローカルのアドレスがあれば、ここに追加してください。
	raddr := net.UnixAddr{Name: "/path/to/unix.sock", Net: "unix"}
	conn, err := d.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("Hello, socket!")); err != nil {
		log.Fatal(err)
	}
}

func ExampleIPv4() {
	fmt.Println(net.IPv4(8, 8, 8, 8))

	// Output:
	// 8.8.8.8
}

func ExampleParseCIDR() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

	// Output:
	// 192.0.2.1
	// 192.0.2.0/24
	// 2001:db8:a0b:12f0::1
	// 2001:db8::/32
}

func ExampleParseIP() {
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))

	// Output:
	// 192.0.2.1
	// 2001:db8::68
	// <nil>
}

func ExampleIP_DefaultMask() {
	ip := net.ParseIP("192.0.2.1")
	fmt.Println(ip.DefaultMask())

	// Output:
	// ffffff00
}

func ExampleIP_Equal() {
	ipv4DNS := net.ParseIP("8.8.8.8")
	ipv4Lo := net.ParseIP("127.0.0.1")
	ipv6DNS := net.ParseIP("0:0:0:0:0:FFFF:0808:0808")

	fmt.Println(ipv4DNS.Equal(ipv4DNS))
	fmt.Println(ipv4DNS.Equal(ipv4Lo))
	fmt.Println(ipv4DNS.Equal(ipv6DNS))

	// Output:
	// true
	// false
	// true
}

func ExampleIP_IsGlobalUnicast() {
	ipv6Global := net.ParseIP("2000::")
	ipv6UniqLocal := net.ParseIP("2000::")
	ipv6Multi := net.ParseIP("FF00::")

	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("8.8.8.8")
	ipv4Broadcast := net.ParseIP("255.255.255.255")

	fmt.Println(ipv6Global.IsGlobalUnicast())
	fmt.Println(ipv6UniqLocal.IsGlobalUnicast())
	fmt.Println(ipv6Multi.IsGlobalUnicast())

	fmt.Println(ipv4Private.IsGlobalUnicast())
	fmt.Println(ipv4Public.IsGlobalUnicast())
	fmt.Println(ipv4Broadcast.IsGlobalUnicast())

	// Output:
	// true
	// true
	// false
	// true
	// true
	// false
}

func ExampleIP_IsInterfaceLocalMulticast() {
	ipv6InterfaceLocalMulti := net.ParseIP("ff01::1")
	ipv6Global := net.ParseIP("2000::")
	ipv4 := net.ParseIP("255.0.0.0")

	fmt.Println(ipv6InterfaceLocalMulti.IsInterfaceLocalMulticast())
	fmt.Println(ipv6Global.IsInterfaceLocalMulticast())
	fmt.Println(ipv4.IsInterfaceLocalMulticast())

	// Output:
	// true
	// false
	// false
}

func ExampleIP_IsLinkLocalMulticast() {
	ipv6LinkLocalMulti := net.ParseIP("ff02::2")
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")

	fmt.Println(ipv6LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv6LinkLocalUni.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalMulticast())

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIP_IsLinkLocalUnicast() {
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv6Global := net.ParseIP("2000::")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")

	fmt.Println(ipv6LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv6Global.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalUnicast())

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIP_IsLoopback() {
	ipv6Lo := net.ParseIP("::1")
	ipv6 := net.ParseIP("ff02::1")
	ipv4Lo := net.ParseIP("127.0.0.0")
	ipv4 := net.ParseIP("128.0.0.0")

	fmt.Println(ipv6Lo.IsLoopback())
	fmt.Println(ipv6.IsLoopback())
	fmt.Println(ipv4Lo.IsLoopback())
	fmt.Println(ipv4.IsLoopback())

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIP_IsMulticast() {
	ipv6Multi := net.ParseIP("FF00::")
	ipv6LinkLocalMulti := net.ParseIP("ff02::1")
	ipv6Lo := net.ParseIP("::1")
	ipv4Multi := net.ParseIP("239.0.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4Lo := net.ParseIP("127.0.0.0")

	fmt.Println(ipv6Multi.IsMulticast())
	fmt.Println(ipv6LinkLocalMulti.IsMulticast())
	fmt.Println(ipv6Lo.IsMulticast())
	fmt.Println(ipv4Multi.IsMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsMulticast())
	fmt.Println(ipv4Lo.IsMulticast())

	// Output:
	// true
	// true
	// false
	// true
	// true
	// false
}

func ExampleIP_IsPrivate() {
	ipv6Private := net.ParseIP("fc00::")
	ipv6Public := net.ParseIP("fe00::")
	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("11.0.0.0")

	fmt.Println(ipv6Private.IsPrivate())
	fmt.Println(ipv6Public.IsPrivate())
	fmt.Println(ipv4Private.IsPrivate())
	fmt.Println(ipv4Public.IsPrivate())

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIP_IsUnspecified() {
	ipv6Unspecified := net.ParseIP("::")
	ipv6Specified := net.ParseIP("fe00::")
	ipv4Unspecified := net.ParseIP("0.0.0.0")
	ipv4Specified := net.ParseIP("8.8.8.8")

	fmt.Println(ipv6Unspecified.IsUnspecified())
	fmt.Println(ipv6Specified.IsUnspecified())
	fmt.Println(ipv4Unspecified.IsUnspecified())
	fmt.Println(ipv4Specified.IsUnspecified())

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIP_Mask() {
	ipv4Addr := net.ParseIP("192.0.2.1")
	// このマスクはIPv4の/24サブネットに対応しています。
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))

	ipv6Addr := net.ParseIP("2001:db8:a0b:12f0::1")
	// このマスクはIPv6の/32サブネットに対応しています。
	ipv6Mask := net.CIDRMask(32, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))

	// Output:
	// 192.0.2.0
	// 2001:db8::
}

func ExampleIP_String() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.String())
	fmt.Println(ipv4.String())

	// Output:
	// fc00::
	// 10.255.0.0
}

func ExampleIP_To16() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.To16())
	fmt.Println(ipv4.To16())

	// Output:
	// fc00::
	// 10.255.0.0
}

func ExampleIP_to4() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.To4())
	fmt.Println(ipv4.To4())

	// Output:
	// <nil>
	// 10.255.0.0
}

func ExampleCIDRMask() {
	// このマスクはIPv4の/31サブネットに対応しています。
	fmt.Println(net.CIDRMask(31, 32))

	// このマスクはIPv6の/64サブネットに対応しています。
	fmt.Println(net.CIDRMask(64, 128))

	// Output:
	// fffffffe
	// ffffffffffffffff0000000000000000
}

func ExampleIPv4Mask() {
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))

	// Output:
	// ffffff00
}

func ExampleUDPConn_WriteTo() {
	// Dialとは異なり、ListenPacketはピアとの関連付けなしで
	// 接続を作成します。
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", "192.0.2.1:2000")
	if err != nil {
		log.Fatal(err)
	}

	// この接続は、指定したアドレスにデータを書き込むことができます。
	_, err = conn.WriteTo([]byte("data"), dst)
	if err != nil {
		log.Fatal(err)
	}
}
