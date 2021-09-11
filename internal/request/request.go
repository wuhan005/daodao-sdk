// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package request

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Request is a request
type Request struct {
	// method is the request HTTP method.
	method string
	// path is the request API path.
	path string
	// query is the HTTP request query.
	query url.Values
	// header is the HTTP request header.
	header http.Header
	// body is the HTTP request body.
	body url.Values
}

// New creates a new Request instance with the given method and path.
func New(method, path string) *Request {
	return &Request{
		method: method,
		path:   path,
		query:  url.Values{},
		header: http.Header{},
	}
}

// Get creates an HTTP GET request with the given path.
func Get(path string) *Request { return New(http.MethodGet, path) }

// Post creates an HTTP POST request with the given path.
func Post(path string) *Request { return New(http.MethodPost, path) }

// Header replaces the request header with the given one.
func (r *Request) Header(header http.Header) *Request {
	r.header = header
	return r
}

// SetHeader sets the request header.
func (r *Request) SetHeader(k, v string) *Request {
	r.header.Set(k, v)
	return r
}

func (r *Request) Query(data url.Values) *Request {
	r.query = data
	return r
}

// Form sets the form body and `application/x-www-form-urlencoded` content-type header.
func (r *Request) Form(data url.Values) *Request {
	r.body = data
	return r
}

// Do send the HTTP request and returns the Response.
func (r *Request) Do(httpClient ...http.Client) (*Response, error) {
	client := http.Client{}
	if len(httpClient) > 0 {
		client = httpClient[0]
	}

	// Sign signature.
	timestamp := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	nonce := uuid.NewV4().String() + timestamp
	if r.method == http.MethodPost {
		r.body.Set("nonce", nonce)
		r.body.Set("timestamp", timestamp)
		r.body.Set("sign", Sign(nonce, timestamp))
	} else if r.method == http.MethodGet {
		r.query.Set("nonce", nonce)
		r.query.Set("timestamp", timestamp)
		r.query.Set("sign", Sign(nonce, timestamp))
	}

	baseURL, _ := url.Parse("https://api.daodao.cn/api")
	baseURL.Path = path.Join("/api", r.path)
	baseURL.RawQuery = r.query.Encode()

	req, err := http.NewRequest(r.method, baseURL.String(), strings.NewReader(r.body.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}

	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	r.SetHeader("User-Agent", "dao dao/1.1.7 (iPhone; iOS 14.7.1; Scale/3.00)")
	req.Header = r.header

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	return &Response{resp}, nil
}
