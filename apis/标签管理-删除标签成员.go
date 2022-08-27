package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqDeltagusersTag 删除标签成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/90351#删除标签成员
type ReqDeltagusersTag struct {
	// Partylist 企业部门ID列表，注意:userlist、partylist不能同时为空，单次请求长度不超过100
	Partylist []int `json:"partylist,omitempty"`
	// Tagid 标签ID，必填
	Tagid int `json:"tagid"`
	// Userlist 企业成员ID列表，注意:userlist、partylist不能同时为空，单次请求长度不超过1000
	Userlist []string `json:"userlist,omitempty"`
}

var _ bodyer = ReqDeltagusersTag{}

func (x ReqDeltagusersTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespDeltagusersTag 删除标签成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/90351#删除标签成员
type RespDeltagusersTag struct {
	CommonResp
}

var _ bodyer = RespDeltagusersTag{}

func (x RespDeltagusersTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeltagusersTag 删除标签成员
// 文档：https://developer.work.weixin.qq.com/document/path/90351#删除标签成员
func (c *ApiClient) ExecDeltagusersTag(req ReqDeltagusersTag) (RespDeltagusersTag, error) {
	var resp RespDeltagusersTag
	err := c.executeWXApiPost("/cgi-bin/tag/deltagusers", req, &resp, true)
	if err != nil {
		return RespDeltagusersTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeltagusersTag{}, bizErr
	}
	return resp, nil
}
