// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// 次の例は、Valueを使用して定期的なプログラム設定の更新と
// 変更のワーカーゴルーチンへの伝播を行う方法を示しています。
func ExampleValue_config() {
	var config atomic.Value // 現在のサーバー設定を保持します
	// 初期設定値を作成し、configに格納します。
	config.Store(loadConfig())
	go func() {
		// 10秒ごとに設定を再読み込みし、
		// 新しいバージョンで設定値を更新します。
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	// 最新の設定値を使用して受信リクエストを処理する
	// ワーカーゴルーチンを作成します。
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				// 設定cを使用してリクエストrを処理します。
				_, _ = r, c
			}
		}()
	}
}

// 次の例は、コピーオンライトのイディオムを使用して、
// 頻繁に読み取られるが、あまり更新されないデータ構造を維持する方法を示しています。
func ExampleValue_readMostly() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex // 書き込み者のみが使用します
	// read関数は、さらなる同期化なしでデータを読み取るために使用できます
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	// insert関数は、さらなる同期化なしでデータを更新するために使用できます
	insert := func(key, val string) {
		mu.Lock() // 他の潜在的な書き込み者と同期します
		defer mu.Unlock()
		m1 := m.Load().(Map) // データ構造の現在の値をロードします
		m2 := make(Map)      // 新しい値を作成します
		for k, v := range m1 {
			m2[k] = v // 現在のオブジェクトから新しいオブジェクトにすべてのデータをコピーします
		}
		m2[key] = val // 必要な更新を行います
		m.Store(m2)   // 現在のオブジェクトを新しいオブジェクトとアトミックに置き換えます
		// この時点で、すべての新しい読み取り者は新しいバージョンで作業を開始します。
		// 既存の読み取り者（もしあれば）がそれを使用し終えると、古いバージョンはガベージコレクションされます。
	}
	_, _ = read, insert
}
