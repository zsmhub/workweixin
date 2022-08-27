package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqMarkTagExternalcontact 编辑客户企业标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/92697#编辑客户企业标签
type ReqMarkTagExternalcontact struct {
	// AddTag 要标记的标签列表
	AddTag []string `json:"add_tag,omitempty"`
	// ExternalUserid 外部联系人userid，必填
	ExternalUserid string `json:"external_userid"`
	// RemoveTag 要移除的标签列表
	RemoveTag []string `json:"remove_tag,omitempty"`
	// Userid 添加外部联系人的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqMarkTagExternalcontact{}

func (x ReqMarkTagExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespMarkTagExternalcontact 编辑客户企业标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/92697#编辑客户企业标签
type RespMarkTagExternalcontact struct {
	CommonResp
}

var _ bodyer = RespMarkTagExternalcontact{}

func (x RespMarkTagExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecMarkTagExternalcontact 编辑客户企业标签
// 文档：https://developer.work.weixin.qq.com/document/path/92697#编辑客户企业标签
func (c *ApiClient) ExecMarkTagExternalcontact(req ReqMarkTagExternalcontact) (RespMarkTagExternalcontact, error) {
	var resp RespMarkTagExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/mark_tag", req, &resp, true)
	if err != nil {
		return RespMarkTagExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMarkTagExternalcontact{}, bizErr
	}
	return resp, nil
}
