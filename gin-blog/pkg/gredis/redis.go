package gredis

import (
	"gin-blog/pkg/setting"
	"time"

	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxAative,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
