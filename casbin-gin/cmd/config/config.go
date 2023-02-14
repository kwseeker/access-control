package config

import (
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

// Config 配置集合
type Config struct {
	Database *Database `yaml:"database"`
	//Databases *map[string]*Database `yaml:"databases"`
	Cache *Cache `yaml:"cache"`
}

func Setup(s string) {
	_cfg = &Settings{
		Settings: Config{
			Database: DatabaseConfig,
			Cache:    CacheConfig,
		},
	}

	yf, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("Config setup failed, err:", err)
	}
	err = yaml.Unmarshal(yf, _cfg)
	if err != nil {
		log.Fatal("Config file Unmarshal failed, err:", err)
	}

	_, err = _cfg.Settings.Cache.Setup()
	if err != nil {
		log.Fatal("Config cache setup failed, err:", err)
	}
}
