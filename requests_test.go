package requests

import (
	"testing"
)


func TestOptions(t *testing.T) {
	// println(Get("https://www.google.com", &Options{
	// 	Timeout: 3*time.Second,
	// 	Proxy: "http://127.0.0.1:8889",
	// }).Text())

	println(Get("https://www.google.cn", nil).Text())

	// Need to release memory by `defer res.Close()`
}
