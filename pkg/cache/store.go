package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	// Cache_Key_STORE_INFO 店铺信息缓存
	Cache_Key_STORE_INFO      Cache_Key = "store:"
	Cache_Key_STORE_USER_INFO Cache_Key = "storeUser:"
)

type CacheStore struct {
	ctx       context.Context
	cacheConn *redis.Client
}

// NewCacheStore
// @Desc：初始化 store 结构
// @param：ctx
// @param：redisConn
// @return：CacheStoreApi
func NewCacheStore(ctx context.Context, redisConn *redis.Client) CacheStoreApi {
	return &CacheStore{
		ctx:       ctx,
		cacheConn: redisConn,
	}
}

// GetInfo
// @Desc：获取店铺信息&店长信息
// @param：storeId
// @return：map[string]string
// @return：error
func (store *CacheStore) GetInfo(storeId int64) (map[string]string, error) {
	info, err := store.cacheConn.HGetAll(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId)).Result()
	if err == redis.Nil {
		return info, nil
	} else if err != nil {
		return nil, err
	} else if len(info) == 0 {
		return map[string]string{}, nil
	}
	suInfo, err := store.cacheConn.HGetAll(store.ctx, fmt.Sprintf("%s%s", Cache_Key_STORE_USER_INFO, info["storeUserId"])).Result()
	if err == redis.Nil {
		return info, nil
	} else if err != nil {
		return nil, err
	} else if len(suInfo) == 0 {
		suInfo["userId"] = ""
	}
	info["userId"] = suInfo["userId"]
	return info, nil
}

// SetInfo
// @Desc：存储店铺信息缓存
// @param：storeId
// @param：info
// @return：error
func (store *CacheStore) SetInfo(storeId int64, info map[string]interface{}, seconds int64) error {
	_, err := store.cacheConn.Pipelined(store.ctx, func(pipe redis.Pipeliner) error {
		pipe.HMSet(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId), info)
		pipe.Expire(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId), time.Duration(seconds)*time.Second)
		return nil
	})
	return err
}

// SetStoreAndStoreUser
// @Desc： 存储店铺信息和店主信息缓存
// @param：storeId
// @param：storeInfo
// @param：storeUserId
// @param：storeUserInfo
// @param：seconds
// @return：error
func (store *CacheStore) SetStoreAndStoreUser(storeId int64, storeInfo map[string]interface{}, storeUserId int64, storeUserInfo map[string]interface{}, seconds int64) error {
	_, err := store.cacheConn.TxPipelined(store.ctx, func(pipe redis.Pipeliner) error {
		pipe.HMSet(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId), storeInfo)
		pipe.Expire(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId), time.Duration(seconds)*time.Second)

		pipe.HMSet(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_USER_INFO, storeUserId), storeUserInfo)
		pipe.Expire(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_USER_INFO, storeUserId), time.Duration(seconds)*time.Second)
		return nil
	})
	return err
}
