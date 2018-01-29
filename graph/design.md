# 設計

## DisjointSetがサポートするmethod

* MakeSet
  - 各nodeだけの集合を作る(初期化)
* Union
  - 2つの集合を結合する
* FindSet
  - edgeを入力に、そのedgeの両端のnodeが含まれる集合を返す


```go

type DisjointSet interface {
  MakeSet(g *graph) []Set
  Union()
  FindSet(e Edge)
}

```

## 状態

* []Set


1. MakeSetにより、各nodeの集合を作る
1. edgeを順番に選ぶ
1. FindSetにより、2.で選んだedgeの両端のnodeが含まれる集合を見つける
1. 3.で見つけた2つの集合が同じか、異なるかを判定する
    - 同じ: cycleあり
    - 異なる かつ 全edgeを探索終了: cycleなし




