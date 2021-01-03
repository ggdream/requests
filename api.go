package requests

import (
	"net/http"
)



func request(method string, url string, options *Options) *Response {
	session :=  New(options)
	res, err := session.Request(method, url, options)
	if err != nil {
		panic(err)
	}

	return res
}

// Get ...
func Get(url string, options *Options) *Response {
	return request(http.MethodGet, url, options)
}
// Post ...
func Post(url string, options *Options) *Response {
	return request(http.MethodPost, url, options)
}
