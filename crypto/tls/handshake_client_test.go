// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in client mode:
//   % go test -test.run "TestRunClient" -connect
// The recorded bytes are written to stdout.
//
// The server private key is:
// -----BEGIN RSA PRIVATE KEY-----
// MIIBPAIBAAJBAJ+zw4Qnlf8SMVIPFe9GEcStgOY2Ww/dgNdhjeD8ckUJNP5VZkVD
// TGiXav6ooKXfX3j/7tdkuD8Ey2//Kv7+ue0CAwEAAQJAN6W31vDEP2DjdqhzCDDu
// OA4NACqoiFqyblo7yc2tM4h4xMbC3Yx5UKMN9ZkCtX0gzrz6DyF47bdKcWBzNWCj
// gQIhANEoojVt7hq+SQ6MCN6FTAysGgQf56Q3TYoJMoWvdiXVAiEAw3e3rc+VJpOz
// rHuDo6bgpjUAAXM+v3fcpsfZSNO6V7kCIQCtbVjanpUwvZkMI9by02oUk9taki3b
// PzPfAfNPYAbCJQIhAJXNQDWyqwn/lGmR11cqY2y9nZ1+5w3yHGatLrcDnQHxAiEA
// vnlEGo8K85u+KwIOimM48ZG8oTk7iFdkqLJR1utT3aU=
// -----END RSA PRIVATE KEY-----
//
// and certificate is:
// -----BEGIN CERTIFICATE-----
// MIICKzCCAdWgAwIBAgIJALE1E2URIMWSMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
// BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
// aWRnaXRzIFB0eSBMdGQwHhcNMTIwNDA2MTcxMDEzWhcNMTUwNDA2MTcxMDEzWjBF
// MQswCQYDVQQGEwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50
// ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ+z
// w4Qnlf8SMVIPFe9GEcStgOY2Ww/dgNdhjeD8ckUJNP5VZkVDTGiXav6ooKXfX3j/
// 7tdkuD8Ey2//Kv7+ue0CAwEAAaOBpzCBpDAdBgNVHQ4EFgQUeKaXmmO1xaGlM7oi
// fCNuWxt6zCswdQYDVR0jBG4wbIAUeKaXmmO1xaGlM7oifCNuWxt6zCuhSaRHMEUx
// CzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRl
// cm5ldCBXaWRnaXRzIFB0eSBMdGSCCQCxNRNlESDFkjAMBgNVHRMEBTADAQH/MA0G
// CSqGSIb3DQEBBQUAA0EAhTZAc8G7GtrUWZ8tonAxRnTsg26oyDxRrzms7EC86CJG
// HZnWRiok1IsFCEv7NRFukrt3uuQSu/TIXpyBqJdgTA==
// -----END CERTIFICATE-----

// $ openssl s_server -tls1_2 -cert server.crt -key server.key \
//     -cipher ECDHE-RSA-AES128-SHA -port 10443
// $ go test -test.run "TestRunClient" -connect -ciphersuites=0xc013 \
//     -minversion=0x0303 -maxversion=0x0303

// $ openssl s_server -tls1_2 -cert server.crt -key server.key \
//     -port 10443 -verify 0
// $ go test -test.run "TestRunClient" -connect -ciphersuites=0xc02f \
//     -maxversion=0x0303

// Script of interaction with openssl implementation:
//
//   openssl s_server -cipher ECDHE-ECDSA-AES128-SHA \
//     -key server.key -cert server.crt -port 10443
//
// The values for this test are obtained by building and running in client mode:
//   % go test -test.run "TestRunClient" -connect -ciphersuites=0xc009
// The recorded bytes are written to stdout.
//
// The server private key is:
//
// -----BEGIN EC PARAMETERS-----
// BgUrgQQAIw==
// -----END EC PARAMETERS-----
// -----BEGIN EC PRIVATE KEY-----
// MIHcAgEBBEIBmIPpCa0Kyeo9M/nq5mHxeFIGlw+MqakWcvHu3Keo7xK9ZWG7JG3a
// XfS01efjqSZJvF2DoL+Sly4A5iBn0Me9mdegBwYFK4EEACOhgYkDgYYABADEoe2+
// mPkLSHM2fsMWVhEi8j1TwztNIT3Na3Xm9rDcmt8mwbyyh/ByMnyzZC8ckLzqaCMQ
// fv7jJcBIOmngKG3TNwDvBGLdDaCccGKD2IHTZDGqnpcxvZawaMCbI952ZD8aXH/p
// Eg5YWLZfcN2b2OrV1/XVzLm2nzBmW2aaIOIn5b/+Ow==
// -----END EC PRIVATE KEY-----
//
// and certificate is:
//
// -----BEGIN CERTIFICATE-----
// MIICADCCAWICCQC4vy1HoNLr9DAJBgcqhkjOPQQBMEUxCzAJBgNVBAYTAkFVMRMw
// EQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0
// eSBMdGQwHhcNMTIxMTIyMTUwNjMyWhcNMjIxMTIwMTUwNjMyWjBFMQswCQYDVQQG
// EwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50ZXJuZXQgV2lk
// Z2l0cyBQdHkgTHRkMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAxKHtvpj5C0hz
// Nn7DFlYRIvI9U8M7TSE9zWt15vaw3JrfJsG8sofwcjJ8s2QvHJC86mgjEH7+4yXA
// SDpp4Cht0zcA7wRi3Q2gnHBig9iB02Qxqp6XMb2WsGjAmyPedmQ/Glx/6RIOWFi2
// X3Ddm9jq1df11cy5tp8wZltmmiDiJ+W//jswCQYHKoZIzj0EAQOBjAAwgYgCQgGI
// ok/r4kXFSH0brPXtmJ2uR3DAXhu2L73xtk23YUDTEaLO7gt+kn7/dp3DO36lP876
// EOJZ7EctfKzaTpcOFaBv0AJCAU38vmcTnC0FDr0/o4wlwTMTgw2UBrvUN3r27HrJ
// hi7d1xFpf4V8Vt77MXgr5Md4Da7Lvp5ONiQxe2oPOZUSB48q
// -----END CERTIFICATE-----
