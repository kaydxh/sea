package provider

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	resolver_ "github.com/kaydxh/golang/pkg/resolver"
)

// Provider for global instance
type Provider struct {
	SqlDB           *sqlx.DB
	RedisDB         *redis.Client
	ResolverService *resolver_.ResolverService
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

func GetResolverService() *resolver_.ResolverService {
	return provider.ResolverService
}
