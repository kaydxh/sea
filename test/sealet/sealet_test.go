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
package sealet_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"testing"
	"time"

	"github.com/google/uuid"
	http_ "github.com/kaydxh/golang/go/net/http"
	"github.com/kaydxh/sea/api/protoapi-spec/date"
	"google.golang.org/protobuf/proto"
)

const (
	serverAddr = "localhost:10001"
	timeout    = 5 * time.Second
)

func TestHttpJsonNow(t *testing.T) {
	testCases := []struct {
		expected string
	}{
		{
			expected: "",
		},
	}

	client, err := http_.NewClient(http_.WithTimeout(timeout))
	if err != nil {
		t.Fatalf("failed to new http client got: %v", err)
	}
	for i := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			req := &date.DateRequest{
				RequestId: uuid.NewString(),
			}
			dataReq, err := json.Marshal(req)
			if err != nil {
				t.Fatalf("failed to json marshal got: %v", err)
			}
			u, _ := url.Parse("http://" + serverAddr)
			u.Path = path.Join(u.Path, "Now")
			t.Logf("url: %v", u.String())

			//dataResp, err := client.PostJson("http://"+serverAddr+"/Now", nil, dataReq)
			dataResp, err := client.PostJson(u.String(), nil, dataReq)
			if err != nil {
				t.Fatalf("failed to post json got: %v", err)
			}

			var resp date.DateResponse
			err = json.Unmarshal(dataResp, &resp)
			if err != nil {
				t.Fatalf("failed to json unmarshal got: %v", err)
			}
			t.Logf("response: %v", resp.String())
		})
	}
}

func TestHttpPbNow(t *testing.T) {
	testCases := []struct {
		expected string
	}{
		{
			expected: "",
		},
	}

	client, err := http_.NewClient(http_.WithTimeout(timeout))
	if err != nil {
		t.Fatalf("failed to new http client got: %v", err)
	}
	for i := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			req := &date.DateRequest{
				RequestId: uuid.NewString(),
			}
			dataReq, err := proto.Marshal(req)
			if err != nil {
				t.Fatalf("failed to proto marshal got: %v", err)
			}

			u, _ := url.Parse("http://" + serverAddr)
			u.Path = path.Join(u.Path, "Now")
			dataResp, err := client.PostPb(u.String(), nil, dataReq)
			if err != nil {
				t.Fatalf("failed to post proto got: %v", err)
			}

			var resp date.DateResponse
			err = proto.Unmarshal(dataResp, &resp)
			if err != nil {
				t.Fatalf("failed to proto unmarshal got: %v", err)
			}
			t.Logf("response: %v", resp.String())
		})
	}
}
