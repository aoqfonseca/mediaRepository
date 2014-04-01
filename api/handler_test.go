package api

import (
	"io"
	"io/ioutil"
	"launchpad.net/gocheck"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct{}

var _ = gocheck.Suite(&S{})

func request(method, url string, b io.Reader, c *gocheck.C) (*httptest.ResponseRecorder, *http.Request) {
	request, err := http.NewRequest(method, url, b)
	c.Assert(err, gocheck.IsNil)
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	return recorder, request
}

func get(url string, b io.Reader, c *gocheck.C) (*httptest.ResponseRecorder, *http.Request) {
	return request("GET", url, b, c)
}

func (s *S) TestHealthCheck(c *gocheck.C) {
	b := strings.NewReader("")
	recorder, request := get("/healthcheck", b, c)
	HealthCheckHandler(recorder, request)
	body, err := ioutil.ReadAll(recorder.Body)
	c.Assert(err, gocheck.IsNil)
	c.Assert(string(body), gocheck.Equals, "WORKING")
}
