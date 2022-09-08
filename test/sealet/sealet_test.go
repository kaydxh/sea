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
	v1 "github.com/kaydxh/sea/api/openapi-spec/sealet/date/v1"
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
			req := &v1.DateRequest{
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

			var resp v1.DateResponse
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
			req := &v1.DateRequest{
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

			var resp v1.DateResponse
			err = proto.Unmarshal(dataResp, &resp)
			if err != nil {
				t.Fatalf("failed to proto unmarshal got: %v", err)
			}
			t.Logf("response: %v", resp.String())
		})
	}
}
