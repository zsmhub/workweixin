package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSimplelistUser 获取部门成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/90336#获取部门成员
type ReqSimplelistUser struct {
	// DepartmentID 获取的部门id，必填
	DepartmentID int `json:"department_id"`
	// FetchChild 是否递归获取子部门下面的成员:1-递归获取，0-只获取本部门
	FetchChild int `json:"fetch_child"`
}

var _ urlValuer = ReqSimplelistUser{}

func (x ReqSimplelistUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespSimplelistUser 获取部门成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/90336#获取部门成员
type RespSimplelistUser struct {
	CommonResp
	Userlist []struct {
		// Department 成员所属部门列表。列表项为部门ID，32位整型
		Department []int `json:"department"`
		// Name 成员名称，代开发自建应用需要管理员授权才返回；此字段从2019年12月30日起，对新创建第三方应用不再返回真实name，使用userid代替name，2020年6月30日起，对所有历史第三方应用不再返回真实name，使用userid代替name，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示名字
		Name string `json:"name"`
		// OpenUserid 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
		OpenUserid string `json:"open_userid"`
		// Userid 成员UserID。对应管理端的帐号
		Userid string `json:"userid"`
	} `json:"userlist"` // 成员列表
}

var _ bodyer = RespSimplelistUser{}

func (x RespSimplelistUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSimplelistUser 获取部门成员
// 文档：https://developer.work.weixin.qq.com/document/path/90336#获取部门成员
func (c *ApiClient) ExecSimplelistUser(req ReqSimplelistUser) (RespSimplelistUser, error) {
	var resp RespSimplelistUser
	err := c.executeWXApiGet("/cgi-bin/user/simplelist", req, &resp, true)
	if err != nil {
		return RespSimplelistUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSimplelistUser{}, bizErr
	}
	return resp, nil
}
