package repository

import "time"

type RedisRepository interface {
	GetByKey(key string, data interface{}) error
	SetByKey(key string, data interface{}, expiration time.Duration) error
	FlushAll() error
}
