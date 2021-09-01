package provider

import "github.com/jmoiron/sqlx"

// Provider for global instance
type Provider struct {
	SqlDB *sqlx.DB
}

var provider = &Provider{}

func GlobalProvider() *Provider {
	return provider
}

func GetSqlDB() *sqlx.DB {
	return provider.SqlDB
}
