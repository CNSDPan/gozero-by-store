package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	// Cache_Key_USER_INFO 用户信息缓存
	Cache_Key_USER_INFO Cache_Key = "user:"
)

type CacheUser struct {
	ctx       context.Context
	cacheConn *redis.Client
}

// NewCacheUser
// @Desc：初始化 user 结构
// @param：ctx
// @param：redisConn
// @return：*CacheUserApi
func NewCacheUser(ctx context.Context, redisConn *redis.Client) CacheUserApi {
	return &CacheUser{
		ctx:       ctx,
		cacheConn: redisConn,
	}
}

// GetInfo
// @Desc：获取用户信息
// @param：userId
// @return：map[string]string
// @return：error
func (user *CacheUser) GetInfo(userId int64) (map[string]string, error) {
	info, err := user.cacheConn.HGetAll(user.ctx, fmt.Sprintf("%s%d", Cache_Key_USER_INFO, userId)).Result()
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
// @param：seconds
// @return：error
func (user *CacheUser) SetInfo(userId int64, info map[string]interface{}, seconds int64) error {
	_, err := user.cacheConn.Pipelined(user.ctx, func(pipe redis.Pipeliner) error {
		pipe.HMSet(user.ctx, fmt.Sprintf("%s%d", Cache_Key_USER_INFO, userId), info)
		pipe.Expire(user.ctx, fmt.Sprintf("%s%d", Cache_Key_USER_INFO, userId), time.Duration(seconds)*time.Second)
		return nil
	})
	return err
}
