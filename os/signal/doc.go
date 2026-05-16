// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package signalは、受信シグナルへのアクセスを実装します。

シグナルは主にUnix系システムで使われます。WindowsとPlan 9での
このパッケージの使い方については、以下を参照してください。

# Types of signals

SIGKILL と SIGSTOP はプログラムによって捕捉できないため、
このパッケージの影響を受けません。

同期シグナルとは、プログラム実行中のエラーによって発生する
シグナルです。SIGBUS、SIGFPE、SIGSEGV がこれに当たります。
これらは、[os.Process.Kill] や kill プログラム、あるいは同様の
仕組みで送られた場合ではなく、プログラム実行によって発生した場合に
のみ同期シグナルと見なされます。一般に、以下で述べる場合を除き、
Goプログラムは同期シグナルを実行時panicに変換します。

残りのシグナルは非同期シグナルです。これらはプログラムのエラーに
よって発生するのではなく、カーネルまたは他のプログラムから送られます。

非同期シグナルのうち、SIGHUP はプログラムが制御端末を失ったときに
送られます。SIGINT は、制御端末上のユーザーが割り込み文字を押したときに
送られます。既定ではその文字は ^C（Control-C）です。SIGQUIT は、
制御端末上のユーザーが終了文字を押したときに送られます。既定ではその文字は
^\（Control-Backslash）です。一般に ^C を押すとプログラムを単純に終了させる
ことができ、^\ を押すとスタックダンプ付きで終了させることができます。

# Default behavior of signals in Go programs

既定では、同期シグナルは実行時panicに変換されます。SIGHUP、SIGINT、SIGTERM
はプログラムを終了させます。SIGQUIT、SIGILL、SIGTRAP、SIGABRT、SIGSTKFLT、
SIGEMT、SIGSYS はスタックダンプ付きでプログラムを終了させます。
SIGTSTP、SIGTTIN、SIGTTOU はシステムの既定動作になります（これらのシグナルは
シェルのジョブ制御で使われます）。SIGPROF は runtime.CPUProfile を実装するために
Goランタイムによって直接処理されます。その他のシグナルは捕捉されますが、
何の処理も行われません。

Goプログラムが SIGHUP または SIGINT を無視する設定（signal handler を SIG_IGN に
設定）で起動された場合、それらは無視されたままになります。

Goプログラムが空でないシグナルマスクを持って起動された場合、通常はその設定が
尊重されます。ただし、一部のシグナルは明示的にブロック解除されます。すなわち、
同期シグナル、SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF、そしてLinuxでは
32（SIGCANCEL）と33（SIGSETXID）です（SIGCANCEL と SIGSETXID は glibc が内部で
使います）。[os.Exec] または [os/exec] によって起動されたサブプロセスは、
変更後のシグナルマスクを継承します。

# Changing the behavior of signals in Go programs

このパッケージの関数を使うと、Goプログラムのシグナル処理方法を変更できます。

Notify は、指定された一連の非同期シグナルに対する既定動作を無効にし、
代わりに1つ以上の登録済みチャネルへそれらを配送します。具体的には、
SIGHUP、SIGINT、SIGQUIT、SIGABRT、SIGTERM に適用されます。また、
SIGTSTP、SIGTTIN、SIGTTOU といったジョブ制御シグナルにも適用され、この場合は
システムの既定動作は起こりません。さらに、通常は何も起こさない一部のシグナル、
すなわち SIGUSR1、SIGUSR2、SIGPIPE、SIGALRM、SIGCHLD、SIGCONT、SIGURG、
SIGXCPU、SIGXFSZ、SIGVTALRM、SIGWINCH、SIGIO、SIGPWR、SIGINFO、SIGTHR、
SIGWAITING、SIGLWP、SIGFREEZE、SIGTHAW、SIGLOST、SIGXRES、SIGJVM1、SIGJVM2、
およびシステムで使われる任意のリアルタイムシグナルにも適用されます。
これらのシグナルのすべてが全システムで利用できるわけではないことに注意してください。

