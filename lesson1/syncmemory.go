package lesson1

import (
	"fmt"
	"sync"
)

// SampleSyncMemory メモリへのアクセスを同期するサンプルプログラム
func SampleSyncMemory() {
	var mu = sync.Mutex{}
	var value int
	go func() {
		mu.Lock() //メモリの排他的アクセスを取得
		value++
		mu.Unlock() // 排他的アクセスの解放
	}()
	mu.Lock() // 排他的アクセス取得
	// ここでメモリアクセスを同期した事により条件分岐でのvalueとPrintする際のvalueが違う値になっている事を防げる
	if value == 0 {
		PrintValue(value)
	} else {
		PrintValue(value)
	}
	mu.Unlock() // 解放
}

func PrintValue(value int) {
	fmt.Printf("the value is %v.\n", value)
}
