package src

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var initializeRedis = false
var redisPool *redis.Pool

const redisMaxIdle = 20
const redisIdleTimeout = 120 * time.Second
const redisMaxActive = 100

func InitRedis() {
	if !initializeRedis {
		redisConnection := fmt.Sprintf("%[1]s:%[2]s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

		redisPool = &redis.Pool{
			MaxIdle:     redisMaxIdle,
			IdleTimeout: redisIdleTimeout,
			MaxActive:   redisMaxActive,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", redisConnection, redis.DialPassword(os.Getenv("REDIS_PASSWORD")), redis.DialDatabase(redisDB))

				if err != nil {
					log.Fatal("Failed connect to redis")
					os.Exit(1)
				}

				return conn, err
			},
		}

		conn := redisPool.Get()
		conn.Close()

		initializeRedis = true
	}
}

func GetRedis() *redis.Pool {
	return redisPool
}

func RedisSet(key string, val string) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}

	return nil
}

func RedisGet(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return s, nil
}

func RedisSetExpire(key string, ttl int) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, ttl)

	if err != nil {
		return err
	}

	return nil
}

func RedisDelete(key string) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)

	if err != nil {
		return err
	}

	return nil
}
