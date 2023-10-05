package mypackage

import (
	remove_me "fmt"
	_ "image/png"
)

// 非公開関数なので削除してほしい
var removeMeVar = remove_me.Print

// 公開変数なので残してほしい
var (
	// 公開変数その1
	ExportedVar1 = "exported"

	// 公開変数その2
	ExportedVar2 = "exported"
)

// 非公開関数なので削除してほしい
func removeMe() {
	removeMeVar("remove me")
}

// 公開関数なので残してほしい
func LeaveMe() {
	// 関数の中身は削除
	removeMe()
}
