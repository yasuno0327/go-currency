# Go言語、並行処理メモ1

## 競合状態
２つの操作が正しい状態で実行されないといけない場面で順番が保証されていない状態
共有リソースに対する排他的アクセスが必要な場所 => クリティカルセクション
### 解決策
1. sync.Mutexを使う事によるメモリの排他的アクセスの取得

## アトミック
操作されている特定のコンテキストの中でそれ以上分割不可能な操作
コンテキストとは特定の操作が含まれるスコープ（範囲）のこと

## デットロック
ロックのかかっているリソースにおじいにアクセスしようとし永遠にリソース解放待ちになる状態
下記Coffman条件を全て満たすとデットロックになる
つまり下記条件のうちいずれか一つを偽とすればデットロックは解消される

### Coffman条件
1. 相互排他
  ある並行プロセスがリソースに対して排他的な権利をどの時点においても保持している
2. 条件待ち
  ある平行プロセスはリソースの保持と追加のリソース待ちを同時に行わなければならない
3. 横取り不可
  ある並行プロセスによって保持されているリソースは、そのプロセスによってのみ解放される
4. 循環待ち
  ある並行プロセス(P1)は他の連なっている平行プロセス(P2)を待たなければならない
  そしてP2はP1を待っている

## ライブロック

## リソース枯渇