package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListLinkCustomerAcquisition 获取获客链接列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接列表
type ReqListLinkCustomerAcquisition struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值100
	Limit int `json:"limit,omitempty"`
}

var _ urlValuer = ReqListLinkCustomerAcquisition{}

func (x ReqListLinkCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListLinkCustomerAcquisition 获取获客链接列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接列表
type RespListLinkCustomerAcquisition struct {
	CommonResp
	// LinkIDList link_id列表
	LinkIDList []string `json:"link_id_list"`
	// NextCursor 分页游标，在下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListLinkCustomerAcquisition{}

func (x RespListLinkCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListLinkCustomerAcquisition 获取获客链接列表
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接列表
func (c *ApiClient) ExecListLinkCustomerAcquisition(req ReqListLinkCustomerAcquisition) (RespListLinkCustomerAcquisition, error) {
	var resp RespListLinkCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/list_link", req, &resp, true)
	if err != nil {
		return RespListLinkCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListLinkCustomerAcquisition{}, bizErr
	}
	return resp, nil
}

// ReqGetCustomerAcquisition 获取获客链接详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接详情
type ReqGetCustomerAcquisition struct {
	// LinkID 获客链接id，必填
	LinkID string `json:"link_id"`
}

var _ urlValuer = ReqGetCustomerAcquisition{}

func (x ReqGetCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetCustomerAcquisition 获取获客链接详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接详情
type RespGetCustomerAcquisition struct {
	CommonResp
	Link struct {
		// CreateTime 创建时间
		CreateTime int `json:"create_time"`
		// LinkID 获客链接的id
		LinkID string `json:"link_id"`
		// LinkName 获客链接的名称
		LinkName string `json:"link_name"`
		// URL 获客链接实际的url
		URL string `json:"url"`
	} `json:"link"`
	Range struct {
		// DepartmentList 该获客链接使用范围的部门列表
		DepartmentList []int `json:"department_list"`
		// UserList 该获客链接使用范围成员列表
		UserList []string `json:"user_list"`
	} `json:"range"`
	// SkipVerify 是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify"`
}

var _ bodyer = RespGetCustomerAcquisition{}

func (x RespGetCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetCustomerAcquisition 获取获客链接详情
// 文档：https://developer.work.weixin.qq.com/document/path/97297#获取获客链接详情
func (c *ApiClient) ExecGetCustomerAcquisition(req ReqGetCustomerAcquisition) (RespGetCustomerAcquisition, error) {
	var resp RespGetCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/get", req, &resp, true)
	if err != nil {
		return RespGetCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCustomerAcquisition{}, bizErr
	}
	return resp, nil
}

// ReqCreateLinkCustomerAcquisition 创建获客链接请求
// 文档：https://developer.work.weixin.qq.com/document/path/97297#创建获客链接
type ReqCreateLinkCustomerAcquisition struct {
	// LinkName 链接名称，必填
	LinkName string `json:"link_name"`
	Range    struct {
		// DepartmentList 此获客链接关联的部门id列表，部门覆盖总人数最多100个
		DepartmentList []int `json:"department_list,omitempty"`
		// UserList 此获客链接关联的userid列表，最多可关联100个
		UserList []string `json:"user_list,omitempty"`
	} `json:"range"`
	// SkipVerify 是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify,omitempty"`
}

var _ urlValuer = ReqCreateLinkCustomerAcquisition{}

func (x ReqCreateLinkCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespCreateLinkCustomerAcquisition 创建获客链接响应
// 文档：https://developer.work.weixin.qq.com/document/path/97297#创建获客链接
type RespCreateLinkCustomerAcquisition struct {
	CommonResp
	Link struct {
		// CreateTime 获客链接创建时间
		CreateTime int `json:"create_time"`
		// LinkID 获客链接的id
		LinkID string `json:"link_id"`
		// LinkName 获客链接名称
		LinkName string `json:"link_name"`
		// URL 获客链接
		URL string `json:"url"`
	} `json:"link"`
}

var _ bodyer = RespCreateLinkCustomerAcquisition{}

func (x RespCreateLinkCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateLinkCustomerAcquisition 创建获客链接
// 文档：https://developer.work.weixin.qq.com/document/path/97297#创建获客链接
func (c *ApiClient) ExecCreateLinkCustomerAcquisition(req ReqCreateLinkCustomerAcquisition) (RespCreateLinkCustomerAcquisition, error) {
	var resp RespCreateLinkCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/create_link", req, &resp, true)
	if err != nil {
		return RespCreateLinkCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateLinkCustomerAcquisition{}, bizErr
	}
	return resp, nil
}

// ReqUpdateLinkCustomerAcquisition 编辑获客链接请求
// 文档：https://developer.work.weixin.qq.com/document/path/97297#编辑获客链接
type ReqUpdateLinkCustomerAcquisition struct {
	// LinkID 获客链接的id，必填
	LinkID string `json:"link_id"`
	// LinkName 更新的链接名称
	LinkName string `json:"link_name,omitempty"`
	Range    struct {
		// DepartmentList 此获客链接关联的部门id列表，部门覆盖总人数最多100个
		DepartmentList []int `json:"department_list,omitempty"`
		// UserList 此获客链接关联的userid列表，最多可关联100个
		UserList []string `json:"user_list,omitempty"`
	} `json:"range"`
	// SkipVerify 是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify,omitempty"`
}

var _ urlValuer = ReqUpdateLinkCustomerAcquisition{}

func (x ReqUpdateLinkCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespUpdateLinkCustomerAcquisition 编辑获客链接响应
// 文档：https://developer.work.weixin.qq.com/document/path/97297#编辑获客链接
type RespUpdateLinkCustomerAcquisition struct {
	CommonResp
}

var _ bodyer = RespUpdateLinkCustomerAcquisition{}

func (x RespUpdateLinkCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateLinkCustomerAcquisition 编辑获客链接
// 文档：https://developer.work.weixin.qq.com/document/path/97297#编辑获客链接
func (c *ApiClient) ExecUpdateLinkCustomerAcquisition(req ReqUpdateLinkCustomerAcquisition) (RespUpdateLinkCustomerAcquisition, error) {
	var resp RespUpdateLinkCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/update_link", req, &resp, true)
	if err != nil {
		return RespUpdateLinkCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateLinkCustomerAcquisition{}, bizErr
	}
	return resp, nil
}

// ReqDeleteLinkCustomerAcquisition 删除获客链接请求
// 文档：https://developer.work.weixin.qq.com/document/path/97297#删除获客链接
type ReqDeleteLinkCustomerAcquisition struct {
	// LinkID 获客链接的id，必填
	LinkID string `json:"link_id"`
}

var _ urlValuer = ReqDeleteLinkCustomerAcquisition{}

func (x ReqDeleteLinkCustomerAcquisition) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDeleteLinkCustomerAcquisition 删除获客链接响应
// 文档：https://developer.work.weixin.qq.com/document/path/97297#删除获客链接
type RespDeleteLinkCustomerAcquisition struct {
	CommonResp
}

var _ bodyer = RespDeleteLinkCustomerAcquisition{}

func (x RespDeleteLinkCustomerAcquisition) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeleteLinkCustomerAcquisition 删除获客链接
// 文档：https://developer.work.weixin.qq.com/document/path/97297#删除获客链接
func (c *ApiClient) ExecDeleteLinkCustomerAcquisition(req ReqDeleteLinkCustomerAcquisition) (RespDeleteLinkCustomerAcquisition, error) {
	var resp RespDeleteLinkCustomerAcquisition
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition/delete_link", req, &resp, true)
	if err != nil {
		return RespDeleteLinkCustomerAcquisition{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteLinkCustomerAcquisition{}, bizErr
	}
	return resp, nil
}
