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
	"time"

	kitdate_ "github.com/kaydxh/sea/pkg/seadate/domain/kit/date"
)

var _ kitdate_.Repository = (*Repository)(nil)

type Repository struct {
}

func (r *Repository) Now(ctx context.Context, req *kitdate_.NowRequest) (resp *kitdate_.NowResponse, err error) {
	resp = &kitdate_.NowResponse{
		Date: time.Now().String(),
	}
	return resp, nil
}

func (r *Repository) NowError(ctx context.Context, req *kitdate_.NowErrorRequest) (resp *kitdate_.NowErrorResponse, err error) {
	err = fmt.Errorf("Internal")
	return nil, fmt.Errorf("Internal")
}
