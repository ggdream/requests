package requests

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Session ...
type Session struct {
	*http.Client
}

// New ...
func New(options *Options) *Session {
	if options == nil {
		return &Session{
			&http.Client{},
		}
	}

	proxyFunc := func () func(r *http.Request) (*url.URL, error) {
		if options.Proxy == "" {
			return nil							// 不设置代理
		}

		proxy, err := url.Parse(options.Proxy)
		if err != nil {
			panic(err)
		}
		return http.ProxyURL(proxy)				// 设置代理
	}

	session := &Session{
		&http.Client{
			Transport: &http.Transport{
				Proxy: proxyFunc(),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: options.SkipVerify,
				},
			},
			Timeout: options.Timeout,
		},
	}

	return session
}



// Request ...
func (s *Session) Request(method string, url string, options *Options) (*Response, error) {
	reader, ctype := func () (io.Reader, string) {
		if options == nil {
			return nil, ""
		}
		return s.args(options)
	}()

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	if options != nil {
		if ctype != "" {
			req.Header.Set("Content-Type", ctype)
		}
		for k, v := range options.Headers {
			req.Header.Set(k, v)
		}
	
		for k, v := range options.Cookies {
			req.AddCookie(&http.Cookie{
				Name: k,
				Value: v,
			})
		}
	}


	res, err := s.Do(req)
	if err != nil {
		return nil, err
	}

	response := &Response{
		rawResponse: res,

		headers: res.Header,
		cookies: res.Cookies(),
	}

	return response, nil
}



func (s *Session) args(options *Options) (io.Reader, string) {
	var reader	io.Reader
	var types	string

	trans := func (data map[string]interface{}) io.Reader {
		res, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		return bytes.NewBuffer(res)
	}

	if options.Params != nil {
		reader = trans(options.Params)
	} else if options.Data != nil {
		types = "application/x-www-form-urlencoded"
		reader = trans(options.Data)
	} else if options.Json != nil {
		types = "application/json"
		reader = trans(options.Json)
	} else {
		reader = nil
	}

	return reader, types
}
