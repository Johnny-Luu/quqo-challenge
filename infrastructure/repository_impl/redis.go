package repository_impl

import (
	"context"
	"encoding/json"
	"quqo_challenge/domain/db"
	"quqo_challenge/domain/repository"
	"time"
)

type RedisRepositoryImpl struct {
	p *db.Persistence
}

func NewRedisRepository(p *db.Persistence) repository.RedisRepository {
	return RedisRepositoryImpl{p: p}
}

func (repo RedisRepositoryImpl) GetByKey(key string, data interface{}) error {
	val, err := repo.p.RedisDb.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return err
	}

	return nil
}

func (repo RedisRepositoryImpl) SetByKey(key string, data interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = repo.p.RedisDb.Set(context.Background(), key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (repo RedisRepositoryImpl) FlushAll() error {
	err := repo.p.RedisDb.FlushAll(context.Background()).Err()
	if err != nil {
		return err
	}

	return nil
}
