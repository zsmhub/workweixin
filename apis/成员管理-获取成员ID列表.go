package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListIdUser 获取成员ID列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/96021#获取成员ID列表
type ReqListIdUser struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 10000
	Limit int `json:"limit,omitempty"`
}

var _ bodyer = ReqListIdUser{}

func (x ReqListIdUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListIdUser 获取成员ID列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/96021#获取成员ID列表
type RespListIdUser struct {
	CommonResp
	NextCursor string `json:"next_cursor"`
	DeptUser   []struct {
		OpenUserid string `json:"open_userid"`
		Department int    `json:"department"`
	} `json:"dept_user"`
}

var _ bodyer = RespListIdUser{}

func (x RespListIdUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListIdUser 获取成员ID列表
// 文档：https://developer.work.weixin.qq.com/document/path/96021#获取成员ID列表
func (c *ApiClient) ExecListIdUser(req ReqListIdUser) (RespListIdUser, error) {
	var resp RespListIdUser
	err := c.executeWXApiPost("/cgi-bin/user/list_id", req, &resp, true)
	if err != nil {
		return RespListIdUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListIdUser{}, bizErr
	}
	return resp, nil
}
