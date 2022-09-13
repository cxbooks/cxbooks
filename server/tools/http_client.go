package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cxbooks/cxbooks/server/zlog"
)

type HttpClient struct {
	*http.Client
	AuthHead string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy:               http.ProxyFromEnvironment,
				TLSHandshakeTimeout: 2 * time.Second,
			},
			Timeout: 30 * time.Second,
		},
	}
}

func (c *HttpClient) PATCH(url string, body []byte) ([]byte, int, error) {

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))

	return c.DoReq(req)

}

func (c *HttpClient) Post(url string, body []byte) ([]byte, int, error) {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	return c.DoReq(req)

}

func (c *HttpClient) Get(url string) ([]byte, int, error) {

	req, _ := http.NewRequest("GET", url, nil)
	return c.DoReq(req)

}

func (c *HttpClient) DoReq(req *http.Request) ([]byte, int, error) {

	req.Header.Set("Content-Type", "application/json")

	if c.AuthHead != "" {
		req.Header.Set("Cookie", c.AuthHead)
	}

	rsp, err := c.Do(req)
	if err != nil {
		zlog.E("do http request", err)
		return nil, 0, err
	}
	defer rsp.Body.Close()

	d, err := ioutil.ReadAll(rsp.Body)

	return d, rsp.StatusCode, err
}
