package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetRedPointResp_GetRedPointRespData struct {
	Point map[uint32]uint32 `protobuf:"bytes,1,rep,name=point,proto3" json:"point,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

type GetRedPointResp struct {
	Code    int32                                `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string                               `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *GetRedPointResp_GetRedPointRespData `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

type GetRedPointRespData struct {
	Point map[uint32]uint32 `protobuf:"bytes,1,rep,name=point,proto3" json:"point,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

var local_GetRedPointRespData = GetRedPointRespData{
	Point: map[uint32]uint32{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	},
}

func Test_GetRedPoint(t *testing.T) {
	Preheat(&GetRedPointResp{}, &GetRedPointRespData{})
	rv := GetRedPointResp{}

	err := Copy(&rv.Data, local_GetRedPointRespData, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, rv.Data.Point, local_GetRedPointRespData.Point)
}
