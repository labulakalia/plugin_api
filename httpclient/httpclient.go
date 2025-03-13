package httpclient

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	userAgent = "Mediagate/v0.0.1"
)

func GetUserAgent() string {
	return userAgent
}

func SetUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", userAgent)
	return
}

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	SetHttpUserAgent(req)
	return c.Client.Do(req)
}
func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	SetHttpUserAgent(req)
	return c.Client.Do(req)
}
func (c *Client) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	SetHttpUserAgent(req)
	return c.Client.Do(req)
}
func (c *Client) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	SetHttpUserAgent(req)

	return c.Client.Do(req)
}
func (c *Client) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
