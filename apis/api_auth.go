package apis

import (
	"net/url"
)

// 获取授权链接
var authUrlBase = "https://open.weixin.qq.com/connect/oauth2/authorize"

// 应用授权作用域
const (
	scopeBase    = "snsapi_base"        // 静默授权，可获取成员的基础信息（UserId与DeviceId）
	scopeInfo    = "snsapi_userinfo"    // 静默授权，可获取成员的详细信息，但不包含头像、二维码等敏感信息
	scopePrivate = "snsapi_privateinfo" // 手动授权，可获取成员的详细信息，包含头像、二维码等敏感信息
)

type GetThirdOauthUrlReq struct {
	RedirectUri string `json:"redirect_uri"` // 授权回调跳转地址
	Scope       string `json:"scope"`        // 应用授权作用域
	State       string `json:"state"`        // 重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值，长度不可超过128个字节
}

// 构造oauth2链接
func (c *ApiClient) GetThirdOauthUrl(req GetThirdOauthUrlReq) string {
	ret, _ := url.Parse(authUrlBase)
	query := ret.Query()
	query.Set("appid", c.AppSuiteId)
	query.Set("redirect_uri", urlEncode(req.RedirectUri))
	query.Set("response_type", "code")
	query.Set("scope", req.Scope)
	if req.State != "" {
		query.Set("state", req.State)
	}
	ret.RawQuery = query.Encode()
	return ret.String() + "#wechat_redirect"
}

func urlEncode(str string) string {
	return url.QueryEscape(str)
}
