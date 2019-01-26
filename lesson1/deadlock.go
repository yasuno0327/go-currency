package lesson1

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

// MakeDeadLock デットロックを再現する関数
func MakeDeadLock() {
	var a, b value
	wg.Add(2) // 待つgoroutineの数を追加
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait() // goroutineが終了するまで待つ
}

// 最初のprintSumでv1にロックがかかると次のprintSumが実行された際v1のアクセスを飛ばしv2をロックするため
// printする際に相互に待ち状態になりデッドロックが発生する。
func printSum(v1, v2 *value) {
	defer wg.Done() // goroutineが終了したことを通知
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum=%v\n", v1.value+v2.value)
}
