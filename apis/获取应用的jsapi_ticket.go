package apis

import "net/url"

// 文档：https://work.weixin.qq.com/api/doc/90001/90144/90539#获取应用的jsapi_ticket

type JsAPITicketAgentConfigReq struct{}

var _ urlValuer = JsAPITicketAgentConfigReq{}

func (x JsAPITicketAgentConfigReq) intoURLValues() url.Values {
	return url.Values{
		"type": {"agent_config"},
	}
}

// ExecGetJSAPITicketAgentConfig 获取应用的jsapi_ticket
func (c *ApiClient) ExecGetJSAPITicketAgentConfig(req JsAPITicketAgentConfigReq) (JsAPITicketResp, error) {
	var resp JsAPITicketResp
	err := c.executeWXApiGet("/cgi-bin/ticket/get", req, &resp, true)
	if err != nil {
		return JsAPITicketResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return JsAPITicketResp{}, bizErr
	}

	return resp, nil
}
