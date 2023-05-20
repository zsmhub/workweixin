package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCustomerCustomerAcquisition 获取获客客户列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/97298#获取获客客户列表
type ReqCustomerCustomerAcquisition struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000
	Limit int `json:"limit,omitempty"`
	// LinkID 获客链接id，必填
	LinkID string `json:"link_id"`
}

var _ urlValuer = ReqCustomerCustomerAcquisition{}

func (x ReqCustomerCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespCustomerCustomerAcquisition 获取获客客户列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/97298#获取获客客户列表
type RespCustomerCustomerAcquisition struct {
	CustomerList []struct {
		// ChatStatus 会话状态，0-客户未发消息 1-客户已发送消息
		ChatStatus int `json:"chat_status"`
		// ExternalUserid 客户external_userid
		ExternalUserid string `json:"external_userid"`
		// State 用于区分客户具体是通过哪个获客链接进行添加，用户可在获客链接后拼接<code>customer_channel=自定义字符串</code>，字符串不超过64字节，超过会被截断。通过点击带有customer_channel参数的链接获取到的客户，调用获客信息接口或<a href="#13878" rel="nofollow">获取客户详情接口</a>时，返回的state参数即为链接后拼接自定义字符串
		State string `json:"state"`
		// Userid 通过获客链接添加此客户的跟进人userid
		Userid string `json:"userid"`
	} `json:"customer_list"`
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespCustomerCustomerAcquisition{}

func (x RespCustomerCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCustomerCustomerAcquisition 获取获客客户列表
// 文档：https://developer.work.weixin.qq.com/document/path/97298#获取获客客户列表
func (c *ApiClient) ExecCustomerCustomerAcquisition(req ReqCustomerCustomerAcquisition) (RespCustomerCustomerAcquisition, error) {
	var resp RespCustomerCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/customer", req, &resp, true)
	if err != nil {
		return RespCustomerCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCustomerCustomerAcquisition{}, bizErr
	}
	return resp, nil
}
