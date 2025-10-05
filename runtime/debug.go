// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// GOMAXPROCSは同時に実行可能なCPUの最大数を設定し、以前の設定値を返します。n < 1の場合は現在の設定を変更しません。
//
// # Default
//
// GOMAXPROCS環境変数が正の整数に設定されている場合、GOMAXPROCSはその値がデフォルトになります。
//
// それ以外の場合、Goランタイムは以下の組み合わせから適切なデフォルト値を選択します。
//   - マシン上の論理CPU数
//   - プロセスのCPUアフィニティマスク
//   - Linuxの場合、cgroup CPUクォータに基づくプロセスの平均CPUスループット制限（存在する場合）
//
// GODEBUG=containermaxprocs=0が設定されていて、GOMAXPROCSが環境変数で設定されていない場合、GOMAXPROCSは[runtime.NumCPU]の値がデフォルトになります。なお、GODEBUG=containermaxprocs=0はバージョン1.24以下では [default] です。
//
// # Updates
//
// Goランタイムは、論理CPU数、CPUアフィニティマスク、cgroupクォータの変更に基づいてデフォルト値を定期的に更新します。
// GOMAXPROCS環境変数を設定するか、GOMAXPROCSを呼び出してカスタム値を設定すると、自動更新は無効になります。
// デフォルト値と自動更新は [SetDefaultGOMAXPROCS] を呼び出すことで元に戻せます。
//
// GODEBUG=updatemaxprocs=0 が設定されている場合、GoランタイムはGOMAXPROCSの自動更新を行いません。
// なお、GODEBUG=updatemaxprocs=0 は言語バージョン1.24以下では [default] です。
//
// # Compatibility
//
// デフォルトのGOMAXPROCSの挙動は、スケジューラの改善に伴い変更される可能性があることに注意してください。特に以下の実装詳細に関しては変更される場合があります。
//
// # Implementation details
//
// デフォルトのGOMAXPROCSをcgroup経由で計算する場合、GoランタイムはcgroupのCPUクォータ／期間として「平均CPUスループット制限」を計算します。cgroup v2ではこれらの値はcpu.maxファイルから取得され、cgroup v1ではそれぞれcpu.cfs_quota_usとcpu.cfs_period_usから取得されます。コンテナランタイムでCPU制限を設定できる場合、この値は通常「CPUリミット」オプションに対応し、「CPUリクエスト」ではありません。
//
// Goランタイムは通常、論理CPU数、CPUアフィニティマスク数、cgroupのCPUスループット制限のうち最小値をデフォルトのGOMAXPROCSとして選択します。ただし、論理CPU数またはCPUアフィニティマスク数が2未満でない限り、GOMAXPROCSを2未満に設定することはありません。
//
// cgroupのCPUスループット制限が整数でない場合、Goランタイムは次の整数に切り上げます。
//
// GOMAXPROCSの更新は最大で1秒に1回、またはアプリケーションがアイドル状態の場合はそれ以下の頻度で行われます。
//
// [default]: https://go.dev/doc/godebug#default
func GOMAXPROCS(n int) int

// SetDefaultGOMAXPROCSは、GOMAXPROCSの設定をランタイムのデフォルト値に更新します。[GOMAXPROCS] で説明されている通り、GOMAXPROCS環境変数は無視されます。
//
// SetDefaultGOMAXPROCSは、GOMAXPROCS環境変数や以前の[GOMAXPROCS]呼び出しによって自動更新が無効化されている場合に、デフォルトの自動更新GOMAXPROCS動作を有効にするため、または論理CPU数・CPUアフィニティマスク・cgroupクォータの合計が変更されたことを呼び出し元が認識している場合に即座に更新を強制するために使用できます。
func SetDefaultGOMAXPROCS()

// NumCPUは現在のプロセスで使用可能な論理CPUの数を返します。
//
// 利用可能なCPUのセットはプロセスの起動時にオペレーティングシステムによって確認されます。
// プロセスの起動後にオペレーティングシステムのCPU割り当てに変更があっても、それは反映されません。
func NumCPU() int

// NumCgoCall は現在のプロセスによって行われた cgo 呼び出しの数を返します。
func NumCgoCall() int64

// NumGoroutineは現在存在するゴルーチンの数を返します。
func NumGoroutine() int
