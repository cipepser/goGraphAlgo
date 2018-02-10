# set

[![Build Status](https://travis-ci.org/cipepser/goGraphAlgo.svg?branch=master)](https://travis-ci.org/cipepser/goGraphAlgo)
[![Coverage Status](https://coveralls.io/repos/github/cipepser/goGraphAlgo/badge.svg?branch=master)](https://coveralls.io/github/cipepser/goGraphAlgo?branch=master)

集合に対する操作を行うライブラリ

## How to Install

```
$ go get github.com/cipepser/goGraphAlgo/...
```

## How to Use

集合に対する操作として以下をサポートしています。

* 集合型を作る(現時点で`IntSet`のみ)
* 要素の追加/削除
* 要素の個数
* 要素が含まれるか
* 集合同士の比較
* 差集合を返す
* 共通部分を返す
* 和集合を返す

### 集合型を作る(`IntSet`)

```go
s := NewIntSet()
```

### 要素の追加/削除

```go
s.Add(0)
s.Remove(0)
```

### 要素の個数

```go
s.Cardinality()
```

### 要素が含まれるか

```go
s.Contains(0)
```

### 集合同士の比較

```go
s.Equal(other)
```

### 差集合を返す

```go
d := s.Difference(other)
```

### 共通部分を返す

```go
i := s.Intersect(other)
```

### 和集合を返す

```go
u := s.Union(other)
```

## References
* [golang-set](https://github.com/deckarep/golang-set/blob/master/set.go)

## License
MIT