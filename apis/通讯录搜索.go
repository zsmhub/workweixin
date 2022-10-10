package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSearchContact 通讯录单个搜索请求
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录单个搜索
type ReqSearchContact struct {
	// Agentid 应用id，若非0则只返回应用可见范围内的用户或者部门信息
	Agentid int `json:"agentid,omitempty"`
	// AuthCorpid 查询的企业corpid，必填
	AuthCorpid string `json:"auth_corpid"`
	// FullMatchField 精确匹配的字段。1:匹配用户名称或者部门名称 2:匹配用户英文名。不填则为模糊匹配
	FullMatchField int `json:"full_match_field,omitempty"`
	// Limit 查询返回的最大数量，默认为50，最多为200，查询返回的数量可能小于limit指定的值
	Limit int `json:"limit,omitempty"`
	// Offset 查询的偏移量，每次调用的offset在上一次offset基础上加上limit
	Offset int `json:"offset,omitempty"`
	// QueryType 查询类型 1:查询用户，返回用户userid列表 2:查询部门，返回部门id列表。 不填该字段或者填0代表同时查询部门跟用户
	QueryType int `json:"query_type,omitempty"`
	// QueryWord 搜索关键词。当查询用户时应为用户名称、名称拼音或者英文名；当查询部门时应为部门名称或者部门名称拼音，必填
	QueryWord string `json:"query_word"`
	// QueryRange 查询范围，仅查询类型包含用户时有效。 0：只查询在职用户 1：同时查询在职和离职用户
	QueryRange int `json:"query_range"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
}

var _ bodyer = ReqSearchContact{}

func (x ReqSearchContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSearchContact 通讯录单个搜索响应
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录单个搜索
type RespSearchContact struct {
	CommonResp
	// IsLast 根据该字段判断是否是最后一页，若为false，开发者需要使用offset+limit继续调用
	IsLast      bool `json:"is_last"`
	QueryResult struct {
		Party struct {
			// DepartmentID 返回的部门id
			DepartmentID []int `json:"department_id"`
		} `json:"party"` // 返回的部门信息 （通过部门名称，拼音匹配）
		User struct {
			// OpenUserid 查询到的用户open_userid
			OpenUserid []string `json:"open_userid"`
			// Userid 查询到的用户userid
			Userid []string `json:"userid"`
		} `json:"user"` // 返回的用户信息（通过用户名称，拼音匹配）
	} `json:"query_result"` // 查询结果
}

var _ bodyer = RespSearchContact{}

func (x RespSearchContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSearchContact 通讯录单个搜索
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录单个搜索
func (c *ApiClient) ExecSearchContact(req ReqSearchContact) (RespSearchContact, error) {
	var resp RespSearchContact
	err := c.executeWXApiPost("/cgi-bin/service/contact/search", req, &resp, true)
	if err != nil {
		return RespSearchContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSearchContact{}, bizErr
	}
	return resp, nil
}

// ReqBatchsearchContact 通讯录批量搜索请求
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录批量搜索
type ReqBatchsearchContact struct {
	// Agentid 应用id，若非0则只返回应用可见范围内的用户或者部门信息
	Agentid int `json:"agentid,omitempty"`
	// AuthCorpid 查询的企业corpid，必填
	AuthCorpid       string `json:"auth_corpid"`
	QueryRequestList []struct {
		// FullMatchField 如果需要精确匹配用户名称或者部门名称或者英文名，不填则默认为模糊匹配；1:匹配用户名称或者部门名称 2:匹配用户英文名
		FullMatchField int `json:"full_match_field,omitempty"`
		// Limit 查询返回的最大数量，默认为50，最多为200，查询返回的数量可能小于limit指定的值
		Limit int `json:"limit,omitempty"`
		// Offset 查询的偏移量，每次调用的offset在上一次offset基础上加上limit
		Offset int `json:"offset,omitempty"`
		// QueryType 查询类型 1:查询用户，返回用户userid列表 2:查询部门，返回部门id列表。 不填该字段或者填0代表同时查询部门跟用户
		QueryType int `json:"query_type,omitempty"`
		// QueryWord 搜索关键词。当查询用户时应为用户名称、名称拼音或者英文名；当查询部门时应为部门名称或者部门名称拼音，必填
		QueryWord string `json:"query_word"`
	} `json:"query_request_list"` // 搜索请求列表,每次搜索列表数量不超过50，必填
}

var _ bodyer = ReqBatchsearchContact{}

func (x ReqBatchsearchContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchsearchContact 通讯录批量搜索响应
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录批量搜索
type RespBatchsearchContact struct {
	CommonResp
	QueryResultList []struct {
		// IsLast 根据该字段判断是否是最后一页，若为false，开发者需要使用offset+limit继续调用
		IsLast       bool `json:"is_last"`
		QueryRequest struct {
			Limit     int    `json:"limit"`
			Offset    int    `json:"offset"`
			QueryType int    `json:"query_type"`
			QueryWord string `json:"query_word"`
		} `json:"query_request"` // 原搜索请求报文
		QueryResult struct {
			Party struct {
				// DepartmentID 返回的部门id
				DepartmentID []int `json:"department_id"`
			} `json:"party"` // 返回的部门信息 （通过部门名称，拼音匹配）
			User struct {
				// OpenUserid 查询到的用户open_userid
				OpenUserid []string `json:"open_userid"`
				// Userid 查询到的用户userid
				Userid []string `json:"userid"`
			} `json:"user"` // 返回的用户信息（通过用户名称，拼音匹配）
		} `json:"query_result"` // 搜索请求对应的查询结果
	} `json:"query_result_list"` // 搜索结果列表
}

var _ bodyer = RespBatchsearchContact{}

func (x RespBatchsearchContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchsearchContact 通讯录批量搜索
// 文档：https://developer.work.weixin.qq.com/document/path/91844#通讯录批量搜索
func (c *ApiClient) ExecBatchsearchContact(req ReqBatchsearchContact) (RespBatchsearchContact, error) {
	var resp RespBatchsearchContact
	err := c.executeWXApiPost("/cgi-bin/service/contact/batchsearch", req, &resp, true)
	if err != nil {
		return RespBatchsearchContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchsearchContact{}, bizErr
	}
	return resp, nil
}
