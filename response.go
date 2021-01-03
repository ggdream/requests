package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"unsafe"
)



// Response your result
type Response struct {
	headers			http.Header
	cookies			[]*http.Cookie

	rawResponse		*http.Response
}


// Body raw reponse body
func (r *Response) Body() io.ReadCloser {
	return r.rawResponse.Body
}

// Data body byte slice
func (r *Response) Data() []byte {
	data, err :=  ioutil.ReadAll(r.Body())
	if err != nil {
		panic(err)
	}
	return data
}

// Text body string
func (r *Response) Text() string {
	data := r.Data()
	return *(*string)(unsafe.Pointer(&data))
}

// Json body json string
func (r *Response) Json() (string, error) {
	data, err := json.Marshal(r.Data())
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&data)), nil
}

// Headers response headers
func (r *Response) Headers() http.Header {
	return r.headers
}

// Cookies response cookies
func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

// Raw get raw response
func (r *Response) Raw() *http.Response {
	return r.rawResponse
}

// Close close raw response body
func (r *Response) Close() error {
	return r.rawResponse.Body.Close()
}

