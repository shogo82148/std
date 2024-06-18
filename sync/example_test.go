// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/sync"
)

// この例では、複数のURLを同時にフェッチし、WaitGroupを使用して、すべてのフェッチが完了するまでブロックします。
func ExampleWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// WaitGroup のカウンターをインクリメントする。
		wg.Add(1)
		// URLを取得するために、ゴルーチンを起動します。
		go func(url string) {
			// ゴルーチンが完了したら、カウンタを減らす。
			defer wg.Done()
			// URLを取得する。
			http.Get(url)
		}(url)
	}
	// すべてのHTTPフェッチが完了するまで待ちます。
	wg.Wait()
}

func ExampleOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Only once
}

// This example uses OnceValue to perform an "expensive" computation just once,
// even when used concurrently.
func ExampleOnceValue() {
	once := sync.OnceValue(func() int {
		sum := 0
		for i := 0; i < 1000; i++ {
			sum += i
		}
		fmt.Println("Computed once:", sum)
		return sum
	})
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			const want = 499500
			got := once()
			if got != want {
				fmt.Println("want", want, "got", got)
			}
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Computed once: 499500
}

// This example uses OnceValues to read a file just once.
func ExampleOnceValues() {
	once := sync.OnceValues(func() ([]byte, error) {
		fmt.Println("Reading file once")
		return os.ReadFile("example_test.go")
	})
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			data, err := once()
			if err != nil {
				fmt.Println("error:", err)
			}
			_ = data // Ignore the data for this example
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Reading file once
}
