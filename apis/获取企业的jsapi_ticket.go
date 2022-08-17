package apis

import "net/url"

// 文档：https://developer.work.weixin.qq.com/document/path/90539#获取企业的jsapi_ticket

type JsAPITicketReq struct{}

var _ urlValuer = JsAPITicketReq{}

func (x JsAPITicketReq) intoURLValues() url.Values {
	return url.Values{}
}

type JsAPITicketResp struct {
	CommonResp

	Ticket        string `json:"ticket"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// ExecGetJSAPITicket 获取企业的jsapi_ticket
func (c *ApiClient) ExecGetJSAPITicket(req JsAPITicketReq) (JsAPITicketResp, error) {
	var resp JsAPITicketResp
	err := c.executeWXApiGet("/cgi-bin/get_jsapi_ticket", req, &resp, true)
	if err != nil {
		return JsAPITicketResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return JsAPITicketResp{}, bizErr
	}
	return resp, nil
}
