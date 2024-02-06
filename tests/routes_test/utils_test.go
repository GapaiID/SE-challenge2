package routes_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestServer struct {
	*httptest.Server
}

func NewTestServer(t *testing.T, h http.Handler) *TestServer {
	ts := httptest.NewServer(h)
	return &TestServer{ts}
}

func (ts *TestServer) Get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
