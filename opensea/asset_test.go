package opensea

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestOpenSeaClient_GetAssetsWithOffset(t *testing.T) {
	type fields struct {
		Log          *zap.SugaredLogger
		apiKey       string
		client       *http.Client
		baseURL      string
		limitAssets  int
		requestDelay time.Duration
	}
	type args struct {
		owner  string
		offset int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		path        string
		fixturePath string
		want        GetAssetsResponse
		wantErr     bool
	}{
		{
			name: "Get assets with offset",
			fields: fields{
				Log:          zaptest.NewLogger(t).Sugar(),
				apiKey:       "",
				client:       &http.Client{},
				baseURL:      "https://api.opensea.io",
				limitAssets:  50,
				requestDelay: time.Millisecond * 250,
			},
			args: args{
				owner:  "0x3b417FaeE9d2ff636701100891DC2755b5321Cc3",
				offset: 0,
			},
			path:        "/api/v1/assets",
			fixturePath: "../testdata/get_assets.json",
			want:        FixtureGetAssetsResp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.path {
					t.Errorf("Expected to request '%s', got: %s", tt.path, r.URL.Path)
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("Expected Content-Type: application/json header, got: %s", r.Header.Get("Accept"))
				}
				w.WriteHeader(http.StatusOK)

				// Read the fixture
				jsonFile, err := os.Open(tt.fixturePath)
				if err != nil {
					t.Errorf("Failed to open fixture file: %s", err)
				}
				defer jsonFile.Close()

				// Write the fixture to the response
				jsonFile.Seek(0, 0)
				_, err = io.Copy(w, jsonFile)
				if err != nil {
					t.Errorf("Failed to write fixture to response: %s", err)
				}
			}))
			defer server.Close()

			c := &OpenSeaClient{
				Log:          tt.fields.Log,
				apiKey:       tt.fields.apiKey,
				client:       tt.fields.client,
				baseURL:      server.URL,
				limitAssets:  tt.fields.limitAssets,
				requestDelay: tt.fields.requestDelay,
			}
			got, err := c.GetAssetsWithOffset(tt.args.owner, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenSeaClient.GetAssetsWithOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenSeaClient.GetAssetsWithOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
