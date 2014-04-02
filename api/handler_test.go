package api

import (
	"bytes"
	"fmt"
	"github.com/tsuru/config"
	"io"
	"io/ioutil"
	"launchpad.net/gocheck"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

func newRequestFormFile(url string, params map[string]string, paramName, path string) (*httptest.ResponseRecorder, *http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))

	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()

	if err != nil {
		return nil, nil, err
	}

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return recorder, request, err
}

func (s *S) SetUpSuite(c *gocheck.C) {
	//Loading config.yml for testing
	err := config.ReadAndWatchConfigFile("../config.yml")
	if err != nil {
		fmt.Println("Erro loading config file")
	}
}

func (s *S) TestHealthCheck(c *gocheck.C) {
	b := strings.NewReader("")
	recorder, request := get("/healthcheck", b, c)
	HealthCheckHandler(recorder, request)
	body, err := ioutil.ReadAll(recorder.Body)
	c.Assert(err, gocheck.IsNil)
	c.Assert(string(body), gocheck.Equals, "WORKING")
}

func (s *S) TestUploadFile(c *gocheck.C) {
	recorder, request, err := newRequestFormFile("/upload", nil, "file", "../fixtures/teste.jpg")
	c.Assert(err, gocheck.IsNil)
	UploadFileHandler(recorder, request)
	body, err := ioutil.ReadAll(recorder.Body)

	c.Assert(err, gocheck.IsNil)
	c.Assert(string(body), gocheck.Equals, "SUCCESS")
}
