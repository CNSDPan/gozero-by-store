package server

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/pkg/consts"
	"store/pkg/types"
	"sync"
)

type Bucket struct {
	Clock      sync.RWMutex
	UserClient map[int64]*Client
	StoresMap  map[int64]*Store
	Routines   chan types.SocketMsgBody
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
			Clock:      sync.RWMutex{},
			UserClient: make(map[int64]*Client, 0),
			StoresMap:  make(map[int64]*Store, 0),
			Routines:   make(chan types.SocketMsgBody, 10000),
			Idx:        uint32(i),
		}
		go func(bucket *Bucket) {
			bucket.RoutineSubscribe()
		}(buckets[i])
	}
	return buckets
}

// AddBucket
// @Desc：将客户端加入连接池
// @param：userId
// @param：storeIds
func (b *Bucket) AddBucket(client *Client, storeIds ...int64) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	for _, storeId := range storeIds {
		if _, ok := b.StoresMap[storeId]; !ok {
			store := &Store{
				StoreId:   storeId,
				StoreName: "",
				UserIds:   make([]int64, 0),
			}
			store.UserIds = append(store.UserIds, client.UserId)
			b.StoresMap[storeId] = store
			b.UserClient[client.UserId] = client
		} else {
			b.UserClient[client.UserId] = client
			b.StoresMap[storeId].UserIds = append(b.StoresMap[storeId].UserIds, client.UserId)
		}
	}
}

// UnBucket
// @Desc：连接池移除客户端
// @param：client
func (b *Bucket) UnBucket(client *Client) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	delete(b.UserClient, client.UserId)
	WsServer.Log.Errorf("连接池StoreIds:%+v", client.StoreIds)
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
			WsServer.Log.Errorf("连接池:%+v", b.StoresMap[storeId])
		}
	}
}

func (b *Bucket) BroadcastMsg(socketMsgBody types.SocketMsgBody) {
	select {
	case b.Routines <- socketMsgBody:
	default:
	}
	return
}

// RoutineSubscribe
// @Desc：每个池子有单独接收订阅者传输过来的数据并处理的协程
func (b *Bucket) RoutineSubscribe() {
	var (
		ok         bool
		dataNormal types.DataByNormal
		store      *Store
		client     *Client
		socketMsg  types.SocketMsg
	)
	for socketMsgBody := range b.Routines {
		// 将data类型map[string]interface{}转成types.DataByNormal
		dByte, err := jsonx.Marshal(socketMsgBody.Event.Data)
		if err != nil {
			continue
		}
		err = jsonx.Unmarshal(dByte, &dataNormal)
		if err != nil {
			continue
		}
		socketMsg = types.SocketMsg{
			Operate:       socketMsgBody.Operate,
			Method:        socketMsgBody.Method,
			StoreId:       dataNormal.StoreId,
			SendUserId:    dataNormal.SendUserId,
			ReceiveUserId: dataNormal.ReceiveUserId,
			Extend:        "",
			Body: types.SocketMsgBody{
				Operate:      socketMsgBody.Operate,
				Method:       socketMsgBody.Method,
				ResponseTime: socketMsgBody.ResponseTime,
				Event: types.Event{
					Params: socketMsgBody.Event.Params,
					Data:   socketMsgBody.Event.Data,
				},
			},
		}
		if socketMsgBody.Operate == consts.OperatePrivate {
			if client, ok = b.UserClient[dataNormal.ReceiveUserId]; ok {
				client.PushMsg(socketMsg)
			}
		} else if socketMsgBody.Operate == consts.OperatePublic {
			if store, ok = b.StoresMap[dataNormal.StoreId]; ok {
				for _, userId := range store.UserIds {
					if client, ok = b.UserClient[userId]; ok {
						client.PushMsg(socketMsg)
					}
				}
			}
		}
	}
}
