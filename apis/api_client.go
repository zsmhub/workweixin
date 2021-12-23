package apis

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"sync"
)

// ApiClient 企业微信客户端
type ApiClient struct {
	opts               AppClientOptions
	CorpId             string // 企业ID
	CorpProviderSecret string // 企业密钥

	// 第三方应用/服务商必填字段
	AppSuiteId     string // 第三方应用唯一身份标识
	AppSuiteSecret string // 第三方应用密钥
	AppSuiteTicket string // 企业微信服务器会定时（每十分钟）推送ticket。ticket会实时变更，并用于后续接口的调用。

	// 授权企业必填字段
	CompanyPermanentCode string // 企业授权给第三方应用的永久授权码
	AgentId              int    // 授权方应用id

	accessTokenName        string // token参数名,默认access_token; 第三方应用为suite_access_token；第三方应用服务商为provider_access_token
	accessToken            *token
	jsapiTicket            *token
	jsapiTicketAgentConfig *token

	ThirdAppClient *ApiClient // 第三方应用client，用于授权企业API客户端获取suite_token，目前用于获取企业凭证接口
}

// 服务商API客户端初始化
func NewProviderApiClient(corpId, corpProviderSecret string) *ApiClient {
	optionsObj := getDefaultAppOptions()

	c := ApiClient{
		opts:               optionsObj,
		CorpId:             corpId,
		CorpProviderSecret: corpProviderSecret,
		accessTokenName:    "provider_access_token",
		accessToken:        &token{mutex: &sync.RWMutex{}},
	}

	c.accessToken.setGetTokenFunc(c.getProviderAccessToken)
	c.SpawnAccessTokenRefresher()

	return &c
}

// 第三方应用API客户端初始化，第一次调用这个接口时，appSuiteTicket为空字符串
func NewThirdAppApiClient(corpId, appSuiteId, appSuiteSecret, appSuiteTicket string) *ApiClient {
	optionsObj := getDefaultAppOptions()

	c := ApiClient{
		opts:            optionsObj,
		CorpId:          corpId,
		AppSuiteId:      appSuiteId,
		AppSuiteSecret:  appSuiteSecret,
		AppSuiteTicket:  appSuiteTicket,
		accessTokenName: "suite_access_token",
		accessToken:     &token{mutex: &sync.RWMutex{}},
	}

	c.accessToken.setGetTokenFunc(c.getSuiteToken)
	c.SpawnAccessTokenRefresher()

	return &c
}

// 授权企业API客户端初始化
func NewAuthCorpApiClient(corpId, companyPermanentCode string, AgentId int, thirdAppClient *ApiClient) *ApiClient {
	optionsObj := getDefaultAppOptions()

	c := ApiClient{
		opts:                   optionsObj,
		CorpId:                 corpId,
		AgentId:                AgentId,
		CompanyPermanentCode:   companyPermanentCode,
		accessTokenName:        "access_token",
		accessToken:            &token{mutex: &sync.RWMutex{}},
		jsapiTicket:            &token{mutex: &sync.RWMutex{}},
		jsapiTicketAgentConfig: &token{mutex: &sync.RWMutex{}},
		ThirdAppClient:         thirdAppClient,
	}

	c.accessToken.setGetTokenFunc(c.getAuthCorpToken)
	c.SpawnAccessTokenRefresher()

	c.jsapiTicket.setGetTokenFunc(c.getJSAPITicket)
	c.SpawnJSAPITicketRefresher()

	c.jsapiTicketAgentConfig.setGetTokenFunc(c.getJSAPITicketAgentConfig)
	c.SpawnJSAPITicketAgentConfigRefresher()

	return &c
}

// 更新suite_ticket
func (c *ApiClient) RefreshSuiteTicket(ticket string) {
	c.AppSuiteTicket = ticket
}

func (c *ApiClient) composeWXApiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	base, err := url.Parse(c.opts.WxAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.WxAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *ApiClient) composeWXURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	wxApiURL := c.composeWXApiURL(path, req)

	if !withAccessToken {
		return wxApiURL
	}

	q := wxApiURL.Query()
	if wxApiURL.Path == PathGetCorpToken { // 兼容获取企业凭证接口
		q.Set("suite_access_token", c.ThirdAppClient.accessToken.getToken())
	} else {
		q.Set(c.accessTokenName, c.accessToken.getToken())
	}
	wxApiURL.RawQuery = q.Encode()

	return wxApiURL
}

func (c *ApiClient) executeWXApiGet(path string, req urlValuer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	commonResp := CommonResp{}
	resp, err := c.opts.restyCli.R().Get(urlStr)
	if err != nil {
		return err
	}

	bodyResp := resp.Body()
	if commonResp.ErrCode == ErrCode42001 || commonResp.ErrCode == ErrCode640014 {
		log.Println("invalid access_token,now retry")
		return err
	}
	err = json.Unmarshal(bodyResp, &objResp)
	return err
}

// 微信端接收的参数中一个数组里包含有多种类型，强类型语言无法支持，只能在前端拼接成str直接传到wx
func (c *ApiClient) executeWXApiJSONPostWithBytesReq(path string, req []byte, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	// resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(req))
	resp, err := c.opts.restyCli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}

	err = json.Unmarshal(resp.Body(), &objResp)

	return err
}

func (c *ApiClient) executeWXApiPost(path string, req bodyer, objResp interface{}, withAccessToken bool) error {
	// defer util.FuncTracer("path", path, "req", req, "resp", objResp)()
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	body, err := req.intoBody()
	if err != nil {
		// TODO: error_chain
		return err
	}

	resp, err := c.opts.restyCli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}

	err = json.Unmarshal(resp.Body(), &objResp)
	return err

}

func (c *ApiClient) executeWXApiMediaUpload(path string, req mediaUploader, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()
	m := req.getMedia()
	resp, err := c.opts.restyCli.R().
		SetFileReader("media", m.filename, m.stream).
		Post(urlStr)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(resp.Body(), &objResp)
	return err
}
