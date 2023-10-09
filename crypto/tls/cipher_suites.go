// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// CipherSuiteはTLSの暗号スイートです。このパッケージのほとんどの関数は、この型の代わりに暗号スイートのIDを受け入れて公開します。
type CipherSuite struct {
	ID   uint16
	Name string

	// Supported versionsは、この暗号スイートをネゴシエートできるTLSプロトコルのバージョンのリストです。
	SupportedVersions []uint16

	// Insecureは、そのプリミティブ、設計、または実装による既知のセキュリティ問題があるため、暗号スイートが安全ではない場合にtrueとなります。
	Insecure bool
}

// CipherSuitesは、このパッケージで現在実装されている暗号スイートのリストを返します。
// ただし、セキュリティ上の問題があるものはInsecureCipherSuitesによって返されます。
//
// このリストはIDでソートされています。このパッケージによって選択されるデフォルトの暗号スイートが、
// 静的なリストでは捉えることができないロジックに依存している場合があり、
// この関数によって返されるものと一致しない場合があります。
func CipherSuites() []*CipherSuite

// InsecureCipherSuitesは、現在このパッケージで実装されているセキュリティ上の問題を抱えた暗号スイートのリストを返します。
// ほとんどのアプリケーションは、このリストに含まれる暗号スイートを使用すべきではありません。代わりに、CipherSuitesで返されるスイートのみを使用するべきです。
func InsecureCipherSuites() []*CipherSuite

// CipherSuiteNameは渡された暗号スイートIDの標準名
// （例：「TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256」）を返します。
// もしこのパッケージで暗号スイートが実装されていない場合は、IDの値をフォールバック表現として返します。
func CipherSuiteName(id uint16) string

// このパッケージで実装されている、または実装されていた暗号スイートのIDのリストです。
//
// 詳細は、https://www.iana.org/assignments/tls-parameters/tls-parameters.xml を参照してください。
const (
	// TLS 1.0 - 1.2の暗号スイート。
	TLS_RSA_WITH_RC4_128_SHA                      uint16 = 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA                 uint16 = 0x000a
	TLS_RSA_WITH_AES_128_CBC_SHA                  uint16 = 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA                  uint16 = 0x0035
	TLS_RSA_WITH_AES_128_CBC_SHA256               uint16 = 0x003c
	TLS_RSA_WITH_AES_128_GCM_SHA256               uint16 = 0x009c
	TLS_RSA_WITH_AES_256_GCM_SHA384               uint16 = 0x009d
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA              uint16 = 0xc007
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          uint16 = 0xc009
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          uint16 = 0xc00a
	TLS_ECDHE_RSA_WITH_RC4_128_SHA                uint16 = 0xc011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA           uint16 = 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA            uint16 = 0xc013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA            uint16 = 0xc014
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       uint16 = 0xc023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256         uint16 = 0xc027
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256         uint16 = 0xc02f
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       uint16 = 0xc02b
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384         uint16 = 0xc030
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       uint16 = 0xc02c
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   uint16 = 0xcca8
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 uint16 = 0xcca9

	// TLS 1.3サイファースイート。
	TLS_AES_128_GCM_SHA256       uint16 = 0x1301
	TLS_AES_256_GCM_SHA384       uint16 = 0x1302
	TLS_CHACHA20_POLY1305_SHA256 uint16 = 0x1303

	// TLS_FALLBACK_SCSVは標準の暗号スイートではなく、クライアントがバージョンのフォールバックを行っていることを示すものです。RFC 7507を参照してください。
	TLS_FALLBACK_SCSV uint16 = 0x5600

	// 正しい_SHA256サフィックスを持つ対応する暗号スイートのためのレガシー名前
	// 互換性のために保持されています。
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305   = TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305 = TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
)
