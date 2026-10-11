package trackerapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestClient(t *testing.T, cfg Config, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	cfg.BaseURL = srv.URL
	return New(cfg)
}

func TestClientDoHeaders(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want map[string]string
	}{
		{
			name: "token only",
			cfg:  Config{Token: "tok"},
			want: map[string]string{"Authorization": "OAuth tok", "X-Org-ID": "", "X-Cloud-Org-ID": ""},
		},
		{
			name: "with org id",
			cfg:  Config{Token: "tok", OrgID: "org1"},
			want: map[string]string{"Authorization": "OAuth tok", "X-Org-ID": "org1", "X-Cloud-Org-ID": ""},
		},
		{
			name: "with cloud org id",
			cfg:  Config{Token: "tok", CloudOrgID: "cloud1"},
			want: map[string]string{"Authorization": "OAuth tok", "X-Org-ID": "", "X-Cloud-Org-ID": "cloud1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got http.Header
			c := newTestClient(t, tt.cfg, func(w http.ResponseWriter, r *http.Request) {
				got = r.Header
				w.WriteHeader(http.StatusOK)
			})
			if _, err := c.do(t.Context(), http.MethodGet, "/x", nil, nil); err != nil {
				t.Fatalf("do() error = %v", err)
			}
			for header, want := range tt.want {
				if got.Get(header) != want {
					t.Errorf("header %s = %q, want %q", header, got.Get(header), want)
				}
			}
		})
	}
}

func TestClientDoRequestBody(t *testing.T) {
	t.Run("nil body sends no body or content-type", func(t *testing.T) {
		var gotContentType string
		var gotBody []byte
		c := newTestClient(t, Config{}, func(w http.ResponseWriter, r *http.Request) {
			gotContentType = r.Header.Get("Content-Type")
			gotBody, _ = io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
		})
		if _, err := c.do(t.Context(), http.MethodGet, "/x", nil, nil); err != nil {
			t.Fatalf("do() error = %v", err)
		}
		if gotContentType != "" {
			t.Errorf("Content-Type = %q, want empty", gotContentType)
		}
		if len(gotBody) != 0 {
			t.Errorf("body = %q, want empty", gotBody)
		}
	})

	t.Run("non-nil body is encoded as JSON", func(t *testing.T) {
		var gotContentType string
		var gotBody map[string]any
		c := newTestClient(t, Config{}, func(w http.ResponseWriter, r *http.Request) {
			gotContentType = r.Header.Get("Content-Type")
			_ = json.NewDecoder(r.Body).Decode(&gotBody)
			w.WriteHeader(http.StatusOK)
		})
		body := map[string]any{"summary": "hello"}
		if _, err := c.do(t.Context(), http.MethodPost, "/x", nil, body); err != nil {
			t.Fatalf("do() error = %v", err)
		}
		if gotContentType != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", gotContentType)
		}
		if gotBody["summary"] != "hello" {
			t.Errorf("body = %v, want summary=hello", gotBody)
		}
	})
}

func TestClientDoQueryParams(t *testing.T) {
	var gotQuery url.Values
	c := newTestClient(t, Config{}, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query()
		w.WriteHeader(http.StatusOK)
	})
	q := url.Values{"perPage": []string{"50"}, "id": []string{"10"}}
	if _, err := c.do(t.Context(), http.MethodGet, "/x", q, nil); err != nil {
		t.Fatalf("do() error = %v", err)
	}
	if gotQuery.Get("perPage") != "50" || gotQuery.Get("id") != "10" {
		t.Errorf("query = %v, want perPage=50&id=10", gotQuery)
	}
}

func TestClientDoResponses(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		respBody   string
		wantRaw    string
		wantErr    *APIError
	}{
		{
			name:       "2xx passes body through unchanged",
			statusCode: http.StatusOK,
			respBody:   `{"id":"1"}`,
			wantRaw:    `{"id":"1"}`,
		},
		{
			name:       "2xx with empty body returns null",
			statusCode: http.StatusNoContent,
			respBody:   "",
			wantRaw:    "null",
		},
		{
			name:       "non-2xx returns APIError",
			statusCode: http.StatusNotFound,
			respBody:   `{"error":"not found"}`,
			wantErr:    &APIError{StatusCode: http.StatusNotFound, Body: `{"error":"not found"}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestClient(t, Config{}, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.respBody))
			})
			raw, err := c.do(t.Context(), http.MethodGet, "/x", nil, nil)

			if tt.wantErr != nil {
				var apiErr *APIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("err = %v, want *APIError", err)
				}
				if apiErr.StatusCode != tt.wantErr.StatusCode || apiErr.Body != tt.wantErr.Body {
					t.Errorf("APIError = %+v, want %+v", apiErr, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("do() error = %v", err)
			}
			if string(raw) != tt.wantRaw {
				t.Errorf("raw = %q, want %q", raw, tt.wantRaw)
			}
		})
	}
}
