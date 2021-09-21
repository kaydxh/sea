package provider

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// Provider for global instance
type Provider struct {
	SqlDB   *sqlx.DB
	RedisDB *redis.Client
}

var provider = &Provider{}

func GlobalProvider() *Provider {
	return provider
}

func GetSqlDB() *sqlx.DB {
	return provider.SqlDB
}

func GetRedisDB() *redis.Client {
	return provider.RedisDB
}
