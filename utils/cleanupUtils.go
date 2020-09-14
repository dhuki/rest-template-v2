package utils

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Cleanup interface {
	Close() error
}

type Dependencies struct {
	GormDB      *gorm.DB
	Server      *http.Server
	RedisClient *redis.Client
}

func (d Dependencies) Close() error {
	if err := d.Server.Close(); err != nil {
		return err
	}
	if _, err := d.GormDB.DB(); err != nil {
		return err
	}
	if err := d.RedisClient.Close(); err != nil {
		return err
	}

	return nil
}
