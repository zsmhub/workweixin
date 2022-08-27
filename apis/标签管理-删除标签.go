package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqDeleteTag 删除标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/90348#删除标签
type ReqDeleteTag struct {
	// Tagid 标签ID，必填
	Tagid int `json:"tagid"`
}

var _ urlValuer = ReqDeleteTag{}

func (x ReqDeleteTag) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDeleteTag 删除标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/90348#删除标签
type RespDeleteTag struct {
	CommonResp
}

var _ bodyer = RespDeleteTag{}

func (x RespDeleteTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeleteTag 删除标签
// 文档：https://developer.work.weixin.qq.com/document/path/90348#删除标签
func (c *ApiClient) ExecDeleteTag(req ReqDeleteTag) (RespDeleteTag, error) {
	var resp RespDeleteTag
	err := c.executeWXApiGet("/cgi-bin/tag/delete", req, &resp, true)
	if err != nil {
		return RespDeleteTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteTag{}, bizErr
	}
	return resp, nil
}
