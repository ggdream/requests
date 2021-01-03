package requests

import (
	"time"
)


// Options ...
type Options struct {
	Headers		map[string]string
	Cookies		map[string]string

	Params		map[string]interface{}
	Data		map[string]interface{}				// application/x-www-form-urlencoded
	Json		map[string]interface{}				// application/json

	Stream		bool
	Redirect	bool
	SkipVerify	bool

	Proxy		string
	Timeout		time.Duration
}
