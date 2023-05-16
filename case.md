测试点
### 1. 名字一样，类型不一样的数据类型。
1.1 slice 和 map(todo)
1.2 map 和slice(todo)
1.3 interface{} 和map(todo)
1.4 map和interface{}(todo)
1.5 array 和 slice(todo)

### 2. src是指针类型的各种排列组合
2.1 基础类型
2.1.1 dst是指针, src基础类型(done)
2.1.2 dst是指针，src是指针(done)
2.1.3 dst是基础类型，src是指针(done)

2.2 基础类型-2级指针
2.2.1 dst是指针, src基础类型(done)
2.2.2 dst是指针，src是指针(done)
2.2.3 dst是基础类型，src是指针(done)

2.3 slice类型
2.3.1 dst是指针, src基础slice类型(done)
2.3.2 dst是指针，src是指针(done)
2.3.3 dst是基础slice类型，src是指针(done)

2.4 map类型
2.4.1 dst是指针, src基础map类型(done)
2.4.2 dst是指针，src是指针(done)
2.4.3 dst是基础map类型，src是指针(done)

2.5 struct类型
2.5.1 dst是指针, src struct类型(done)
2.5.2 dst是指针，src是指针(done)
2.5.3 dst是struct类型，src是指针(done)

### 3. 特殊类型
1. map[interface{}]interface{}
2. time.Time-结构体，特殊处理了 (done) 
3. []byte 和[]uint8处理一样 (done) 

### 4. slice类型
1. 基础slice
1.1 dst的元素是指针slice src元素不是(done)
1.2 dst的元素不是指针 slice的元素是指针(done)
1.3 dst的元素是指针, src的元素指针(done)

// 1。slice里面套基础map
// type sliceMap struct {
// 	// 1.1 slice里面套基础map，map里面套基础类型
// 	SliceWithMap []map[string]string
// 	// 1.2 slice里面套基础map，map里面套struct
// 	SliceWithMapStruct []map[string]DCopyDst_BaseSlice
// 	// 1.3 slice里面套基础map，map里面套interface
// 	// SliceWithMapInterface []map[string]interface{}
// }

// 2. slice里面套map

// 3. slice里面套slice

// 4. slice里面套struct

// 5. slice里面套interface

// 6. slice里面套基础类型

// 7. slice里面套slice里面套基础类型

// 8. slice里面套slice里面套map

// 9. slice里面套slice里面套struct

// 10. slice里面套slice里面套interface

// 11. slice里面套slice里面套slice
