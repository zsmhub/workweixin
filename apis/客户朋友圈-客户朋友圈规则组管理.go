package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListMomentStrategy 获取规则组列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组列表
type ReqListMomentStrategy struct {
	// Cursor 分页查询游标，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 分页大小,默认为1000，最大不超过1000
	Limit int `json:"limit,omitempty"`
}

var _ urlValuer = ReqListMomentStrategy{}

func (x ReqListMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListMomentStrategy 获取规则组列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组列表
type RespListMomentStrategy struct {
	CommonResp
	// NextCursor 分页游标，用于查询下一个分页的数据，无更多数据时不返回
	NextCursor string `json:"next_cursor"`
	Strategy   []struct {
		// StrategyID 规则组id
		StrategyID int `json:"strategy_id"`
	} `json:"strategy"`
}

var _ bodyer = RespListMomentStrategy{}

func (x RespListMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListMomentStrategy 获取规则组列表
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组列表
func (c *ApiClient) ExecListMomentStrategy(req ReqListMomentStrategy) (RespListMomentStrategy, error) {
	var resp RespListMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/list", req, &resp, true)
	if err != nil {
		return RespListMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListMomentStrategy{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentStrategy 获取规则组详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组详情
type ReqGetMomentStrategy struct {
	// StrategyID 规则组id，必填
	StrategyID int `json:"strategy_id"`
}

var _ urlValuer = ReqGetMomentStrategy{}

func (x ReqGetMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetMomentStrategy 获取规则组详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组详情
type RespGetMomentStrategy struct {
	CommonResp
	Strategy struct {
		// AdminList 规则组管理员userid列表
		AdminList []string `json:"admin_list"`
		// CreateTime 规则组创建时间戳
		CreateTime int `json:"create_time"`
		// ParentID 父规则组id， 如果当前规则组没父规则组，则为0
		ParentID  int `json:"parent_id"`
		Privilege struct {
			// ManageMomentCoverAndSign 配置封面和签名，默认为true
			ManageMomentCoverAndSign bool `json:"manage_moment_cover_and_sign"`
			// SendMoment 允许成员发表客户朋友圈，默认为true
			SendMoment bool `json:"send_moment"`
			// ViewMomentList 允许查看成员的全部客户朋友圈发表
			ViewMomentList bool `json:"view_moment_list"`
		} `json:"privilege"`
		// StrategyID 规则组id
		StrategyID int `json:"strategy_id"`
		// StrategyName 规则组名称
		StrategyName string `json:"strategy_name"`
	} `json:"strategy"`
}

var _ bodyer = RespGetMomentStrategy{}

func (x RespGetMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentStrategy 获取规则组详情
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组详情
func (c *ApiClient) ExecGetMomentStrategy(req ReqGetMomentStrategy) (RespGetMomentStrategy, error) {
	var resp RespGetMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/get", req, &resp, true)
	if err != nil {
		return RespGetMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentStrategy{}, bizErr
	}
	return resp, nil
}

// ReqGetRangeMomentStrategy 获取规则组管理范围请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组管理范围
type ReqGetRangeMomentStrategy struct {
	// Cursor 分页游标
	Cursor string `json:"cursor,omitempty"`
	// Limit 每个分页的成员/部门节点数，默认为1000，最大为1000
	Limit int `json:"limit,omitempty"`
	// StrategyID 规则组id，必填
	StrategyID int `json:"strategy_id"`
}

var _ urlValuer = ReqGetRangeMomentStrategy{}

func (x ReqGetRangeMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetRangeMomentStrategy 获取规则组管理范围响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组管理范围
type RespGetRangeMomentStrategy struct {
	CommonResp
	// NextCursor 分页游标，用于查询下一个分页的数据，无更多数据时不返回
	NextCursor string `json:"next_cursor"`
	Range      []struct {
		// Partyid 管理范围内配置的部门partyid，仅<code>type</code>为2时返回
		Partyid int `json:"partyid"`
		// Type 节点类型，1-成员 2-部门
		Type int `json:"type"`
		// Userid 管理范围内配置的成员userid，仅<code>type</code>为1时返回
		Userid string `json:"userid"`
	} `json:"range"`
}

var _ bodyer = RespGetRangeMomentStrategy{}

func (x RespGetRangeMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetRangeMomentStrategy 获取规则组管理范围
// 文档：https://developer.work.weixin.qq.com/document/path/94890#获取规则组管理范围
func (c *ApiClient) ExecGetRangeMomentStrategy(req ReqGetRangeMomentStrategy) (RespGetRangeMomentStrategy, error) {
	var resp RespGetRangeMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/get_range", req, &resp, true)
	if err != nil {
		return RespGetRangeMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetRangeMomentStrategy{}, bizErr
	}
	return resp, nil
}

// ReqCreateMomentStrategy 创建新的规则组请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#创建新的规则组
type ReqCreateMomentStrategy struct {
	ParentID     int      `json:"parent_id"`
	StrategyName string   `json:"strategy_name"`
	AdminList    []string `json:"admin_list"`
	Privilege    struct {
		SendMoment               bool `json:"send_moment"`
		ViewMomentList           bool `json:"view_moment_list"`
		ManageMomentCoverAndSign bool `json:"manage_moment_cover_and_sign"`
	} `json:"privilege"`
	Range []struct {
		Type    int    `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int    `json:"partyid,omitempty"`
	} `json:"range"`
}

var _ urlValuer = ReqCreateMomentStrategy{}

func (x ReqCreateMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespCreateMomentStrategy 创建新的规则组响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#创建新的规则组
type RespCreateMomentStrategy struct {
	CommonResp
	StrategyID int `json:"strategy_id"`
}

var _ bodyer = RespCreateMomentStrategy{}

func (x RespCreateMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateMomentStrategy 创建新的规则组
// 文档：https://developer.work.weixin.qq.com/document/path/94890#创建新的规则组
func (c *ApiClient) ExecCreateMomentStrategy(req ReqCreateMomentStrategy) (RespCreateMomentStrategy, error) {
	var resp RespCreateMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/create", req, &resp, true)
	if err != nil {
		return RespCreateMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateMomentStrategy{}, bizErr
	}
	return resp, nil
}

// ReqEditMomentStrategy 编辑规则组及其管理范围请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#编辑规则组及其管理范围
type ReqEditMomentStrategy struct {
	// AdminList 管理员列表，如果为空则不对负责人做编辑，如果有则<strong>覆盖旧的负责人列表</strong>
	AdminList []string `json:"admin_list,omitempty"`
	Privilege struct {
		ManageMomentCoverAndSign bool `json:"manage_moment_cover_and_sign"`
		SendMoment               bool `json:"send_moment"`
		ViewMomentList           bool `json:"view_moment_list"`
	} `json:"privilege,omitempty"` // 权限配置，如果为空则不对权限做编辑，如果有则<strong>覆盖旧的权限配置</strong>
	RangeAdd []struct {
		// Partyid 向管理范围添加部门的partyid，仅type为2时有效
		Partyid int `json:"partyid,omitempty"`
		// Type 向管理范围添加的节点类型 1-成员 2-部门
		Type int `json:"type,omitempty"`
		// Userid 向管理范围添加成员的userid,仅type为1时有效
		Userid string `json:"userid,omitempty"`
	} `json:"range_add"`
	RangeDel []struct {
		// Partyid 从管理范围删除的部门的partyid，仅type为2时有效
		Partyid int `json:"partyid,omitempty"`
		// Type 从管理范围删除的节点类型 1-成员 2-部门
		Type int `json:"type,omitempty"`
		// Userid 从管理范围删除的成员的userid,仅type为1时有效
		Userid string `json:"userid,omitempty"`
	} `json:"range_del"`
	// StrategyID 规则组id，必填
	StrategyID int `json:"strategy_id"`
	// StrategyName 规则组名称
	StrategyName string `json:"strategy_name,omitempty"`
}

var _ urlValuer = ReqEditMomentStrategy{}

func (x ReqEditMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespEditMomentStrategy 编辑规则组及其管理范围响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#编辑规则组及其管理范围
type RespEditMomentStrategy struct {
	CommonResp
}

var _ bodyer = RespEditMomentStrategy{}

func (x RespEditMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecEditMomentStrategy 编辑规则组及其管理范围
// 文档：https://developer.work.weixin.qq.com/document/path/94890#编辑规则组及其管理范围
func (c *ApiClient) ExecEditMomentStrategy(req ReqEditMomentStrategy) (RespEditMomentStrategy, error) {
	var resp RespEditMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/edit", req, &resp, true)
	if err != nil {
		return RespEditMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespEditMomentStrategy{}, bizErr
	}
	return resp, nil
}

// ReqDelMomentStrategy 删除规则组请求
// 文档：https://developer.work.weixin.qq.com/document/path/94890#删除规则组
type ReqDelMomentStrategy struct {
	// StrategyID 规则组id，必填
	StrategyID int `json:"strategy_id"`
}

var _ urlValuer = ReqDelMomentStrategy{}

func (x ReqDelMomentStrategy) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDelMomentStrategy 删除规则组响应
// 文档：https://developer.work.weixin.qq.com/document/path/94890#删除规则组
type RespDelMomentStrategy struct {
	CommonResp
}

var _ bodyer = RespDelMomentStrategy{}

func (x RespDelMomentStrategy) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDelMomentStrategy 删除规则组
// 文档：https://developer.work.weixin.qq.com/document/path/94890#删除规则组
func (c *ApiClient) ExecDelMomentStrategy(req ReqDelMomentStrategy) (RespDelMomentStrategy, error) {
	var resp RespDelMomentStrategy
	err := c.executeWXApiGet("/cgi-bin/externalcontact/moment_strategy/del", req, &resp, true)
	if err != nil {
		return RespDelMomentStrategy{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDelMomentStrategy{}, bizErr
	}
	return resp, nil
}
