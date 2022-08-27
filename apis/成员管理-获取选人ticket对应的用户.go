package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListSelectedTicketUserUser 成获取选人ticket对应的用户请求
// 文档：https://developer.work.weixin.qq.com/document/path/94894#成获取选人ticket对应的用户
type ReqListSelectedTicketUserUser struct {
	// SelectedTicket <a href="#30288" rel="nofollow">选人jsapi返回的selectedTicket</a>，必填
	SelectedTicket string `json:"selected_ticket"`
}

var _ bodyer = ReqListSelectedTicketUserUser{}

func (x ReqListSelectedTicketUserUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListSelectedTicketUserUser 成获取选人ticket对应的用户响应
// 文档：https://developer.work.weixin.qq.com/document/path/94894#成获取选人ticket对应的用户
type RespListSelectedTicketUserUser struct {
	CommonResp
	// OpenUseridList 此次选人操作中，在应用可见范围内的open_userid列表
	OpenUseridList []string `json:"open_userid_list"`
	// OperatorOpenUserid 选人用户的open_userid
	OperatorOpenUserid string `json:"operator_open_userid"`
	// Total 用户选择的总人数
	Total int `json:"total"`
	// UnauthOpenUseridList 此次选人操作中，不在应用可见范围内的open_userid列表
	UnauthOpenUseridList []string `json:"unauth_open_userid_list"`
}

var _ bodyer = RespListSelectedTicketUserUser{}

func (x RespListSelectedTicketUserUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListSelectedTicketUserUser 成获取选人ticket对应的用户
// 文档：https://developer.work.weixin.qq.com/document/path/94894#成获取选人ticket对应的用户
func (c *ApiClient) ExecListSelectedTicketUserUser(req ReqListSelectedTicketUserUser) (RespListSelectedTicketUserUser, error) {
	var resp RespListSelectedTicketUserUser
	err := c.executeWXApiPost("/cgi-bin/user/list_selected_ticket_user", req, &resp, true)
	if err != nil {
		return RespListSelectedTicketUserUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListSelectedTicketUserUser{}, bizErr
	}
	return resp, nil
}
