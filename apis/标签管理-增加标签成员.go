package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqAddtagusersTag 增加标签成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/90350#增加标签成员
type ReqAddtagusersTag struct {
	// Partylist 企业部门ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过100
	Partylist []int `json:"partylist,omitempty"`
	// Tagid 标签ID，必填
	Tagid int `json:"tagid"`
	// Userlist 企业成员ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过1000
	Userlist []string `json:"userlist,omitempty"`
}

var _ bodyer = ReqAddtagusersTag{}

func (x ReqAddtagusersTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespAddtagusersTag 增加标签成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/90350#增加标签成员
type RespAddtagusersTag struct {
	CommonResp
}

var _ bodyer = RespAddtagusersTag{}

func (x RespAddtagusersTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddtagusersTag 增加标签成员
// 文档：https://developer.work.weixin.qq.com/document/path/90350#增加标签成员
func (c *ApiClient) ExecAddtagusersTag(req ReqAddtagusersTag) (RespAddtagusersTag, error) {
	var resp RespAddtagusersTag
	err := c.executeWXApiPost("/cgi-bin/tag/addtagusers", req, &resp, true)
	if err != nil {
		return RespAddtagusersTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddtagusersTag{}, bizErr
	}
	return resp, nil
}
