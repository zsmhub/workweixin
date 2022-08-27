package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetUnassignedListExternalcontact 获取待分配的离职成员列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/92273#获取待分配的离职成员列表
type ReqGetUnassignedListExternalcontact struct {
	// Cursor 分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
	Cursor string `json:"cursor"`
	// PageID 分页查询，要查询页号，从0开始
	PageID int `json:"page_id"`
	// PageSize 每次返回的最大记录数，默认为1000，最大值为1000
	PageSize int `json:"page_size"`
}

var _ bodyer = ReqGetUnassignedListExternalcontact{}

func (x ReqGetUnassignedListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetUnassignedListExternalcontact 获取待分配的离职成员列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/92273#获取待分配的离职成员列表
type RespGetUnassignedListExternalcontact struct {
	CommonResp
	Info []struct {
		// DimissionTime 成员离职时间
		DimissionTime int `json:"dimission_time"`
		// ExternalUserid 外部联系人userid
		ExternalUserid string `json:"external_userid"`
		// HandoverUserid 离职成员的userid
		HandoverUserid string `json:"handover_userid"`
	} `json:"info"`
	// IsLast 是否是最后一条记录
	IsLast bool `json:"is_last"`
	// NextCursor 分页查询游标,已经查完则返回空(&#34;&#34;)，使用<code>page_id</code>作为查询参数时不返回
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetUnassignedListExternalcontact{}

func (x RespGetUnassignedListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUnassignedListExternalcontact 离职继承-获取待分配的离职成员列表
// 文档：https://developer.work.weixin.qq.com/document/path/92273#获取待分配的离职成员列表
func (c *ApiClient) ExecGetUnassignedListExternalcontact(req ReqGetUnassignedListExternalcontact) (RespGetUnassignedListExternalcontact, error) {
	var resp RespGetUnassignedListExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_unassigned_list", req, &resp, true)
	if err != nil {
		return RespGetUnassignedListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUnassignedListExternalcontact{}, bizErr
	}
	return resp, nil
}
