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

	errors_ "github.com/kaydxh/golang/go/errors"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	kitdate_ "github.com/kaydxh/sea/pkg/seadate/domain/kit/date"
)

var _ Repository = (*SeaDate)(nil)

type SeaDate struct {
	DateRepository kitdate_.Repository
}

type NowRequest struct {
	RequestId string
}

type NowResponse struct {
	Date string
}

type NowErrorRequest struct {
	RequestId string
}

type NowErrorResponse struct {
	Date string
}

func (s *SeaDate) Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.NowRequest{}
	dateResp, err := s.DateRepository.Now(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call Now, err: %v", err)
		return nil, errors_.Errore(
			err,
			ErrInterval,
		)
	}

	resp = &NowResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}

func (s *SeaDate) NowError(ctx context.Context, req *NowErrorRequest) (resp *NowErrorResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.NowErrorRequest{}
	dateResp, err := s.DateRepository.NowError(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call NowError, err: %v", err)
		return nil, errors_.Errore(
			err,
			ErrInterval,
		)
	}

	resp = &NowErrorResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
