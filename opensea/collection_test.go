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

func TestOpenSeaClient_GetCollection(t *testing.T) {
	type fields struct {
		Log          *zap.SugaredLogger
		apiKey       string
		client       *http.Client
		baseURL      string
		limitAssets  int
		requestDelay time.Duration
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		path        string
		fixturePath string
		want        Collection
		wantErr     bool
	}{
		{
			name: "Get collection",
			fields: fields{
				Log:          zaptest.NewLogger(t).Sugar(),
				apiKey:       "",
				client:       &http.Client{},
				baseURL:      "https://api.opensea.io",
				limitAssets:  50,
				requestDelay: time.Millisecond * 250,
			},
			args: args{
				slug: "boredapeyachtclub",
			},
			path:        "/api/v1/collection/boredapeyachtclub",
			fixturePath: "../testdata/get_collection.json",
			want:        FixtureGetCollectionResp,
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
			got, err := c.GetCollection(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenSeaClient.GetCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenSeaClient.GetCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
