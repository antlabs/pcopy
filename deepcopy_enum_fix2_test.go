// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DiaryResourceTypeDst int32

const (
	// 文本
	DiaryResourceType_DiaryResourceTypeTextDst DiaryResourceTypeDst = 0
	// 图片
	DiaryResourceType_DiaryResourceTypeImageDst DiaryResourceTypeDst = 1
	// 视频
	DiaryResourceType_DiaryResourceTypeVideoDst DiaryResourceTypeDst = 2
	// Ours类型
	DiaryResourceType_DiaryResourceTypeOursDst DiaryResourceTypeDst = 3
)

type SquareDiaryItemDst struct {
	// 资源类型
	ResourceType DiaryResourceTypeDst `protobuf:"varint,7,opt,name=ResourceType,proto3,enum=topic.v1.DiaryResourceType" json:"ResourceType,omitempty"`
	// 资源链接数组
}

type GetRecommendedListResp_GetRecommendedListRespData struct {
	// 列表
	List []*SquareDiaryItemDst `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	// 是否有更多记录
	HasMore bool `protobuf:"varint,2,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	// 记录位置
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
	// 相同记录的偏移量
	Offset int32 `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

type DiaryResourceTypeSrc int32

const (
	// 文本
	DiaryResourceType_DiaryResourceTypeText DiaryResourceTypeSrc = 0
	// 图片
	DiaryResourceType_DiaryResourceTypeImage DiaryResourceTypeSrc = 1
	// 视频
	DiaryResourceType_DiaryResourceTypeVideo DiaryResourceTypeSrc = 2
	// Ours类型
	DiaryResourceType_DiaryResourceTypeOurs DiaryResourceTypeSrc = 3
)

type SquareDiaryItemSrc struct {
	// 资源类型
	ResourceType DiaryResourceTypeSrc `protobuf:"varint,7,opt,name=ResourceType,proto3,enum=topic.v1.DiaryResourceType" json:"ResourceType,omitempty"`
	// 资源链接数组
}

type GetRecommendedListResp struct {
	// 列表
	List []*SquareDiaryItemSrc `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	// 是否有更多记录
	HasMore bool `protobuf:"varint,2,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	// 记录位置
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
	// 相同记录的偏移量
	Offset int32 `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func Test_Fix2(t *testing.T) {
	var r *GetRecommendedListResp_GetRecommendedListRespData
	resp := GetRecommendedListResp{
		// 帮我优化这个代码
		List:    []*SquareDiaryItemSrc{{ResourceType: DiaryResourceType_DiaryResourceTypeVideo}},
		HasMore: true,
		Pos:     10,
		Offset:  11,
	}
	err := Copy(&r, &resp)
	assert.NoError(t, err)
	assert.NotEqual(t, r, nil)
	assert.NotEqual(t, len(r.List), 0)
}
