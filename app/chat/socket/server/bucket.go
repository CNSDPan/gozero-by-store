package server

import (
	"store/pkg/types"
	"sync"
)

type Bucket struct {
	Clock      sync.RWMutex
	UserClient map[int64]*Client
	StoresMap  map[int64]*Store
	Routines   chan types.SocketMsg
	Idx        uint32
}

type Store struct {
	StoreId   int64
	StoreName string
	UserIds   []int64
}

// NewBucket
// @Desc：初始化池子
// @param：cpu 设置几个池子
// @return：[]*Bucket
func NewBucket(cpu uint) []*Bucket {
	buckets := make([]*Bucket, cpu)
	for i := uint(0); i < cpu; i++ {
		buckets[i] = &Bucket{
			Clock: sync.RWMutex{},
			Idx:   uint32(i),
		}
	}
	return buckets
}

// AddBucket
// @Desc：将客户端加入连接池
// @param：userId
// @param：storeIds
func (b *Bucket) AddBucket(userId int64, storeIds ...int64) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	for _, storeId := range storeIds {
		if _, ok := b.StoresMap[storeId]; !ok {
			store := &Store{
				StoreId:   storeId,
				StoreName: "",
				UserIds:   make([]int64, 0),
			}
			store.UserIds = append(store.UserIds, userId)
			b.StoresMap[storeId] = store
		} else {
			b.StoresMap[storeId].UserIds = append(b.StoresMap[storeId].UserIds, userId)
		}
	}
}

// UnBucket
// @Desc：连接池移除客户端
// @param：client
func (b *Bucket) UnBucket(client *Client) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	for _, storeId := range client.StoreIds {
		if store, ok := b.StoresMap[storeId]; ok {
			var newStoresMap = b.StoresMap[storeId].UserIds[:0]
			for _, userId := range store.UserIds {
				if userId == client.UserId {
					continue
				}
				newStoresMap = append(newStoresMap, userId)
			}
			b.StoresMap[storeId].UserIds = newStoresMap
		}
	}
}
