## 作用
[![Go](https://github.com/antlabs/deepcopy/workflows/Go/badge.svg)](https://github.com/antlabs/deepcopy/actions)
[![codecov](https://codecov.io/gh/antlabs/deepcopy/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/deepcopy)

deepcopy.Copy主要用作结构之间的深度拷贝[从零实现]


## feature

## 内容
- [Installation](#Installation)
- [Quick start](#quick-start)
- [example](#example)
    - [1. 控制拷贝结构体最多深度](#max-copy-depth)
    - [2. 只拷贝设置tag的结构体成员](#copy-only-the-specified-tag)

## Installation
```
go get github.com/antlabs/deepcopy
```

## Quick start
```go
package main

import (
    "fmt"
    "github.com/antlabs/deepcopy"
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
   deepcopy.Copy(&d, &s).Do()
   fmt.Printf("%#v\n", d)
   
}

```

## max copy depth
如果src的结构体嵌套了两套，MaxDepth可以控制只拷贝一层
```go
    deepcopy.Copy(dst, src).MaxDepth(1).Do()
```

## copy only the specified   tag
只拷贝结构体里面有copy tag的字段
```go
    deepcopy.Copy(dst, src).RegisterTagName("copy").Do()
```
