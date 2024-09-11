package cache

type Cache_Key string

// CacheUserApi
// @Description: 该接口只是实现了获取缓存数据，没有其他逻辑
type CacheUserApi interface {
	GetInfo(iid int64) (map[string]string, error)
	SetInfo(iid int64, info map[string]interface{}, seconds int64) error
}

// CacheStoreApi
// @Description: 该接口只是实现了获取缓存数据，没有其他逻辑
type CacheStoreApi interface {
	GetInfo(iid int64) (map[string]string, error)
	SetInfo(iid int64, info map[string]interface{}, seconds int64) error
	SetStoreAndStoreUser(storeId int64, storeInfo map[string]interface{}, storeUserId int64, storeUserInfo map[string]interface{}, seconds int64) error
}

type CacheItem struct {
	User  CacheUserApi
	Store CacheStoreApi
}

// NewCache
// @Desc：初始化接口
// @param：user
// @param：store
// @return：*CacheItem
func NewCache(user CacheUserApi, store CacheStoreApi) *CacheItem {
	return &CacheItem{
		User:  user,
		Store: store,
	}
}
