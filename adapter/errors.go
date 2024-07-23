package adapter

import "errors"

var (
	V4APINotSupported  = errors.New("v4 API is not supported in this prism central environment")
	MethodNotSupported = errors.New("method is not supported in this prism central environment")
)
