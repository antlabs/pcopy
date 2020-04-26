## 作用
deepcopy.Copy主要用作结构之间的深度拷贝[从零实现]


## feature

## 内容
- [Installation](#Installation)
- [Quick start](#quick-start)
- [example](#example)
    - [1. 控制拷贝结构体最多深度](#max-copy-depth)
    - [2. 只拷贝设置tag的结构体成员](#copy-only-the-specified-tag)
    - [3. 只拷贝指定类型数据](#copy-only-the-specified-type)

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
```go
    deepcopy.Copy().MaxDepth(1).Do()
```

## copy only the specified   tag
只拷贝结构体里面有copy tag的字段
```go
    deepcopy.Copy().RegisterTagName("copy").Do()
```

## copy only the specified type
只拷贝指定类型
```go
deepcopy.Copy().OnlyType(reflect.String, reflect.Int).Do()
```