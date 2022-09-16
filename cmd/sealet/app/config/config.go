package config

import (
	"github.com/go-playground/validator/v10"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	v1 "github.com/kaydxh/sea/api/openapi-spec/sealet/v1"
	"github.com/spf13/viper"
)

type Config struct {
	Proto     v1.Configuration
	Validator *validator.Validate

	opts struct {
		viper *viper.Viper
	}
}

type completedConfig struct {
	*Config
	completeError error
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Validate checks Config.
func (c *completedConfig) Validate() error {
	return c.Validator.Struct(c)
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	err := c.loadViper()
	if err != nil {
		return CompletedConfig{
			&completedConfig{
				Config:        c,
				completeError: err,
			}}
	}

	if c.Validator == nil {
		c.Validator = validator.New()
	}

	return CompletedConfig{&completedConfig{Config: c}}
}

func (c *Config) loadViper() error {
	if c.opts.viper != nil {
		return viper_.UnmarshalProtoMessageWithJsonPb(c.opts.viper, &c.Proto)
	}

	return nil
}

func (c completedConfig) New() (*v1.Configuration, error) {
	if c.completeError != nil {
		return nil, c.completeError
	}

	return &c.Proto, nil
}

func NewConfig(options ...ConfigOption) *Config {
	var c Config
	c.ApplyOptions(options...)

	return &c
}