プログラムが SIGHUP または SIGINT を無視する設定で起動されていて、[Notify] が
どちらかのシグナルに対して呼ばれると、そのシグナル用の signal handler が
インストールされ、もはや無視されなくなります。その後 [Reset] または [Ignore] が
そのシグナルに対して呼ばれるか、あるいはそのシグナルに対して Notify に渡したすべての
チャネルに対して [Stop] が呼ばれると、そのシグナルは再び無視されます。Reset は
そのシグナルのシステム既定動作を復元し、Ignore はそのシグナルを完全に無視するようにします。

プログラムが空でないシグナルマスクで起動された場合、前述のとおり一部のシグナルは
明示的にブロック解除されます。ブロックされたシグナルに対して Notify が呼ばれると、
そのシグナルはブロック解除されます。その後、そのシグナルに対して Reset が呼ばれるか、
またはそのシグナルに対して Notify に渡したすべてのチャネルに対して Stop が呼ばれると、
そのシグナルは再びブロックされます。

# SIGPIPE

Goプログラムが壊れたパイプに書き込むと、カーネルは SIGPIPE シグナルを発生させます。

プログラムが SIGPIPE シグナルを受け取るための Notify を呼んでいない場合、
動作はファイルディスクリプタ番号に依存します。ファイルディスクリプタ1または2
（標準出力または標準エラー）で壊れたパイプに書き込むと、SIGPIPE シグナルにより
プログラムは終了します。その他のファイルディスクリプタで壊れたパイプに書き込んだ場合、
SIGPIPE シグナルに対しては何の処理も行われず、書き込みは [syscall.EPIPE] エラーで失敗します。

プログラムが SIGPIPE シグナルを受け取るために Notify を呼んでいる場合は、
ファイルディスクリプタ番号は関係ありません。SIGPIPE シグナルは Notify チャネルへ
配送され、書き込みは [syscall.EPIPE] エラーで失敗します。

つまり既定では、コマンドラインプログラムは典型的なUnixのコマンドラインプログラムと
同じように動作し、その他のプログラムは閉じたネットワーク接続への書き込みで SIGPIPE によって
クラッシュしません。

# Go programs that use cgo or SWIG

非Goコード、通常は cgo または SWIG でアクセスされるC/C++コードを含むGoプログラムでは、
通常はGoの起動コードが最初に実行されます。非Goの起動コードが実行される前に、
Goランタイムが期待するように signal handler を設定します。非Goの起動コードが独自の
signal handler をインストールしたい場合、Goを正常に動かし続けるためにいくつかの手順を
踏む必要があります。この節ではその手順と、非Goコードによる signal handler 設定の変更が
Goプログラムへ与える全体的な影響を説明します。まれに、Goコードより先に非Goコードが
実行されることがあり、その場合は次の節も適用されます。

Goプログラムから呼ばれる非Goコードが signal handler やマスクを変更しない場合、
動作は純粋なGoプログラムと同じです。

非Goコードが signal handler をインストールする場合は、sigaction で SA_ONSTACK フラグを
使わなければなりません。これを怠ると、シグナル受信時にプログラムがクラッシュする可能性が
高くなります。Goプログラムは通常、制限されたスタックで実行されるため、別の signal stack を
設定します。

非Goコードが同期シグナル（SIGBUS、SIGFPE、SIGSEGV）のいずれかに対する signal handler を
インストールする場合、既存のGo signal handler を記録しておくべきです。これらのシグナルが
Goコード実行中に発生した場合は、Go signal handler を呼び出すべきです（シグナルがGoコード
実行中に発生したかどうかは、signal handler に渡された PC を見れば判断できます）。そうしないと、
一部のGo実行時panicが期待どおりに発生しません。

非Goコードが非同期シグナルのいずれかに対する signal handler をインストールする場合、
Go signal handler を呼び出すかどうかは任意です。もちろん、Go signal handler を呼び出さなければ、
上で説明したGoの動作は発生しません。特に SIGPROF シグナルでは問題になることがあります。

非Goコードは、Goランタイムによって作成されたスレッドの signal mask を変更すべきではありません。
非Goコード自身が新しいスレッドを開始する場合、そのスレッドは好きなように signal mask を設定できます。

