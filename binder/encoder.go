package binder

import (
	"net/http"

	"github.com/afikrim/pot/option"
)

type RequestEncoder struct {
	Opts    *option.BinderOptions
	Request *http.Request
}

type ResponseEncoder struct {
	ResponseWriter http.ResponseWriter
}

func NewRequestEncoder(r *http.Request, opts ...option.BinderOption) *RequestEncoder {
	return &RequestEncoder{
		Opts:    option.NewBinderOptions(opts...),
		Request: r,
	}
}
