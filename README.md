## 作用
[![Go](https://github.com/antlabs/pcopy/workflows/Go/badge.svg)](https://github.com/antlabs/pcopy/actions)
[![codecov](https://codecov.io/gh/antlabs/pcopy/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/pcopy)

`pcopy.Copy`主要用于两个类型间的深度拷贝, 前身是[deepcopy](https://github.com/antlabs/deepcopy)

新加预热函数。Copy时打开加速开关，达到性能提升4-10倍的效果。

警告: 

高性能的同时可能会有些bug, 如果发现bug可以去掉`pcopy.WithUsePreheat()`试下， 结果不一致，可以提issue。

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
go get github.com/antlabs/pcopy
```

## Quick start
```go
package main

import (
    "fmt"
    "github.com/antlabs/pcopy"
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
   pcopy.Preheat(&dst{}, &src{}) // 一对类型只要预热一次
   pcopy.Copy(&d, &s, pcopy.WithUsePreheat())
   fmt.Printf("%#v\n", d)
   
}

```

## copy slice
```go
package main

import (
        "fmt"

        "github.com/antlabs/pcopy"
)

func main() {
        i := []int{1, 2, 3, 4, 5, 6}
        var o []int

        pcopy.Preheat(&o, &i)
        pcopy.Copy(&o, &i, pcopy.WithUsePreheat())

        fmt.Printf("%#v\n", o)
}

```

## copy map
```go
package main

import (
        "fmt"

        "github.com/antlabs/pcopy"
)

func main() {
        i := map[string]int{
                "cat":  100,
                "head": 10,
                "tr":   3,
                "tail": 44,
        }

        var o map[string]int
        pcopy.Preheat(&o, &i)
        pcopy.Copy(&o, &i, pcopy.WithUsePreheat())

        fmt.Printf("%#v\n", o)
}

```
## simplify business code development
经常看到，对同一个结构体的，有值更新操作，都是一堆手工if 然后赋值的代码。不仅容易出错，还累。快使用pcopy解放双手。
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

        pcopy.Preheat(&a, &b) //只要预热一次
        //可以约化成
        pcopy.Copy(&a, &b, pcopy.WithUsePreheat())
}
```
# benchmark
从零实现的pcopy相比json序列化与反序列化方式拥有更好的性能

[压测仓库位置](https://github.com/1whour/deepcopy-benchmark)
```
goos: darwin
goarch: arm64
pkg: benchmark
Benchmark_Use_reflectValue_MiniCopy-8   	  334728	      3575 ns/op
Benchmark_Use_reflectValue_DeepCopy-8   	  595302	      1956 ns/op
Benchmark_Use_reflectValue_Copier-8     	  203574	      5860 ns/op
Benchmark_Use_Ptr_jsoniter-8            	  821113	      1477 ns/op
Benchmark_Use_Ptr_pcopy-8               	 3390382	       354.0 ns/op
Benchmark_Use_Ptr_coven-8               	 1414197	       848.7 ns/op
PASS
ok  	benchmark	9.771s
```

### 本项目压测
从下面的压测数据可以看到，基本提供了4-10倍的性能提升
```
goos: darwin
goarch: arm64
pkg: github.com/antlabs/pcopy
Benchmark_BaseMap_Unsafe_Pcopy-8               	  529747	      2343 ns/op
Benchmark_BaseMap_miniCopy-8                   	   62181	     19212 ns/op
Benchmark_BaseMap_Reflect-8                    	   93810	     12756 ns/op
Benchmark_BaseSlice_Unsafe_Pcopy-8             	 2013764	       595.1 ns/op
Benchmark_BaseSlice_miniCopy-8                 	  154918	      7728 ns/op
Benchmark_BaseSlice_Reflect-8                  	  188720	      6393 ns/op
Benchmark_BaseType_Unsafe_Pcopy-8              	 4872112	       243.8 ns/op
Benchmark_BaseType_MiniCopy-8                  	  517814	      2278 ns/op
Benchmark_BaseType_Pcopy-8                     	  635156	      1886 ns/op
Benchmark_CompositeMap_Unsafe_Pcopy-8          	  486253	      2409 ns/op
Benchmark_CompositeMap_miniCopy-8              	  229674	      5173 ns/op
Benchmark_CompositeMap_Reflect-8               	  475243	      2490 ns/op
Benchmark_GetLikeFavorited_Unsafe_Pcopy2-8     	  446907	      2662 ns/op
Benchmark_GetLikeFavorited_Unsafe_Pcopy-8      	  470217	      2572 ns/op
Benchmark_GetLikeFavorited_MiniCopy-8          	   85674	     13989 ns/op
Benchmark_GetLikeFavorited_Reflect_Pcopy-8     	  121603	      9856 ns/op
Benchmark_GetRedPoint_Unsafe_Pcopy-8           	 1626688	       736.1 ns/op
Benchmark_GetRedPoint_MiniCopy-8               	  650004	      1871 ns/op
Benchmark_GetRedPoint_Reflect_Pcopy-8          	 1669778	       722.0 ns/op
Benchmark_Interface_Unsafe_Pcopy-8             	 2869022	       421.3 ns/op
Benchmark_Interface_MiniCopy-8                 	  413936	      2704 ns/op
Benchmark_Interface_Pcopy-8                    	  440250	      2688 ns/op
Benchmark_Interface_BaseSlice_Unsafe_Pcopy-8   	 1266501	       947.4 ns/op
Benchmark_Interface_BaseSlice_MiniCopy-8       	  141610	      8422 ns/op
Benchmark_Interface_BaseSlice_Pcopy-8          	  203906	      5917 ns/op
Benchmark_Ptr_BaseType1_Unsafe_Pcopy-8         	  910153	      1310 ns/op
Benchmark_Ptr_BaseType1_Reflect_Pcopy-8        	  391117	      3026 ns/op
Benchmark_Ptr_BaseSlice_Unsafe_Pcopy-8         	  698156	      1704 ns/op
Benchmark_Ptr_BaseSlice_Reflect_Pcopy-8        	  219999	      5415 ns/op
Benchmark_SliceWithStruct_Unsafe_Pcopy-8       	 1395982	       860.3 ns/op
Benchmark_SliceWithStruct_miniCopy-8           	  163154	      7298 ns/op
Benchmark_SliceWithStruct_Reflect_Pcopy-8      	  190728	      6213 ns/op
```
