package common

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

////////////////////redis cache service///////////////////////////

type CacheService struct {
	con *redis.Client
}

func (s *CacheService) Init() {
	if s.con == nil {
		s.conn()
	}
}

func (s *CacheService) conn() {
	db, err := strconv.Atoi(Cfg.RedisDatabase)
	if err != nil {
		panic("get redis db form config err:" + err.Error())
	}
	c := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         Cfg.RedisHost,
		Password:     Cfg.RedisPass,
		DB:           db,
		DialTimeout:  60 * time.Second,
		PoolSize:     1000,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	})
	_, err = c.Ping().Result()
	if err != nil {
		panic("init redis err:" + err.Error())
	}
	s.con = c
}

// 获取缓存信息
func (s *CacheService) Get(key string) (string, error) {
	v, err := s.con.Get(key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (s *CacheService) GetInt(key string) (int64, error) {
	v, err := s.con.Get(key).Int64()
	if err != nil {
		return int64(0), err
	}
	return v, nil
}

// 设置缓存信息
func (s *CacheService) Save(key string, val interface{}, expire int) bool {
	dur := time.Second * time.Duration(expire)
	err := s.con.Set(key, val, dur).Err()
	if err != nil {
		return false
	}
	return true
}

// 删除缓存信息
func (s *CacheService) Del(key string) bool {
	err := s.con.Del(key).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取缓存（哈希）
func (s *CacheService) HGet(key, index string) (string, error) {
	v, err := s.con.HGet(key, index).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

// 设置缓存（哈希）
func (s *CacheService) HSet(key, index, data string) bool {
	err := s.con.HSet(key, index, data).Err()
	if err != nil {
		return false
	}
	return true
}

// 设置哈希缓存（带过期时间）
func (s *CacheService) HSetEx(key, index, data string, expire int) bool {
	err := s.con.HSet(key, index, data).Err()
	if err != nil {
		return false
	}
	// 设置过期时间
	dur := time.Second * time.Duration(expire)
	s.con.Expire(key, dur)
	return true
}

// 删除缓存（哈希）
func (s *CacheService) HDel(key, index string) bool {
	err := s.con.HDel(key, index).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取 哈希表中所有域（field）列表
func (s *CacheService) HKeys(key string) ([]string, error) {
	v, err := s.con.HKeys(key).Result()
	if err != nil {
		return []string{}, err
	}
	return v, nil
}
