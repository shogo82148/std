// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql_test

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/database/sql"
	"github.com/shogo82148/std/flag"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/os/signal"
)

func Example_openDBCLI() {
	id := flag.Int64("id", 0, "person ID to find")
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}
	var err error

	// ドライバを開く場合、通常、データベースへの接続は試みられません。
	pool, err = sql.Open("driver-name", *dsn)
	if err != nil {

		// これは接続エラーではなく、DSN解析エラーや他の初期化エラーです。
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		stop()
	}()

	Ping(ctx)

	Query(ctx, *id)
}
