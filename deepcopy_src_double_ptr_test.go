// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Msg struct {
	FromUid int32  `protobuf:"varint,1,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	ToUid   int32  `protobuf:"varint,2,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Msg     string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

// 测试src是双指针的情况
// 为了提升deepcopy库的兼容性, src是双指针, dst是单指针, 也可以正常拷贝
func Test_Src_DoublePtr_Test(t *testing.T) {
	m1 := Msg{FromUid: 1, ToUid: 2, Msg: "1to2"}

	// 单指针
	m1Ptr := &m1

	var m2 Msg
	// dst(m2) 是单指针
	// src(m1Ptr) 是双指针
	err := Copy(&m2, &m1Ptr).Do()

	assert.NoError(t, err)
	assert.Equal(t, m1, m2)
}

func Test_Src_DoublePtr_Test_CopyEx(t *testing.T) {
	m1 := Msg{FromUid: 1, ToUid: 2, Msg: "1to2"}

	// 单指针
	m1Ptr := &m1

	var m2 Msg
	// dst(m2) 是单指针
	// src(m1Ptr) 是双指针
	err := CopyEx(&m2, &m1Ptr)

	assert.NoError(t, err)
	assert.Equal(t, m1, m2)
}
