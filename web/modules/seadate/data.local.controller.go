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
	"context"

	middleware_ "github.com/kaydxh/golang/pkg/middleware"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
)

type LocalController struct {
	*Controller
}

// 日期查询
func (c *LocalController) Now(
	ctx context.Context,
	req *v1.NowRequest,
) (resp *v1.NowResponse, err error) {

	return middleware_.LocalMiddlewareWrap(c.Controller.Now)(ctx, req)
}

func (c *LocalController) NowError(
	ctx context.Context,
	req *v1.NowErrorRequest,
) (resp *v1.NowErrorResponse, err error) {
	return middleware_.LocalMiddlewareWrap(c.Controller.NowError)(ctx, req)
}
