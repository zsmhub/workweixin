package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSortContact 通讯录userid排序请求
// 文档：https://developer.work.weixin.qq.com/document/path/92093#通讯录userid排序
type ReqSortContact struct {
	// AuthCorpid 查询的企业corpid，必填
	AuthCorpid  string `json:"auth_corpid"`
	SortOptions []struct {
		// SortField 排序属性。0:  根据姓名拼音排序，1: 根据别名排序
		SortField int `json:"sort_field,omitempty"`
		// SortType 排序方式。0:  升序排列，1: 降序排列
		SortType int `json:"sort_type,omitempty"`
	} `json:"sort_options,omitempty"` // 排序选项列表。如果指定多个，将根据优先级排序。如示例参数，先根据姓名拼音降序排列，如果姓名拼音相同，再根据别名升序排列
	// Useridlist 要排序的userid列表，最多支持1000个，必填
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqSortContact{}

func (x ReqSortContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSortContact 通讯录userid排序响应
// 文档：https://developer.work.weixin.qq.com/document/path/92093#通讯录userid排序
type RespSortContact struct {
	CommonResp
	// Useridlist 排序后的userid列表
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = RespSortContact{}

func (x RespSortContact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSortContact 通讯录userid排序
// 文档：https://developer.work.weixin.qq.com/document/path/92093#通讯录userid排序
func (c *ApiClient) ExecSortContact(req ReqSortContact) (RespSortContact, error) {
	var resp RespSortContact
	err := c.executeWXApiPost("/cgi-bin/service/contact/sort", req, &resp, true)
	if err != nil {
		return RespSortContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSortContact{}, bizErr
	}
	return resp, nil
}
