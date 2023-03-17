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
package date

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	kitdate_ "github.com/kaydxh/sea/pkg/sea-date/domain/kit/date"
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
