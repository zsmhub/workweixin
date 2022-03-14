package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
)

// ApiClient 企业微信客户端
type ApiClient struct {
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
	c := ApiClient{
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
	c := ApiClient{
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
	c := ApiClient{
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

	base, err := url.Parse(DefaultQYAPIHost)
	if err != nil {
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", DefaultQYAPIHost, err))
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

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetMethod(http.MethodGet)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	return json.Unmarshal(httpResp.Body(), &objResp)
}

func (c *ApiClient) executeWXApiPost(path string, req bodyer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	body, err := req.intoBody()
	if err != nil {
		return err
	}

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType("application/json")
	httpReq.SetBody(body)
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	return json.Unmarshal(httpResp.Body(), &objResp)
}

func (c *ApiClient) executeWXApiMediaUpload(path string, req mediaUploader, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)

	urlStr := wxUrlWithToken.String()

	m := req.getMedia()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	// 新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	// 创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	// 从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("media", m.filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = io.Copy(fileWriter, m.stream)
	if err != nil {
		return err
	}

	// 停止写入
	_ = bodyWriter.Close()

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType(bodyWriter.FormDataContentType())
	httpReq.SetBody(bodyBufer.Bytes())
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	return json.Unmarshal(httpResp.Body(), &objResp)
}
