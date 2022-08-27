package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetCorpTagListExternalcontact 获取企业标签库请求
// 文档：https://developer.work.weixin.qq.com/document/path/92696#获取企业标签库
type ReqGetCorpTagListExternalcontact struct {
	// GroupID 要查询的标签组id，返回该标签组以及其下的所有标签信息
	GroupID []string `json:"group_id,omitempty"`
	// TagID 要查询的标签id
	TagID []string `json:"tag_id,omitempty"`
}

var _ urlValuer = ReqGetCorpTagListExternalcontact{}

func (x ReqGetCorpTagListExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}

	return ret
}

// RespGetCorpTagListExternalcontact 获取企业标签库响应
// 文档：https://developer.work.weixin.qq.com/document/path/92696#获取企业标签库
type RespGetCorpTagListExternalcontact struct {
	CommonResp
	TagGroup []struct {
		// CreateTime 标签组创建时间
		// CreateTime 标签创建时间
		// CreateTime 标签组创建时间
		// CreateTime 标签创建时间
		CreateTime int `json:"create_time"`
		// Deleted 标签组是否已经被删除，只在指定tag_id进行查询时返回
		// Deleted 标签是否已经被删除，只在指定tag_id/group_id进行查询时返回
		// Deleted 标签组是否已经被删除，只在指定tag_id进行查询时返回
		// Deleted 标签是否已经被删除，只在指定tag_id/group_id进行查询时返回
		Deleted bool `json:"deleted"`
		// GroupID 标签组id
		GroupID string `json:"group_id"`
		// GroupName 标签组名称
		GroupName string `json:"group_name"`
		// Order 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
		Order int `json:"order"`
		Tag   []struct {
			CreateTime int  `json:"create_time"`
			Deleted    bool `json:"deleted"`
			// ID 标签id
			ID string `json:"id"`
			// Name 标签名称
			Name  string `json:"name"`
			Order int    `json:"order"`
		} `json:"tag"` // 标签组内的标签列表
	} `json:"tag_group"` // 标签组列表
}

var _ bodyer = RespGetCorpTagListExternalcontact{}

func (x RespGetCorpTagListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetCorpTagListExternalcontact 获取企业标签库
// 文档：https://developer.work.weixin.qq.com/document/path/92696#获取企业标签库
func (c *ApiClient) ExecGetCorpTagListExternalcontact(req ReqGetCorpTagListExternalcontact) (RespGetCorpTagListExternalcontact, error) {
	var resp RespGetCorpTagListExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get_corp_tag_list", req, &resp, true)
	if err != nil {
		return RespGetCorpTagListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCorpTagListExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqAddCorpTagExternalcontact 添加企业客户标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/92696#添加企业客户标签

type ReqAddCorpTagExternalcontact struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid int `json:"agentid,omitempty"`
	// GroupID 标签组id
	GroupID string `json:"group_id,omitempty"`
	// GroupName 标签组名称，最长为30个<strong>字符</strong>
	GroupName string `json:"group_name,omitempty"`
	// Order 标签组次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty"`
	Tag   []struct {
		// Name 添加的标签名称，最长为30个<strong>字符</strong>，必填
		Name  string `json:"name"`
		Order int    `json:"order,omitempty"`
	} `json:"tag"`
}

var _ urlValuer = ReqAddCorpTagExternalcontact{}

func (x ReqAddCorpTagExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}

	return ret
}

// RespAddCorpTagExternalcontact 添加企业客户标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/92696#添加企业客户标签

type RespAddCorpTagExternalcontact struct {
	CommonResp
	TagGroup struct {
		// CreateTime 标签组创建时间
		// CreateTime 标签创建时间
		// CreateTime 标签组创建时间
		// CreateTime 标签创建时间
		CreateTime int `json:"create_time"`
		// GroupID 标签组id
		GroupID string `json:"group_id"`
		// GroupName 标签组名称
		GroupName string `json:"group_name"`
		// Order 标签组次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签组次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
		// Order 标签次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
		Order int `json:"order"`
		Tag   []struct {
			CreateTime int `json:"create_time"`
			// ID 新建标签id
			ID string `json:"id"`
			// Name 新建标签名称
			Name  string `json:"name"`
			Order int    `json:"order"`
		} `json:"tag"` // 标签组内的标签列表
	} `json:"tag_group"`
}

var _ bodyer = RespAddCorpTagExternalcontact{}

func (x RespAddCorpTagExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddCorpTagExternalcontact 添加企业客户标签
// 文档：https://developer.work.weixin.qq.com/document/path/92696#添加企业客户标签
func (c *ApiClient) ExecAddCorpTagExternalcontact(req ReqAddCorpTagExternalcontact) (RespAddCorpTagExternalcontact, error) {
	var resp RespAddCorpTagExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/add_corp_tag", req, &resp, true)
	if err != nil {
		return RespAddCorpTagExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddCorpTagExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqEditCorpTagExternalcontact 编辑企业客户标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/92696#编辑企业客户标签

type ReqEditCorpTagExternalcontact struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid int `json:"agentid,omitempty"`
	// ID 标签或标签组的id，必填
	ID string `json:"id"`
	// Name 新的标签或标签组名称，最长为30个<strong>字符</strong>
	Name string `json:"name,omitempty"`
	// Order 标签/标签组的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty"`
}

var _ urlValuer = ReqEditCorpTagExternalcontact{}

func (x ReqEditCorpTagExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}

	return ret
}

// RespEditCorpTagExternalcontact 编辑企业客户标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/92696#编辑企业客户标签

type RespEditCorpTagExternalcontact struct {
	CommonResp
}

var _ bodyer = RespEditCorpTagExternalcontact{}

func (x RespEditCorpTagExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecEditCorpTagExternalcontact 编辑企业客户标签
// 文档：https://developer.work.weixin.qq.com/document/path/92696#编辑企业客户标签
func (c *ApiClient) ExecEditCorpTagExternalcontact(req ReqEditCorpTagExternalcontact) (RespEditCorpTagExternalcontact, error) {
	var resp RespEditCorpTagExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/edit_corp_tag", req, &resp, true)
	if err != nil {
		return RespEditCorpTagExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespEditCorpTagExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqDelCorpTagExternalcontact 删除企业客户标签请求
// 文档：https://developer.work.weixin.qq.com/document/path/92696#删除企业客户标签

type ReqDelCorpTagExternalcontact struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid int `json:"agentid,omitempty"`
	// GroupID 标签组的id列表
	GroupID []string `json:"group_id,omitempty"`
	// TagID 标签的id列表
	TagID []string `json:"tag_id,omitempty"`
}

var _ urlValuer = ReqDelCorpTagExternalcontact{}

func (x ReqDelCorpTagExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDelCorpTagExternalcontact 删除企业客户标签响应
// 文档：https://developer.work.weixin.qq.com/document/path/92696#删除企业客户标签

type RespDelCorpTagExternalcontact struct {
	CommonResp
}

var _ bodyer = RespDelCorpTagExternalcontact{}

func (x RespDelCorpTagExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDelCorpTagExternalcontact 删除企业客户标签
// 文档：https://developer.work.weixin.qq.com/document/path/92696#删除企业客户标签
func (c *ApiClient) ExecDelCorpTagExternalcontact(req ReqDelCorpTagExternalcontact) (RespDelCorpTagExternalcontact, error) {
	var resp RespDelCorpTagExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/del_corp_tag", req, &resp, true)
	if err != nil {
		return RespDelCorpTagExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDelCorpTagExternalcontact{}, bizErr
	}
	return resp, nil
}
