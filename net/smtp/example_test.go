// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/smtp"
)

func Example() {
	// リモートSMTPサーバーに接続する。
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// まず、送信者と受信者を設定します
	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("recipient@example.net"); err != nil {
		log.Fatal(err)
	}

	// メール本文を送信する。
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// QUITコマンドを送信し、接続を閉じます。
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}

func ExamplePlainAuth() {
	// hostnameはPlainAuthによってTLS証明書の検証に使用されます。
	hostname := "mail.example.com"
	auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSendMail() {
	// 認証情報を設定する。
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// サーバーに接続し、認証し、送信元と受信者を設定し、
	// メールを一括で送信します。
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
