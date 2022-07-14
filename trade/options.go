package trade

import (
	"github.com/longbridgeapp/openapi-go/http"
	lb "github.com/longbridgeapp/openapi-protocol/go/client"
)

const DefaultTradeUrl = "wss://openapi-trade.longbridgeapp.com"

// Options for quote context
type Options struct {
	TradeURL   string
	HttpClient http.Client
	LBClient   lb.Client
}

// Option
type Option func(*Options)

// WithTradeURL to set trade url for trade context
func WithTradeURL(url string) Option {
	return func(o *Options) {
		if url != "" {
			o.TradeURL = url
		}
	}
}

// WithHttpClient to set http client for trade context
func WithHttpClient(client http.Client) Option {
	return func(o *Options) {
		if client != nil {
			o.HttpClient = client
		}
	}
}

// WithLBClient to set Longbridge protocol client for trade context
func WithLBClient(client lb.Client) Option {
	return func(o *Options) {
		if client != nil {
			o.LBClient = client
		}
	}
}

func newOptions(opt ...Option) *Options {
	opts := Options{
		TradeURL: DefaultTradeUrl,
	}
	for _, o := range opt {
		o(&opts)
	}
	return &opts
}
