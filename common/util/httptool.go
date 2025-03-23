package util

import (
	"context"
	"net/http"
	"time"
)

// 请求方法
// 请求头
// 请求体
// traceId

type requestOption struct {
	ctx     context.Context
	header  map[string]string
	data    []byte
	timeout time.Duration
}

type Option func(*requestOption) error

func (f Option) Apply(opts *requestOption) error {
	return f(opts)
}

func defaultRequestOption() *requestOption {
	return &requestOption{
		ctx:     context.Background(),
		header:  map[string]string{},
		timeout: 10 * time.Second,
	}
}

func WithContext(ctx context.Context) Option {
	return Option(func(req *requestOption) error {
		req.ctx = ctx
		return nil
	})
}

func WithHeader(header map[string]string) Option {
	return Option(func(req *requestOption) error {
		req.header = header
		return nil
	})
}

func WithData(data []byte) Option {
	return Option(func(req *requestOption) error {
		req.data = data
		return nil
	})
}

func WithTimeout(timeout time.Duration) Option {
	return Option(func(req *requestOption) error {
		req.timeout = timeout
		return nil
	})
}

func Request(method string, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	start := time.Now()
	reqOpts := defaultRequestOption()
	for i := range opts {
		if err = opts[i].Apply(reqOpts); err != nil {
			return 0, nil, err
		}
	}
}

func Get(ctx context.Context, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	opts = append(opts, WithContext(ctx))
	return Request("GET", url, opts...)
}

func Post(ctx context.Context, data []byte, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	opts = append(opts,
		WithContext(ctx),
		WithHeader(map[string]string{"Content-Type": "application/json"}),
		WithData(data))
	return Request("POST", url, opts...)
}
