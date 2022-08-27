package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateTag 创建标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/90346#创建标签
type ReqCreateTag struct {
	// Tagid 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。
	Tagid int `json:"tagid,omitempty"`
	// Tagname 标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名。，必填
	Tagname string `json:"tagname"`
}

var _ bodyer = ReqCreateTag{}

func (x ReqCreateTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateTag 创建标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/90346#创建标签
type RespCreateTag struct {
	CommonResp
	// Tagid 标签id
	Tagid int `json:"tagid"`
}

var _ bodyer = RespCreateTag{}

func (x RespCreateTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateTag 创建标签
// 文档：https://developer.work.weixin.qq.com/document/path/90346#创建标签
func (c *ApiClient) ExecCreateTag(req ReqCreateTag) (RespCreateTag, error) {
	var resp RespCreateTag
	err := c.executeWXApiPost("/cgi-bin/tag/create", req, &resp, true)
	if err != nil {
		return RespCreateTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateTag{}, bizErr
	}
	return resp, nil
}
