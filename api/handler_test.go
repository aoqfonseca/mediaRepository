package api

import (
    "net/http",
    "net/http/httptest",
    "testing",
    "strings",
    "gocheck"
)

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


func (s *S) TestHealthCheck(c *gocheck.C){
    b := io.NewReader()
    recorder, request := get('/healthcheck', b, c)
    HealthCheckHandler(recorder, request)
    boy, err := io.ReadAll(recorder.Body)
    c.Assert(err, gocheck.isNill)
    c.Assert(body, gocheck.Equals,"WORKING")
}

