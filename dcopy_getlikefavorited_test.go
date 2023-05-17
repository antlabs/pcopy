// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommentModule int32

const (
	CommentModule_CommentModuleUndefined CommentModule = 0
	CommentModule_CommentModuleDiary     CommentModule = 1
	CommentModule_CommentModuleHotPot    CommentModule = 2
)

type LikeType int32

const (
	LikeType_LIKE    LikeType = 0
	LikeType_COLLECT LikeType = 1
)

type GetMessageLikeFavoritedResp struct {
	Code    int32
	Message string
	Data    *GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
}

type MessageWhoAndTime struct {
	HeadPic  string
	Nickname string
	Time     string
	DID      uint64
	Seq      int64
}

type CommentType int32

const (
	CommentType_DIARY CommentType = 0

	CommentType_COMMENT CommentType = 1
)

type MessageResourceType int32

const (
	MessageResourceType_MessageResourceTypeLike MessageResourceType = 2

	MessageResourceType_MessageResourceTypeFavorited MessageResourceType = 3

	MessageResourceType_MessageesourceTypeText MessageResourceType = 4

	MessageResourceType_MessageResourceTypeImage MessageResourceType = 5

	MessageResourceType_MessageResourceTypeVideo MessageResourceType = 6
)

type GetMessageLikeFavoritedRespItem struct {
	Info *MessageWhoAndTime

	LikeType LikeType

	CommentType  CommentType
	ResourceType MessageResourceType

	PostID uint64

	RefMessage       string
	RefMessageID     uint64
	ParentCid        uint64
	CommentModule    CommentModule
	ArticleThumbnail string
}

type GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData struct {
	Items []*GetMessageLikeFavoritedRespItem
}

var local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData = GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{
	Items: []*GetMessageLikeFavoritedRespItem{
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "headpic",
				Nickname: "nickname",
				Time:     "time",
				DID:      1,
				Seq:      1,
			},
			LikeType:         LikeType_COLLECT,
			CommentType:      CommentType_COMMENT,
			ResourceType:     MessageResourceType_MessageResourceTypeLike,
			PostID:           1,
			RefMessage:       "refmessage",
			RefMessageID:     1,
			ParentCid:        1,
			CommentModule:    CommentModule_CommentModuleDiary,
			ArticleThumbnail: "articlethumbnail",
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "headpic2",
				Nickname: "nickname2",
				Time:     "time2",
				DID:      2,
				Seq:      2,
			},
			LikeType:         LikeType_COLLECT,
			CommentType:      CommentType_COMMENT,
			ResourceType:     MessageResourceType_MessageResourceTypeLike,
			PostID:           2,
			RefMessage:       "refmessage",
			RefMessageID:     2,
			ParentCid:        2,
			CommentModule:    CommentModule_CommentModuleDiary,
			ArticleThumbnail: "articlethumbnail",
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "headpic2",
				Nickname: "nickname2",
				Time:     "time2",
				DID:      2,
				Seq:      2,
			},
			LikeType:         LikeType_COLLECT,
			CommentType:      CommentType_COMMENT,
			ResourceType:     MessageResourceType_MessageResourceTypeLike,
			PostID:           2,
			RefMessage:       "refmessage",
			RefMessageID:     2,
			ParentCid:        2,
			CommentModule:    CommentModule_CommentModuleDiary,
			ArticleThumbnail: "articlethumbnail",
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "headpic2",
				Nickname: "nickname2",
				Time:     "time2",
				DID:      2,
				Seq:      2,
			},
			LikeType:         LikeType_COLLECT,
			CommentType:      CommentType_COMMENT,
			ResourceType:     MessageResourceType_MessageResourceTypeLike,
			PostID:           2,
			RefMessage:       "refmessage",
			RefMessageID:     2,
			ParentCid:        2,
			CommentModule:    CommentModule_CommentModuleDiary,
			ArticleThumbnail: "articlethumbnail",
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "headpic2",
				Nickname: "nickname2",
				Time:     "time2",
				DID:      2,
				Seq:      2,
			},
			LikeType:         LikeType_COLLECT,
			CommentType:      CommentType_COMMENT,
			ResourceType:     MessageResourceType_MessageResourceTypeLike,
			PostID:           2,
			RefMessage:       "refmessage",
			RefMessageID:     2,
			ParentCid:        2,
			CommentModule:    CommentModule_CommentModuleDiary,
			ArticleThumbnail: "articlethumbnail",
		},
	},
}

func Test_GetMessageLikeFavorited_1(t *testing.T) {
	err := Preheat(&GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{}, &GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{})
	assert.NoError(t, err)

	var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
	err = Copy(&dst, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, WithUsePreheat())
	assert.NoError(t, err)
	fmt.Printf("%d\n", len(dst.Items))
	assert.Equal(t, local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, dst)
}

// 左边传参数是个二级指针
func Test_GetMessageLikeFavorited_2(t *testing.T) {
	err := Preheat(&GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{}, &GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{})
	assert.NoError(t, err)
	var resp GetMessageLikeFavoritedResp
	// var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
	err = Copy(&resp.Data, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, *resp.Data)
}
