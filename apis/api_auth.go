package apis

import (
	"net/url"
)

// 获取授权链接
var authUrlBase = "https://open.weixin.qq.com/connect/oauth2/authorize"

const (
	scopeBase    = "snsapi_base"        // 基础信息
	scopeInfo    = "snsapi_userinfo"    // 详细信息
	scopePrivate = "snsapi_privateinfo" // 包含手机号,已禁止
)

// redirectUrl 授权回调跳转地址
// fromSource  从哪儿授权的,可以为空
func (c *ApiClient) GetThirdPartyAuthUrl(redirectUrl, fromSource string) (string, error) {
	ret, _ := url.Parse(authUrlBase)
	query := ret.Query()
	query.Set("response_type", "code")
	query.Set("scope", scopeBase)
	query.Set("redirect_uri", redirectUrl)
	if fromSource != "" {
		query.Set("state", fromSource)
	}
	query.Set("app_id", c.AppSuiteId)

	ret.RawQuery = query.Encode()
	return ret.String() + "#wechat_redirect", nil
}
