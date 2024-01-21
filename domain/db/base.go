package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"quqo_challenge/infrastructure/config"
	"time"
)

type Persistence struct {
	AppDb   *gorm.DB
	RedisDb *redis.Client
}

func NewPersistence(a config.Configuration) *Persistence {
	appDb, err := newAppDb(a)
	if err != nil {
		panic(err)
	}

	rdb, err := newRedisDb(a)
	if err != nil {
		panic(err)
	}

	return &Persistence{
		AppDb:   appDb,
		RedisDb: rdb,
	}
}

func (p *Persistence) ClosePersistence() {
	appDb, errQ := p.AppDb.DB()
	if errQ != nil {
		panic(errQ)
	}

	errDbClose := appDb.Close()
	if errDbClose != nil {
		panic(errDbClose)
	}

	errRedisClose := p.RedisDb.Close()
	if errRedisClose != nil {
		panic(errRedisClose)
	}
}

func newAppDb(a config.Configuration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=verify-full",
		a.DbConfig.Host,
		a.DbConfig.User,
		a.DbConfig.Password,
		a.DbConfig.DbName,
		a.DbConfig.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func newRedisDb(a config.Configuration) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     a.RedisDbConfig.Address,
		Password: a.RedisDbConfig.Password,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
