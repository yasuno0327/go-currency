package lesson1

import "fmt"

// MakeConflict 競合状態を再現する簡単なプログラム
func MakeConflict() {
	var data int
	// goroutineを起動し並行処理でdataをインクリメントする
	go func() {
		data++
	}()
	// ここの条件分岐ではdataがインクリメントされている状態とdataが0の場合の二通りになるケースが考えられる。
	// 平行で実行されているプロセスは同期的には実行されないため
	if data == 0 {
		// ここではdataが1の場合とdataが0で出力される場合が考えられる
		// 条件分岐後にdataがインクリメントされる可能性を考えられるため
		fmt.Printf("the value is %v.\n", data)
	}
}
