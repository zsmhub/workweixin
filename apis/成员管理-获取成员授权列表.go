package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListMemberAuthUser 获取成员授权列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/94513#获取成员授权列表
type ReqListMemberAuthUser struct {
	// Cursor 上一次调用时返回的next_cursor，第一次拉取可以不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 每次拉取的数据量，默认值和最大值都为1000
	Limit int `json:"limit,omitempty"`
}

var _ bodyer = ReqListMemberAuthUser{}

func (x ReqListMemberAuthUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListMemberAuthUser 获取成员授权列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/94513#获取成员授权列表
type RespListMemberAuthUser struct {
	CommonResp
	MemberAuthList []struct {
		// OpenUserid 当用户为某个企业内部成员时返回。全局唯一，对于同一个服务商不同的应用open_userid是相同的；同一用户，对于不同服务商open_userid是不同的
		OpenUserid string `json:"open_userid"`
	} `json:"member_auth_list"` // 成员授权列表
	// NextCursor 当前数据最后一个key值，如果下次调用带上该值则从该key值往后拉，用于实现分页拉取，返回空字符串代表已经是最后一页
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListMemberAuthUser{}

func (x RespListMemberAuthUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListMemberAuthUser 获取成员授权列表
// 文档：https://developer.work.weixin.qq.com/document/path/94513#获取成员授权列表
func (c *ApiClient) ExecListMemberAuthUser(req ReqListMemberAuthUser) (RespListMemberAuthUser, error) {
	var resp RespListMemberAuthUser
	err := c.executeWXApiPost("/cgi-bin/user/list_member_auth", req, &resp, true)
	if err != nil {
		return RespListMemberAuthUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListMemberAuthUser{}, bizErr
	}
	return resp, nil
}
