package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	// Cache_Key_STORE_INFO 店铺信息缓存
	Cache_Key_STORE_INFO Cache_Key = "store:"
)

type CacheStore struct {
	ctx       context.Context
	cacheConn *redis.Client
}

// NewCacheStore
// @Desc：初始化 store 结构
// @param：ctx
// @param：redisConn
// @return：CacheApi
func NewCacheStore(ctx context.Context, redisConn *redis.Client) CacheApi {
	return &CacheStore{
		ctx:       ctx,
		cacheConn: redisConn,
	}
}

// GetInfo
// @Desc：获取店铺信息
// @param：storeId
// @return：map[string]string
// @return：error
func (store *CacheStore) GetInfo(storeId int64) (map[string]string, error) {
	info, err := store.cacheConn.HGetAll(store.ctx, fmt.Sprintf("%s%d", Cache_Key_STORE_INFO, storeId)).Result()
	if err == redis.Nil {
		return info, nil
	} else {
		return info, err
	}
}

// SetInfo
// @Desc：存储店铺信息缓存
// @param：storeId
// @param：info
// @return：error
func (store *CacheStore) SetInfo(storeId int64, info map[string]interface{}, seconds int64) error {
	_, err := store.cacheConn.Pipelined(store.ctx, func(pipe redis.Pipeliner) error {
		pipe.HMSet(store.ctx, fmt.Sprintf("%s%d", Cache_Key_USER_INFO, storeId), info)
		pipe.Expire(store.ctx, fmt.Sprintf("%s%d", Cache_Key_USER_INFO, storeId), time.Duration(seconds)*time.Second)
		return nil
	})
	return err
}
