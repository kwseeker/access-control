package runtime

import (
	"casbin-gin/common/component/storage"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"sync"
)

var ApplicationContext Runtime = NewApp()

type Application struct {
	dbs     map[string]*gorm.DB
	casbins map[string]*casbin.SyncedEnforcer
	//engine      http.Handler
	//crontab     map[string]*cron.Cron
	mux sync.RWMutex
	//middlewares map[string]interface{}
	cache storage.AdapterCache
	//queue       storage.AdapterQueue
	//locker      storage.AdapterLocker
	//memoryQueue storage.AdapterQueue
	//handler     map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	//routers     []Router
	//configs     map[string]interface{} // 系统参数
	//appRouters  []func()               // app路由
}

// SetDb 设置对应key的db
func (e *Application) SetDb(key string, db *gorm.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.dbs[key] = db
}

// GetDb 获取所有map里的db数据
func (e *Application) GetDb() map[string]*gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.dbs
}

// GetDbByKey 根据key获取db
func (e *Application) GetDbByKey(key string) *gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	if db, ok := e.dbs["*"]; ok {
		return db
	}
	return e.dbs[key]
}

// SetCacheAdapter 设置缓存
func (e *Application) SetCacheAdapter(c storage.AdapterCache) {
	e.cache = c
}

// GetCacheAdapter 获取缓存
func (e *Application) GetCacheAdapter() storage.AdapterCache {
	return storage.NewCache("", e.cache, "")
}

// GetCachePrefix 获取带租户标记的cache
func (e *Application) GetCachePrefix(key string) storage.AdapterCache {
	return storage.NewCache(key, e.cache, "")
}

func (e *Application) SetCasbin(key string, enforcer *casbin.SyncedEnforcer) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.casbins[key] = enforcer
}

func (e *Application) GetCasbin() map[string]*casbin.SyncedEnforcer {
	return e.casbins
}

// GetCasbinKey 根据key获取casbin
func (e *Application) GetCasbinKey(key string) *casbin.SyncedEnforcer {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.casbins["*"]; ok {
		return e
	}
	return e.casbins[key]
}

// NewApp 默认值
func NewApp() *Application {
	return &Application{
		dbs:     make(map[string]*gorm.DB),
		casbins: make(map[string]*casbin.SyncedEnforcer),
		//crontab:     make(map[string]*cron.Cron),
		//middlewares: make(map[string]interface{}),
		//memoryQueue: queue.NewMemory(10000),
		//handler:     make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
		//routers:     make([]Router, 0),
		//configs:     make(map[string]interface{}),
	}
}
