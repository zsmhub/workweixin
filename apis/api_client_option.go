package apis

import (
    "github.com/go-resty/resty/v2"
    "net/http"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

type AppClientOptions struct {
    WxAPIHost string
    HTTP      *http.Client
    restyCli  *resty.Client
}

// impl Default for AppClientOptions
func getDefaultAppOptions() (opt AppClientOptions) {
    opt = AppClientOptions{
        WxAPIHost: DefaultQYAPIHost,
        HTTP:      &http.Client{},
    }
    opt.restyCli = resty.NewWithClient(opt.HTTP)
    return
}

// ChangeAppClientOption 客户端对象构造参数
type ChangeAppClientOption interface {
    applyTo(*AppClientOptions)
}

type withQYAPIHost struct {
    x string
}

// WithQYAPIHost 覆盖默认企业微信 API 域名
func WithQYAPIHost(host string) ChangeAppClientOption {
    return &withQYAPIHost{x: host}
}

var _ ChangeAppClientOption = (*withQYAPIHost)(nil)

func (x *withQYAPIHost) applyTo(y *AppClientOptions) {
    if x.x != "" {
        y.WxAPIHost = x.x
    }
}

type withHTTPClient struct {
    x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) ChangeAppClientOption {
    return &withHTTPClient{x: client}
}

var _ ChangeAppClientOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) applyTo(y *AppClientOptions) {
    y.HTTP = x.x
    y.restyCli = resty.NewWithClient(x.x)
}
