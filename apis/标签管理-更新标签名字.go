package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUpdateTag 更新标签名字请求
// 文档：https://developer.work.weixin.qq.com/document/path/90347#更新标签名字
type ReqUpdateTag struct {
	// Tagid 标签ID，必填
	Tagid int `json:"tagid"`
	// Tagname 标签名称，长度限制为32个字（汉字或英文字母），标签不可与其他标签重名。，必填
	Tagname string `json:"tagname"`
}

var _ bodyer = ReqUpdateTag{}

func (x ReqUpdateTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateTag 更新标签名字响应
// 文档：https://developer.work.weixin.qq.com/document/path/90347#更新标签名字
type RespUpdateTag struct {
	CommonResp
}

var _ bodyer = RespUpdateTag{}

func (x RespUpdateTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateTag 更新标签名字
// 文档：https://developer.work.weixin.qq.com/document/path/90347#更新标签名字
func (c *ApiClient) ExecUpdateTag(req ReqUpdateTag) (RespUpdateTag, error) {
	var resp RespUpdateTag
	err := c.executeWXApiPost("/cgi-bin/tag/update", req, &resp, true)
	if err != nil {
		return RespUpdateTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateTag{}, bizErr
	}
	return resp, nil
}
