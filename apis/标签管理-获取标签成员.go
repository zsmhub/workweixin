package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetTag 获取标签成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/90349#获取标签成员
type ReqGetTag struct {
	// Tagid 标签ID，必填
	Tagid int `json:"tagid"`
}

var _ urlValuer = ReqGetTag{}

func (x ReqGetTag) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetTag 获取标签成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/90349#获取标签成员
type RespGetTag struct {
	CommonResp
	// Partylist 标签中包含的部门id列表
	Partylist []int `json:"partylist"`
	// Tagname 标签名
	Tagname  string `json:"tagname"`
	Userlist []struct {
		// Name 成员名称，代开发自建应用需要管理员授权才返回该字段；此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示名字
		Name string `json:"name"`
		// Userid 成员帐号
		Userid string `json:"userid"`
	} `json:"userlist"` // 标签中包含的成员列表
}

var _ bodyer = RespGetTag{}

func (x RespGetTag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetTag 获取标签成员
// 文档：https://developer.work.weixin.qq.com/document/path/90349#获取标签成员
func (c *ApiClient) ExecGetTag(req ReqGetTag) (RespGetTag, error) {
	var resp RespGetTag
	err := c.executeWXApiGet("/cgi-bin/tag/get", req, &resp, true)
	if err != nil {
		return RespGetTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetTag{}, bizErr
	}
	return resp, nil
}
