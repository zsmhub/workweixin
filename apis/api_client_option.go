package apis

import (
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

const HttpTTL = 1 * time.Minute

var FastClient = CreateFastHttpClient()

func CreateFastHttpClient() fasthttp.Client {
	var defaultDialer = &fasthttp.TCPDialer{Concurrency: 300} // tcp 并发200

	return fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			idx := 3 // 重试三次
			for {
				idx--
				conn, err := defaultDialer.DialTimeout(addr, 10*time.Second) // tcp连接超时时间10s
				if err != fasthttp.ErrDialTimeout || idx == 0 {
					return conn, err
				}
			}
		},
	}
}
