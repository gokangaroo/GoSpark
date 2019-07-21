package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() error {
	beego.Info("redis config...")
	addr := fmt.Sprintf("%s:%s", beego.AppConfig.String("redis_host"), beego.AppConfig.String("redis_port"))
	password := beego.AppConfig.String("redis_password")
	Redis = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:     password,
		DB:           0,
		MinIdleConns: 10,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		return err
	}
	Logger.Info("Initialize redis successfully!", pong)

	return nil
}
