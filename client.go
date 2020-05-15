package htest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Client struct {
		handler http.Handler
		*testing.T
	}
)

func NewClient(t *testing.T) *Client {
	return &Client{T: t}
}

func (c Client) To(handler http.Handler) *Client {
	c.handler = handler
	return &c
}

func (c Client) ToFunc(handlerFunc http.HandlerFunc) *Client {
	c.handler = handlerFunc
	return &c
}

func (c Client) NewRequest(req *http.Request) *Request {
	return &Request{
		Request: req,
		Handler: c.handler,
		T:       c.T,
	}
}

func (c Client) request(method, path string, body io.Reader) *Request {
	req, err := http.NewRequest(method, path, body)
	assert.Nil(c.T, err)
	return c.NewRequest(req)
}

func (c Client) Get(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(GET, path, bytes.NewReader([]byte("")))
	}
	return c.request(GET, path, body[0])
}

func (c Client) Head(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(HEAD, path, bytes.NewReader([]byte("")))
	}
	return c.request(HEAD, path, body[0])
}

func (c Client) Trace(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(TRACE, path, bytes.NewReader([]byte("")))
	}
	return c.request(TRACE, path, body[0])
}

func (c Client) Options(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(OPTIONS, path, bytes.NewReader([]byte("")))
	}
	return c.request(OPTIONS, path, body[0])
}

func (c Client) Connect(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(CONNECT, path, bytes.NewReader([]byte("")))
	}
	return c.request(CONNECT, path, body[0])
}

func (c Client) Delete(path string, body ...io.Reader) *Request {
	if len(body) == 0 {
		return c.request(DELETE, path, bytes.NewReader([]byte("")))
	}
	return c.request(DELETE, path, body[0])
}

func (c Client) Post(path string, body io.Reader) *Request {
	return c.request(POST, path, body)
}

func (c Client) Put(path string, body io.Reader) *Request {
	return c.request(PUT, path, body)
}

func (c Client) Patch(path string, body io.Reader) *Request {
	return c.request(PATCH, path, body)
}
