// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"
)

func Benchmark_GetLikeFavorited_Unsafe_dcopy2(b *testing.B) {
	err := Preheat(&GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{}, &GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		err = Preheat(&GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{}, &GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{})
		if err != nil {
			b.Fatal(err)
		}
		var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
		err = Copy(&dst, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_GetLikeFavorited_Unsafe_dcopy(b *testing.B) {
	err := Preheat(&GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{}, &GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData{})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
		err = Copy(&dst, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_GetLikeFavorited_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
		dst.Items = make([]*GetMessageLikeFavoritedRespItem, len(local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items))
		for i := range local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items {
			dst.Items[i] = &GetMessageLikeFavoritedRespItem{
				LikeType:         local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].LikeType,
				CommentType:      local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].CommentType,
				ResourceType:     local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].ResourceType,
				PostID:           local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].PostID,
				RefMessage:       local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].RefMessage,
				RefMessageID:     local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].RefMessageID,
				ParentCid:        local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].ParentCid,
				CommentModule:    local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].CommentModule,
				ArticleThumbnail: local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].ArticleThumbnail,
				Info: &MessageWhoAndTime{
					HeadPic:  local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].Info.HeadPic,
					Nickname: local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].Info.Nickname,
					DID:      local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].Info.DID,
					Seq:      local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].Info.Seq,
					Time:     local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData.Items[i].Info.Time,
				},
			}
		}
	}
}

func Benchmark_GetLikeFavorited_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
		err := miniCopy(&dst, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData)
		if err != nil {
			b.Fatal(err)
		}
		// miniCopy(&dst, &td)
	}
}

func Benchmark_GetLikeFavorited_dcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData
		err := Copy(&dst, &local_GetMessageLikeFavoritedResp_GetMessageLikeFavoritedRespData)
		if err != nil {
			b.Fatal(err)
		}

	}
}
