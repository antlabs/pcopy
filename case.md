测试点
### 1. 名字一样，类型不一样的数据类型。
1.1 slice 和 map
1.2 map 和slice
1.3 interface{} 和map
1.4 map和interface{}
1.5 array 和 slice
1.6 

### 2. src是指针类型的各种排列组合
2.1 基础类型
2.1.1 dst是指针, src基础类型(done)
2.1.2 dst是指针，src是指针(done)
2.1.3 dst是基础类型，src是指针(done)

2.2 slice类型
2.1.1 dst是指针, src基础slice类型
2.1.2 dst是指针，src是指针
2.1.3 dst是基础slice类型，src是指针

2.3 map类型
2.1.1 dst是指针, src基础map类型
2.1.2 dst是指针，src是指针
2.1.3 dst是基础map类型，src是指针

2.4 struct类型
2.1.1 dst是指针, src struct类型
2.1.2 dst是指针，src是指针
2.1.3 dst是struct类型，src是指针

### 3. 特殊类型
map[interface{}]interface{}
time.Time
[]byte
