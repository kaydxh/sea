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
package seadate

import (
	"errors"

	errors_ "github.com/kaydxh/golang/go/errors"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/seadate/domain/date"
)

func DateError(err error, handled bool) (error, bool) {
	if handled {
		return err, true
	}

	if errors.Is(err, date.ErrInterval) {
		return errors_.Errorf(v1.SeaDateReasonEnum_INTERNAL, err.Error()), true
	}

	return err, false
}

func APIError(err error) error {
	if err == nil {
		return nil
	}
	return errors_.ErrorChain(DateError)(err, false)
}
