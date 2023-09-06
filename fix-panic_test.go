package pcopy

import (
	"database/sql"
	"testing"
	"time"
)

type TopicItem struct {
	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	OwnerRid   uint64 `protobuf:"varint,4,opt,name=OwnerRid,proto3" json:"OwnerRid,omitempty"`
}

type OmGroupUserTopic struct {
	ID              int64
	Rid             int64
	GroupID         string
	TopicID         string
	Name            sql.NullString
	CreateTime      time.Time
	CreateTimeMilli int64
	OwnerRid        int64
}

func Test_Panic(t *testing.T) {
	// 类型不匹配
	t.Run("类型不一样导致的panic, dst是uint, src是int", func(t *testing.T) {
		v := &TopicItem{}
		Copy(v, &OmGroupUserTopic{
			OwnerRid: 1,
		})
		if v.OwnerRid != 1 {
			t.Fatal("OwnerRid not equal")
		}

		Preheat(&TopicItem{}, &OmGroupUserTopic{})
		dst := &TopicItem{}
		src := &OmGroupUserTopic{OwnerRid: 1}
		Copy(&dst, src, WithUsePreheat())
		if dst.OwnerRid != 1 {
			t.Fatal("OwnerRid not equal")
		}
	})

	t.Run("类型不一样导致的panic, dst是int, src是uint", func(t *testing.T) {
		src := &TopicItem{
			OwnerRid: 1,
		}
		dst := &OmGroupUserTopic{}
		Copy(dst, src)
		if dst.OwnerRid != 1 {
			t.Fatal("OwnerRid not equal")
		}

		Preheat(&OmGroupUserTopic{}, &TopicItem{})
		dst = &OmGroupUserTopic{}
		src = &TopicItem{OwnerRid: 1}
		Copy(dst, src, WithUsePreheat())
		if dst.OwnerRid != 1 {
			t.Fatal("OwnerRid not equal")
		}
	})
}
