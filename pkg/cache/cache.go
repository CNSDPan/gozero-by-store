package cache

type Cache_Key string

// CacheApi
// @Description: 该接口只是实现了获取缓存数据，没有其他逻辑
type CacheApi interface {
	GetInfo(iid int64) (map[string]string, error)
	SetInfo(iid int64, info map[string]interface{}, seconds int64) error
}

type CacheItem struct {
	User  CacheApi
	Store CacheApi
}

// NewCache
// @Desc：初始化接口
// @param：user
// @param：store
// @return：*CacheItem
func NewCache(user CacheApi, store CacheApi) *CacheItem {
	return &CacheItem{
		User:  user,
		Store: store,
	}
}
