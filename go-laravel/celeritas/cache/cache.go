package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]interface{}

func (c *RedisCache) Has(key string) (bool, error) {
	innerKey := fmt.Sprintf("%s:%s", c.Prefix, key)
	conn := c.Conn.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)

	ok, err := redis.Bool(conn.Do("EXISTS", innerKey))
	if err != nil {
		return false, err
	}
	return ok, nil
}
