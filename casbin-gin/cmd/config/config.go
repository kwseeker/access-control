package config

import (
	"casbin-gin/cmd/config/database"
	"casbin-gin/common/component/casbin"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	_cfg *Settings
)

type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

// OnChange 用于热更新
func (e *Settings) OnChange() {
	e.Setup()
	log.Println("!!! config change and reload")
}

func (e *Settings) Setup() {
	e.Settings.setup()
	log.Println("!!! config init")
}

// Config 配置集合
type Config struct {
	//下面的对象都是存储配置信息的
	Application *Application
	Database    *database.Database `yaml:"database"`
	//Databases *map[string]*Database `yaml:"databases"`
	Cache *Cache `yaml:"cache"`
}

// 通过解析获取到的配置数据初始化组件实例
func (c *Config) setup() {
	//数据库
	_cfg.Settings.Database.Setup()
	//Casbin
	casbin.Setup()
	//缓存
	_cfg.Settings.Cache.Setup()
}

func ParseAndSetup(fp string) {
	//1配置解析
	parse(fp)
	//2组件安装
	_cfg.Setup()
}

func parse(fp string) {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Database:    database.Config,
			Cache:       CacheConfig,
		},
	}

	yf, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal("Config setup failed, err:", err)
	}
	err = yaml.Unmarshal(yf, _cfg)
	if err != nil {
		log.Fatal("Config file Unmarshal failed, err:", err)
	}
}
