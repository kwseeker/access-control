package runtime

import (
	"casbin-gin/common/component/storage"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

type Runtime interface {
	// SetDb 多db设置，SetDbs不允许并发,可以根据自己的业务，例如app分库、host分库
	SetDb(key string, db *gorm.DB)
	GetDb() map[string]*gorm.DB
	GetDbByKey(key string) *gorm.DB

	SetCasbin(key string, enforcer *casbin.SyncedEnforcer)
	GetCasbin() map[string]*casbin.SyncedEnforcer
	GetCasbinKey(key string) *casbin.SyncedEnforcer

	//// SetEngine 使用的路由
	//SetEngine(engine http.Handler)
	//GetEngine() http.Handler

	// SetCacheAdapter cache
	SetCacheAdapter(storage.AdapterCache)
	GetCacheAdapter() storage.AdapterCache
	GetCachePrefix(string) storage.AdapterCache
}
