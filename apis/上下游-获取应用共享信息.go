package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListAppShareInfoCorp 获取应用共享信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/95324#获取应用共享信息
type ReqListAppShareInfoCorp struct {
	// Agentid 上级/上游企业应用agentid，必填
	Agentid int `json:"agentid"`
	// BusinessType 填0则为企业互联/局校互联，填1则表示上下游企业
	BusinessType int `json:"business_type,omitempty"`
	// Corpid 下级/下游企业corpid，若指定该参数则表示拉取该下级/下游企业的应用共享信息
	Corpid string `json:"corpid,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值100，默认情况或者值为0表示下拉取全量数据，建议分页拉取或者通过指定corpid参数拉取。
	Limit int `json:"limit,omitempty"`
}

var _ bodyer = ReqListAppShareInfoCorp{}

func (x ReqListAppShareInfoCorp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListAppShareInfoCorp 获取应用共享信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/95324#获取应用共享信息
type RespListAppShareInfoCorp struct {
	CorpList []struct {
		// Agentid 下级/下游企业应用id
		Agentid int `json:"agentid"`
		// CorpName 下级/下游企业名称
		CorpName string `json:"corp_name"`
		// Corpid 下级/下游企业corpid
		Corpid string `json:"corpid"`
	} `json:"corp_list"` // 应用共享信息
	// Ending 1表示拉取完毕，0表示数据没有拉取完
	Ending int `json:"ending"`
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListAppShareInfoCorp{}

func (x RespListAppShareInfoCorp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListAppShareInfoCorp 获取应用共享信息
// 文档：https://developer.work.weixin.qq.com/document/path/95324#获取应用共享信息
func (c *ApiClient) ExecListAppShareInfoCorp(req ReqListAppShareInfoCorp) (RespListAppShareInfoCorp, error) {
	var resp RespListAppShareInfoCorp
	err := c.executeWXApiPost("/cgi-bin/corpgroup/corp/list_app_share_info", req, &resp, true)
	if err != nil {
		return RespListAppShareInfoCorp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListAppShareInfoCorp{}, bizErr
	}
	return resp, nil
}
