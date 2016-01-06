package tldr_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pranavraja/tldr/lib/tldr"
)

type testServer struct {
	originalRequest *http.Request
	statusCode      int
	response        string
}

func (t *testServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.originalRequest = r
	w.WriteHeader(t.statusCode)
	io.WriteString(w, t.response)
}

func (t *testServer) Intercept(test func()) {
	server := httptest.NewServer(t)
	defer server.Close()
	remote = server.URL
	test()
}

func TestGetPageForPlatform_404(t *testing.T) {
	server := testServer{statusCode: 404, response: "NOT FOUND BRO"}
	var resp io.ReadCloser
	var err error
	server.Intercept(func() {
		resp, err = tldr.GetPageForPlatform("tldr", "osx")
	})
	if resp != nil {
		t.Errorf("Expected a nil response but got a non-nil response")
	}
	if err == nil {
		t.Errorf("Expected an error for an invalid status code, but got none")
	}
}

func TestGetPageForPlatform(t *testing.T) {
	server := testServer{statusCode: 200, response: "DO IT BRO"}
	var resp io.ReadCloser
	var err error
	server.Intercept(func() {
		resp, err = tldr.GetPageForPlatform("tldr", "osx")
	})
	defer resp.Close()
	if err != nil {
		t.Error(err)
	}
	if expected := "/osx/tldr.md"; server.originalRequest.URL.Path != expected {
		t.Errorf("Page requested from wrong url: %s", server.originalRequest.URL.Path)
	}
	if body, _ := ioutil.ReadAll(resp); string(body) != "DO IT BRO" {
		t.Errorf("Read wrong body: %s")
	}
}
