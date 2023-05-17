## 作用
[![Go](https://github.com/antlabs/dcopy/workflows/Go/badge.svg)](https://github.com/antlabs/dcopy/actions)
[![codecov](https://codecov.io/gh/antlabs/dcopy/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/dcopy)

`dcopy.Copy`主要用于两个类型间的深度拷贝, 相比上个版本deepcopy.Copy，主要是性能方面的优化
dcopy.Copy的前身是deepcopy.Copy优化版本。新加预热函数。Copy时新加使用预热接口函数，达到性能提升4-10倍的效果。
警告: 高性能的同时可能会有些bug, 如果发现bug可以去掉`dcopy.WithUsePreheat()`试下， 结果不一致，可以提issue。

## feature
* 高性能, 相对第一个版本提升4-10倍的性能
* 支持异构结构体拷贝, dst和src可以是不同的类型，会拷贝dst和src交集的部分
* 多类型支持struct/map/slice/array/int...int64/uint...uint64/ 等等

## 内容
- [Installation](#Installation)
- [Quick start](#quick-start)
- [example](#example)
    - [1.拷贝slice](#copy-slice)
    - [2.拷贝map](#copy-map)
    - [3.简化业务代码开发](#simplify-business-code-development)
- [性能压测](#benchmark)
## Installation
```
go get github.com/antlabs/dcopy
```

## Quick start
```go
package main

import (
    "fmt"
    "github.com/antlabs/dcopy"
)

type dst struct {
    ID int
    Result string
}

type src struct{
    ID int
    Text string
}
func main() {
   d, s := dst{}, src{ID:3}
   dcopy.Preheat(&dst{}, &src{}) // 一对类型只要预热一次
   dcopy.Copy(&d, &s, dcopy.WithUsePreheat())
   fmt.Printf("%#v\n", d)
   
}

```

## copy slice
```go
package main

import (
        "fmt"

        "github.com/antlabs/dcopy"
)

func main() {
        i := []int{1, 2, 3, 4, 5, 6}
        var o []int

        dcopy.Preheat(&o, &i)
        dcopy.Copy(&o, &i, dcopy.WithUsePreheat())

        fmt.Printf("%#v\n", o)
}

```

## copy map
```go
package main

import (
        "fmt"

        "github.com/antlabs/dcopy"
)

func main() {
        i := map[string]int{
                "cat":  100,
                "head": 10,
                "tr":   3,
                "tail": 44,
        }

        var o map[string]int
        dcopy.Preheat(&o, &i)
        dcopy.Copy(&o, &i, dcopy.WithUsePreheat())

        fmt.Printf("%#v\n", o)
}

```
## simplify business code development
经常看到，对同一个结构体的，有值更新操作，都是一堆手工if 然后赋值的代码。不仅容易出错，还累。快使用dcopy解放双手。
```go
type option struct {
        Int int
        Float64 float64
        S  string
}

func main() {
        var a, b option
        if b.Int != 0 {
                a.Int = b.Int
        }

        if b.Float64 != 0.0 {
                a.Float64 = b.Float64
        }

        if b.S != "" {
                a.S = b.S
        }

        dcopy.Preheat(&a, &b) //只要预热一次
        //可以约化成
        dcopy.Copy(&a, &b, dcopy.WithUsePreheat())
}
```
# benchmark
从零实现的dcopy相比json序列化与反序列化方式拥有更好的性能
```
goos: linux
goarch: amd64
pkg: github.com/antlabs/dcopy
Benchmark_MiniCopy-16    	  159084	      6737 ns/op
Benchmark_dcopy-16    	  374920	      2895 ns/op
PASS
ok  	github.com/antlabs/dcopy	2.275s

```
