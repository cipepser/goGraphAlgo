# graph

[![Build Status](https://travis-ci.org/cipepser/goGraphAlgo.svg?branch=master)](https://travis-ci.org/cipepser/goGraphAlgo)
[![Coverage Status](https://coveralls.io/repos/github/cipepser/goGraphAlgo/badge.svg?branch=master)](https://coveralls.io/github/cipepser/goGraphAlgo?branch=master)


グラフ理論のアルゴリズム向けのデータ構造をGolangで実装したものになります。

## How to Install

```
$ go get github.com/cipepser/goGraphAlgo/...
```

可視化もできるようにしているので、可視化したい場合は以下もインストールしてください。

```
$ brew install graphviz
$ go get github.com/awalterschulze/gographviz
```

## How to Use

基本的な操作として、以下ができます。

* 無向グラフ/有向グラフ
* 頂点(Vertex)の追加/削除
* 辺(Edge)の追加/削除

具体的な使い方を以下に示します。
いずれも`error`を返すようにしているのて適宜エラーチェックをしてください。

### 無向グラフの作成

```go
g := graph.NewGraph()
```
※内部で`isDirected`フィールドを持っていますが、ゼロ値が`false`のためデフォルトで無向グラフになります。

### 有向グラフの作成

```go
g := graph.NewGraph()
g.SetDir(true)
```


### 頂点(Vertex)の追加/削除

```go
g.AddVertex(0) // vertex 0を追加
g.AddVertex(1) // vertex 1を追加
g.AddVertex(2) // vertex 2を追加

g.RemoveVertex(2) // vertex 2を削除
g.RemoveVertex(1) // vertex 1を削除
```

### 辺(Edge)の追加/削除

`AddEdge`の第3引数はedgeの重み(`int`)になります。

```go
// edgeを追加する前にvertexの追加が必要
g.AddVertex(0)
g.AddVertex(1)

g.AddEdge(0, 1, 3) // 0→1に重み3のedgeを追加

g.RemoveEdge(0, 1) // 0→1のedgeを削除
```

### 可視化

```go
g.Visualize()
```

## Example

簡単な有向グラフを作成し、可視化してみます。

```go
package main

import (
	"github.com/cipepser/goGraphAlgo/graph"
)

func main() {

	g := graph.NewGraph()
	g.SetDir(true)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 2, 0)
	g.AddEdge(0, 3, 0)
	g.AddEdge(2, 3, 0)

	if err := g.Visualize(); err != nil {
		panic(err)
	}

}
```

上記を`main.go`として実行(`go run main.go`)すると`gv.dot`ファイルが生成されます。
これを以下のようにgraphgizでpngにします。

```
$ go run main.go
$ dot -T png ./gv.dot -o ./gv.png
$ open ./gv.png
```

![結果](https://github.com/cipepser/goGraphAlgo/blob/master/graph/example/example.png)

## References
* [Graphviz - Graph Visualization Software](https://www.graphviz.org/)
* [gographviz](https://github.com/awalterschulze/gographviz)
* [graph](https://github.com/cipepser/goGraphAlgo/tree/master/graph)
* [graph-go](https://github.com/arnauddri/algorithms/blob/master/data-structures/graph/graph.go)
## License
MIT