// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql_test

import (
	"github.com/shogo82148/std/database/sql"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
)

func Example_openDBService() {
	// ドライバを開くと、通常はデータベースに接続しようとは試みません。
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {

		// これは接続エラーではなく、DSNの解析エラーや他の初期化エラーです。
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &Service{db: db}

	http.ListenAndServe(":8080", s)
}