非Goコードが新しいスレッドを開始し、signal mask を変更し、そのスレッドでGo関数を呼び出すと、
Goランタイムは自動的に一部のシグナルをブロック解除します。すなわち、同期シグナル、
SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF、SIGCANCEL、SIGSETXID です。
Go関数が戻ると、非Go側の signal mask は復元されます。

Go signal handler がGoコードを実行していない非Goスレッド上で呼び出されると、
通常は以下のように信号を非Goコードへ転送します。シグナルが SIGPROF の場合、
Go handler は何もしません。それ以外の場合、Go handler は自分自身を外し、
そのシグナルをブロック解除してもう一度 raise し、非Go handler または既定のシステム handler を
呼び出します。プログラムが終了しない場合、Go handler はその後再インストールされ、
プログラムの実行を継続します。

SIGPIPE シグナルを受信した場合、SIGPIPE がGoスレッド上で受信されたときは、Goプログラムは
上で説明した特別な処理を行います。SIGPIPE が非Goスレッド上で受信された場合は、
そのシグナルは（存在すれば）非Go handler に転送されます。もし handler が存在しなければ、
既定のシステム handler によりプログラムは終了します。

# Non-Go programs that call Go code

-buildmode=c-shared のようなオプションでGoコードがビルドされると、既存の非Goプログラムの
一部として実行されます。Goコードが開始される時点で、非Goコードはすでに signal handler を
インストールしているかもしれません（cgo や SWIG を使う通常でないケースでも同様で、その場合は
ここでの説明が当てはまります）。-buildmode=c-archive の場合、Goランタイムはグローバルな
コンストラクタの時点でシグナルを初期化します。-buildmode=c-shared の場合、共有ライブラリが
ロードされたときにGoランタイムがシグナルを初期化します。

Goランタイムが SIGCANCEL または SIGSETXID シグナル（Linuxでのみ使われます）に対する
既存の signal handler を見つけた場合、SA_ONSTACK フラグを有効にし、それ以外はその signal handler
を保持します。

同期シグナルと SIGPIPE に対しては、Goランタイムが signal handler をインストールします。
既存の signal handler があれば保存します。非Goコード実行中に同期シグナルが到来した場合、
GoランタイムはGo signal handler の代わりに既存の signal handler を呼び出します。

-buildmode=c-archive または -buildmode=c-shared でビルドされたGoコードは、既定では
他の signal handler をインストールしません。既存の signal handler がある場合、Goランタイムは
SA_ONSTACK フラグを有効にし、それ以外はその signal handler を保持します。非同期シグナルに対して
Notify が呼ばれると、そのシグナル用にGo signal handler がインストールされます。その後、その
シグナルに対して Reset が呼ばれると、元の処理が再インストールされ、存在する場合は非Goの
signal handler が復元されます。

-buildmode=c-archive でも -buildmode=c-shared でもなくビルドされたGoコードは、上で列挙した
非同期シグナルに対する signal handler をインストールし、既存の signal handler を保存します。
シグナルが非Goスレッドへ配送された場合は上で説明したように動作しますが、既存の非Go signal
handler がある場合は、シグナルを raise する前にその handler がインストールされます。

# Windows

Windowsでは、通常 ^C（Control-C）または ^BREAK（Control-Break）でプログラムが終了します。
[os.Interrupt] に対して Notify が呼ばれている場合、^C または ^BREAK によって
[os.Interrupt] がチャネルへ送られ、プログラムは終了しません。Reset が呼ばれるか、
Notify に渡したすべてのチャネルに対して Stop が呼ばれると、既定の動作が復元されます。

さらに、Notify が呼ばれていて、Windows が CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT、
CTRL_SHUTDOWN_EVENT をプロセスへ送った場合、Notify は syscall.SIGTERM を返します。
Control-C や Control-Break と違って、CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT、
CTRL_SHUTDOWN_EVENT のいずれかを受信しても Notify はプロセスの動作を変更しません。
そのため、プロセスが終了しない限り、依然として終了されます。ただし、syscall.SIGTERM を
受信すれば、終了前にクリーンアップする機会が得られます。

# Plan 9

Plan 9 では、シグナルの型は syscall.Note で、文字列です。syscall.Note を指定して
Notify を呼ぶと、その文字列が note として投稿されたときに、その値がチャネルへ送られます。
*/
package signal
