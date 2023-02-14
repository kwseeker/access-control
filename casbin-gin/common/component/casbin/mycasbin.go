package casbin

import (
	"log"
	"sync"

	"casbin-gin/cmd/config"
	"github.com/casbin/casbin/v2"
	redisWatcher "github.com/casbin/redis-watcher/v2"
	xormAdapter "github.com/casbin/xorm-adapter/v2"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

const (
	modelConf = "conf/model.conf"
)

var (
	Enforcer *casbin.SyncedEnforcer
	once     sync.Once
)

func Setup() *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, err := xormAdapter.NewAdapter(config.DatabaseConfig.Driver, config.DatabaseConfig.Source, true)
		//Enforcer, err = casbin.NewEnforcer(modelConf, adapter)
		Enforcer, err = casbin.NewSyncedEnforcer(modelConf, adapter)
		if err != nil {
			log.Fatal("Load casbin failed, err:", err)
		}
		err = Enforcer.LoadPolicy()
		if err != nil {
			log.Fatal("Load casbin policy failed, err:", err)
		}

		//用于多个Casbin实例之间保持一致， TODO 原理？
		if config.CacheConfig.Redis != nil {
			w, err := redisWatcher.NewWatcher(config.CacheConfig.Redis.Addr, redisWatcher.WatcherOptions{
				Options: redis.Options{
					Network:  "tcp",
					Password: config.CacheConfig.Redis.Password,
				},
				Channel:    "/casbin",
				IgnoreSelf: false,
			})
			if err != nil {
				panic(err)
			}

			err = w.SetUpdateCallback(updateCallback)
			if err != nil {
				panic(err)
			}
			err = Enforcer.SetWatcher(w)
			if err != nil {
				panic(err)
			}
		}

		//log.SetLogger(&Logger{})
		//Enforcer.EnableLog(true)
	})

	log.Println("Casbin enforcer setup")
	return Enforcer
}

func updateCallback(msg string) {
	log.Printf("casbin updateCallback msg: %v\n", msg)
	err := Enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("casbin LoadPolicy err: %v\n", err)
	}
}
