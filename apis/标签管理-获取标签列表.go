package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListTag 获取标签列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/90352#获取标签列表
type ReqListTag struct{}

var _ urlValuer = ReqListTag{}

func (x ReqListTag) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListTag 获取标签列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/90352#获取标签列表
type RespListTag struct {
	CommonResp
	Taglist []struct {
		// Tagid 标签id
		Tagid int `json:"tagid"`
		// Tagname 标签名
		Tagname string `json:"tagname"`
	} `json:"taglist"` // 标签列表
}

var _ bodyer = RespListTag{}

func (x RespListTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListTag 获取标签列表
// 文档：https://developer.work.weixin.qq.com/document/path/90352#获取标签列表
func (c *ApiClient) ExecListTag(req ReqListTag) (RespListTag, error) {
	var resp RespListTag
	err := c.executeWXApiGet("/cgi-bin/tag/list", req, &resp, true)
	if err != nil {
		return RespListTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListTag{}, bizErr
	}
	return resp, nil
}
