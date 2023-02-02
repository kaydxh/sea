package date

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	kitdate_ "github.com/kaydxh/sea/pkg/seadate/domain/kit/date"
)

type FactoryConfigFunc func(c *FactoryConfig) error

type FactoryConfig struct {
	Validator *validator.Validate

	DateRepository kitdate_.Repository
}

func (fc *FactoryConfig) ApplyOptions(configFuncs ...FactoryConfigFunc) error {

	for _, f := range configFuncs {
		err := f(fc)
		if err != nil {
			return fmt.Errorf("failed to apply factory config, err: %v", err)
		}
	}

	return nil
}

func (fc FactoryConfig) Validate() error {
	valid := fc.Validator
	if valid == nil {
		valid = validator.New()
	}
	return valid.Struct(fc)
}

type Factory struct {
	fc FactoryConfig
}

func NewFactory(fc FactoryConfig, configFuncs ...FactoryConfigFunc) (Factory, error) {
	err := fc.ApplyOptions(configFuncs...)
	if err != nil {
		return Factory{}, err
	}

	err = fc.Validate()
	if err != nil {
		return Factory{}, err
	}

	return Factory{fc: fc}, nil
}

func (f Factory) NewSeaDate(ctx context.Context) (*SeaDate, error) {
	s := &SeaDate{
		DateRepository: f.fc.DateRepository,
	}
	return s, nil
}
