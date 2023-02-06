package shippinglabel

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	HeaderContentTypeJSON = "application/json; charset=utf-8"
	HeaderContentTypeForm = "application/x-www-form-urlencoded"
)

type BodyParser func() (io.ReadCloser, error)

type ResponseHandler = func(*http.Response) error

type request struct {
	baseURL     string
	path        string
	method      string
	body        BodyParser
	respHandler ResponseHandler
	headers     map[string]string
}

func newRequest(url string) *request {
	return &request{
		baseURL: url,
		method:  http.MethodGet,
	}
}

// SetURL sets the base URL
func (r *request) SetURL(url string) *request {
	r.baseURL = url
	return r
}

// SetPath sets the request path
func (r *request) SetPath(path string) *request {
	r.path = path
	return r
}

// SetPathf sets the request path
func (r *request) SetPathf(format string, a ...any) *request {
	return r.SetPath(fmt.Sprintf(format, a...))
}

// SetMethod sets the http request method
func (r *request) SetMethod(method string) *request {
	r.method = method
	return r
}

// Header

// SetHeader adds an http header
func (r *request) SetHeader(key string, value string) *request {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[http.CanonicalHeaderKey(key)] = value
	return r
}

// GetHeader returns a header by key
func (r *request) GetHeader(key string) string {
	if r.headers == nil {
		return ""
	}

	return r.headers[key]
}

// SetContentType sets the Content-Type header
func (r *request) SetContentType(contentType string) *request {
	return r.SetHeader("Content-Type", contentType)
}

// SetAccept sets the Accept header
func (r *request) SetAccept(contentType string) *request {
	return r.SetHeader("Accept", contentType)
}

// SetBearer sets a bearer token
func (r *request) SetBearer(token string) *request {
	return r.SetHeader("Authorization", "Bearer "+token)
}

// SetBasicAuth sets the basic auth header
func (r *request) SetBasicAuth(username string, password string) *request {
	return r.SetHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
}

// Body

// SetJSON converts a struct into a JSON object and sets the json Content-Type header
func (r *request) SetJSON(body any) *request {
	r.body = func() (io.ReadCloser, error) {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(b)), nil
	}
	return r.SetContentType(HeaderContentTypeJSON)
}

// SetFormValues sets the form values body
func (r *request) SetFormValues(val url.Values) *request {
	r.body = func() (r io.ReadCloser, err error) {
		return io.NopCloser(strings.NewReader(val.Encode())), nil
	}
	return r.SetContentType(HeaderContentTypeForm)
}

// Response

// ToJSON converts the JSON response to the specified variable
func (r *request) ToJSON(v any) *request {
	r.respHandler = func(res *http.Response) error {
		return json.NewDecoder(res.Body).Decode(v)
	}
	return r.SetAccept(HeaderContentTypeJSON)
}

// ToBytesBuffer writes the response body to the bytes.Buffer
func (r *request) ToBytesBuffer(buf *bytes.Buffer) *request {
	r.respHandler = func(res *http.Response) error {
		_, err := io.Copy(buf, res.Body)
		return err
	}
	return r
}

// HTTPRequest converts the Request struct into a http.Request
func (r *request) HTTPRequest(ctx context.Context) (req *http.Request, err error) {
	var body io.ReadCloser
	if r.body != nil {
		body, err = r.body()
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequestWithContext(ctx, r.method, r.baseURL+r.path, body)
	if err != nil {
		return nil, err
	}

	if len(r.headers) > 0 {
		for key, val := range r.headers {
			req.Header.Add(key, val)
		}
	}

	return req, nil
}
