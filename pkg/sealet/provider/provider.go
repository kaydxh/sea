package provider

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	reslover_ "github.com/kaydxh/golang/pkg/reslover"
)

// Provider for global instance
type Provider struct {
	SqlDB           *sqlx.DB
	RedisDB         *redis.Client
	ResolverService *reslover_.ResloverService
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

func GetResolverService() *reslover_.ResloverService {
	return provider.ResolverService
}
