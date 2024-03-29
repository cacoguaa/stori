package config

import (
	"stori/utils"
	"sync"
	"time"
)

type Config struct {
	Filepath string
	DBConfig
	EmailConfig
}

type DBConfig struct {
	Host            string
	UserName        string
	Password        string
	DbName          string
	Port            int
	PoolSize        int32
	MaxConnLifetime time.Duration
}

type EmailConfig struct {
	Host string
	Port string
}

var once sync.Once
var config Config

func Environments() Config {
	once.Do(func() {
		config = Config{
			Filepath: utils.Get("FILEPATH"),
			DBConfig: DBConfig{
				Host:            utils.Get("DB_HOST"),
				UserName:        utils.Get("DB_USER"),
				Password:        utils.Get("DB_PASS"),
				DbName:          utils.Get("DB_NAME"),
				Port:            utils.GetInt("DB_PORT"),
				PoolSize:        int32(utils.GetInt("DB_MAX_POOL_SIZE")),
				MaxConnLifetime: 240 * time.Second,
			},
			EmailConfig: EmailConfig{
				Host: "smtp.gmail.com",
				Port: "587",
			},
		}
	})

	return config
}
