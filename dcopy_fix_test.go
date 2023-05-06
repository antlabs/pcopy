// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// src
type TopicItemSrc struct {
	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

// src
type AllTopicResp struct {
	TopicItem []*TopicItemSrc `protobuf:"bytes,1,rep,name=TopicItem,proto3" json:"TopicItem,omitempty"`
	Total     int32           `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"` // 总条数
}

// desc
type AllTopicResp_AllTopicRespData struct {
	TopicItem []*TopicItemDst `protobuf:"bytes,1,rep,name=TopicItem,proto3" json:"TopicItem,omitempty"`
	Total     int32           `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"` // 总条数
}

type TopicItemDst struct {
	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

type AllTopicResp2 struct {
	Code    int32                          `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string                         `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *AllTopicResp_AllTopicRespData `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func TestCopy(t *testing.T) {
	resp := AllTopicResp{Total: 100, TopicItem: []*TopicItemSrc{
		{ID: "111", Name: "111"},
		{ID: "111", Name: "222"},
	}}

	rsp := AllTopicResp2{Data: &AllTopicResp_AllTopicRespData{}}
	Copy(rsp.Data, resp)

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", rsp.Data)
	assert.NotNil(t, rsp.Data)
	assert.Equal(t, rsp.Data.Total, resp.Total)
	assert.Equal(t, len(rsp.Data.TopicItem), len(resp.TopicItem))
	assert.Greater(t, len(rsp.Data.TopicItem), 0)
	assert.Greater(t, len(resp.TopicItem), 0)
	for i, v := range resp.TopicItem {
		assert.Equal(t, v.ID, rsp.Data.TopicItem[i].ID)
		assert.Equal(t, v.Name, rsp.Data.TopicItem[i].Name)
	}
}
