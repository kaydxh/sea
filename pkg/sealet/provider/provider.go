/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
package provider

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	resolver_ "github.com/kaydxh/golang/pkg/resolver"
	v1 "github.com/kaydxh/sea/api/openapi-spec/sealet/v1"
)

// Provider for global instance
type Provider struct {
	Config *v1.Configuration

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
