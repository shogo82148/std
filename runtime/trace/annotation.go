// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/context"
)

// NewTaskはタスクの種類（taskType）でタスクインスタンスを作成し、
// タスクを持つContextとともに返します。
// もし入力のContextにタスクが含まれている場合、新しいタスクはそのサブタスクとなります。
//
// タスクタイプはタスクインスタンスを分類するために使用されます。Go実行トレーサのような
// 分析ツールは、システムに一意のタスクタイプが有限であると見なすことがあります。
//
// 返されるTaskの[Task.End]メソッドはタスクの終了をマークするために使用されます。
// トレースツールは、タスクの作成とEndメソッドの呼び出しの間の時間をタスクのレイテンシとして測定し、
// レイテンシの分布をタスクタイプごとに提供します。
// Endメソッドが複数回呼び出された場合、レイテンシの測定には最初の呼び出しぶんのみ使用されます。
//
// ctx、task := trace.NewTask(ctx, "awesomeTask")
// trace.WithRegion(ctx, "preparation", prepWork)
// // タスクの準備
// go func() {  // 別のゴルーチンでタスクの処理を続ける。
//
//	    defer task.End()
//	    trace.WithRegion(ctx, "remainingWork", remainingWork)
//	}()
func NewTask(pctx context.Context, taskType string) (ctx context.Context, task *Task)

// Taskは、ユーザー定義の論理的な操作をトレースするためのデータ型です。
type Task struct {
	id uint64
}

// End は [Task] によって表される操作の終了を示します。
func (t *Task) End()

// Logは与えられたカテゴリとメッセージでワンオフのイベントを送信します。
// カテゴリは空にすることができ、APIはシステム内に一握りのユニークなカテゴリしか存在しないと仮定します。
func Log(ctx context.Context, category, message string)

// Logfは[Log]と似ていますが、値は指定されたフォーマット仕様を使用して整形されます。
func Logf(ctx context.Context, category, format string, args ...any)

// WithRegionは、呼び出し元のgoroutineに関連付けられた領域を開始し、fnを実行し、その後領域を終了します。もしコンテキストにタスクがある場合、領域はそのタスクに関連付けられます。そうでない場合、領域はバックグラウンドのタスクにアタッチされます。
// regionTypeは領域を分類するために使用されるため、ユニークなregionTypeはごくわずかであるべきです。
func WithRegion(ctx context.Context, regionType string, fn func())

// StartRegionはリージョンを開始して返します。
// 戻り値となるリージョンの[Region.End]メソッドは、
// リージョンを開始した同じゴルーチンから呼び出す必要があります。
// 各ゴルーチン内では、リージョンはネストする必要があります。
// つまり、このリージョンを終了する前に、このリージョンよりも後に開始されたリージョンを終了する必要があります。
// 推奨される使用法は
//
//	defer trace.StartRegion(ctx, "myTracedRegion").End()
func StartRegion(ctx context.Context, regionType string) *Region

// Regionは、実行時間の区間がトレースされるコードの領域です。
type Region struct {
	id         uint64
	regionType string
}

// Endはトレースされたコード領域の終わりを示します。
func (r *Region) End()

// IsEnabled はトレースが有効かどうかを報告します。
// この情報はアドバイザリー(助言的)です。トレースの状態は
// この関数が返るまでに変更されている可能性があります。
func IsEnabled() bool
