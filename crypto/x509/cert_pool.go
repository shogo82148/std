// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// CertPoolは証明書のセットです。
type CertPool struct {
	byName map[string][]int

	// lazyCertsには、必要に応じて遅延的に解析/展開される証明書を返す関数が含まれています。
	lazyCerts []lazyCert

	// haveSum は sum224(cert.Raw) を true にマッピングします。これは、AddCert の重複検出のためにのみ使用されます。CertPool.contains の呼び出しを AddCert パスで避けるためです（なぜなら、contains メソッドは getCert を呼び出し、怠惰な getCert からの節約を否定することができるためです）。
	haveSum map[sum224]bool

	// systemPoolは、システムルートから派生した特別なプールであることを示します。追加のルートを含む場合、呼び出し元によって提供されたルートを使用して1つの検証、およびシステムプラットフォームの検証装置を使用してもう1つの検証が必要です。
	systemPool bool
}

// NewCertPoolは新しい、空のCertPoolを返します。
func NewCertPool() *CertPool

// Cloneはsのコピーを返します。
func (s *CertPool) Clone() *CertPool

// SystemCertPoolはシステム証明書プールのコピーを返します。
//
// macOS以外のUnixシステムでは、環境変数SSL_CERT_FILEとSSL_CERT_DIRを使用して、
// SSL証明書ファイルとSSL証明書ファイルのディレクトリのシステムのデフォルト場所を上書きすることができます。後者はコロンで区切られたリストになります。
//
// 返されたプールへの変更はディスクに書き込まれず、SystemCertPoolによって返される他のプールに影響を与えません。
//
// システム証明書プールの新しい変更は、後続の呼び出しで反映されない場合があります。
func SystemCertPool() (*CertPool, error)

// AddCertは証明書をプールに追加します。
func (s *CertPool) AddCert(cert *Certificate)

// AppendCertsFromPEMは、一連のPEMエンコードされた証明書を解析しようとします。
// 見つかった証明書をsに追加し、成功した証明書があるかどうかを報告します。
//
// 多くのLinuxシステムでは、/etc/ssl/cert.pemには、この関数に適した形式でシステム全体のルートCAセットが含まれています。
func (s *CertPool) AppendCertsFromPEM(pemCerts []byte) (ok bool)

// Subjectsはプール内のすべての証明書のDERエンコードされたサブジェクトのリストを返します。
//
// Deprecated: sが [SystemCertPool] から返された場合、Subjectsにはシステムルートは含まれません。
func (s *CertPool) Subjects() [][]byte

// Equalは、sとotherが等しいかどうかを報告します。
func (s *CertPool) Equal(other *CertPool) bool

// AddCertWithConstraint adds a certificate to the pool with the additional
// constraint. When Certificate.Verify builds a chain which is rooted by cert,
// it will additionally pass the whole chain to constraint to determine its
// validity. If constraint returns a non-nil error, the chain will be discarded.
// constraint may be called concurrently from multiple goroutines.
func (s *CertPool) AddCertWithConstraint(cert *Certificate, constraint func([]*Certificate) error)
