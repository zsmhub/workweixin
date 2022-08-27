package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAgent 获取指定的应用详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取指定的应用详情
type ReqGetAgent struct {
	// Agentid 应用id，必填
	Agentid int `json:"agentid"`
}

var _ urlValuer = ReqGetAgent{}

func (x ReqGetAgent) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetAgent 获取指定的应用详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取指定的应用详情
type RespGetAgent struct {
	// Agentid 企业应用id
	Agentid     int `json:"agentid"`
	AllowPartys struct {
		Partyid []int `json:"partyid"`
	} `json:"allow_partys"` // 企业应用可见范围（部门）
	AllowTags struct {
		Tagid []int `json:"tagid"`
	} `json:"allow_tags"` // 企业应用可见范围（标签）
	AllowUserinfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"` // 企业应用可见范围（人员），其中包括userid
	// Close 企业应用是否被停用
	Close int `json:"close"`
	// CustomizedPublishStatus 代开发自建应用返回该字段，表示代开发发布状态。0:待开发（企业已授权，服务商未创建应用）；1:开发中（服务商已创建应用，未上线）；2:已上线（服务商已上线应用且不存在未上线版本）；3:存在未上线版本（服务商已上线应用但存在未上线版本）
	CustomizedPublishStatus int `json:"customized_publish_status"`
	// Description 企业应用详情
	Description string `json:"description"`
	CommonResp
	// HomeURL 应用主页url
	HomeURL string `json:"home_url"`
	// Isreportenter 是否上报用户进入应用事件。0:不接收；1:接收
	Isreportenter int `json:"isreportenter"`
	// Name 企业应用名称
	Name string `json:"name"`
	// RedirectDomain 企业应用可信域名
	RedirectDomain string `json:"redirect_domain"`
	// ReportLocationFlag 企业应用是否打开地理位置上报 0:不上报；1:进入会话上报；
	ReportLocationFlag int `json:"report_location_flag"`
	// SquareLogoURL 企业应用方形头像
	SquareLogoURL string `json:"square_logo_url"`
}

var _ bodyer = RespGetAgent{}

func (x RespGetAgent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAgent 获取指定的应用详情
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取指定的应用详情
func (c *ApiClient) ExecGetAgent(req ReqGetAgent) (RespGetAgent, error) {
	var resp RespGetAgent
	err := c.executeWXApiGet("/cgi-bin/agent/get", req, &resp, true)
	if err != nil {
		return RespGetAgent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAgent{}, bizErr
	}
	return resp, nil
}

// ReqListAgent 获取access_token对应的应用列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取access_token对应的应用列表
type ReqListAgent struct{}

var _ urlValuer = ReqListAgent{}

func (x ReqListAgent) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListAgent 获取access_token对应的应用列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取access_token对应的应用列表
type RespListAgent struct {
	Agentlist []struct {
		Agentid       int    `json:"agentid"`
		Name          string `json:"name"`
		SquareLogoURL string `json:"square_logo_url"`
	} `json:"agentlist"`
	CommonResp
}

var _ bodyer = RespListAgent{}

func (x RespListAgent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListAgent 获取access_token对应的应用列表
// 文档：https://developer.work.weixin.qq.com/document/path/90363#获取access_token对应的应用列表
func (c *ApiClient) ExecListAgent(req ReqListAgent) (RespListAgent, error) {
	var resp RespListAgent
	err := c.executeWXApiGet("/cgi-bin/agent/list", req, &resp, true)
	if err != nil {
		return RespListAgent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListAgent{}, bizErr
	}
	return resp, nil
}
