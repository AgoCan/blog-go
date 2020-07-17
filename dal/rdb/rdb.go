package rdb

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"

	"blog-go/config"
)

// RDB redis 单例
var RDB redis.Conn

// NewPool 实例链接池
func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("PING")
			if err != nil {
				_, err := c.Do("AUTH", password)
				if err != nil {

					fmt.Println("redis auth failed: ", err)
					return nil, err
				}

			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// InitRedis 初始化链接
func InitRedis() (err error) {
	redisPool := NewPool(config.RedisServer, config.RedisPassword)
	RDB, err = redisPool.Dial()
	if err != nil {
		return err
	}
	_, err = RDB.Do("select", config.RedisDb)
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭数据库
func Close() {
	err := RDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
