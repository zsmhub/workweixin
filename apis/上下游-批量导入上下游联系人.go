package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqImportChainContactCorpgroup 批量导入上下游联系人请求
// 文档：https://developer.work.weixin.qq.com/document/path/95821
type ReqImportChainContactCorpgroup struct {
	// ChainID 上下游id。文件中的联系人将会被导入此上下游中，必填
	ChainID     string `json:"chain_id"`
	ContactList []struct {
		ContactInfoList []struct {
			// IdentityType 联系人身份类型。1:成员，2:负责人。，必填
			IdentityType int `json:"identity_type"`
			// Mobile 手机号。支持国内、国际手机号（国内手机号直接输入手机号即可，格式示例:“138****0001”；国际手机号必须包含加号以及国家地区码，格式示例:“+85259****45”，必填
			Mobile string `json:"mobile"`
			// Name 上下游联系人姓名。长度为1～32个utf8字符，必填
			Name string `json:"name"`
			// UserCustomID 上下游用户自定义 id。类型为字符串，暂时只支持传入64比特无符号整型，取值范围1到2^64-2，必须是全数字，不得传入前置0，且不能为11位或13位数字。
			UserCustomID string `json:"user_custom_id,omitempty"`
		} `json:"contact_info_list"` // 上下游联系人信息列表，必填
		// CorpName 上下游企业名称。长度为1-32个utf8字符。只能由中文、字母、数字和“ -_()（）”六种字符组成，必填
		CorpName string `json:"corp_name"`
		// CustomID 上下游企业自定义 id。长度为0～64 个字节，只能由数字和字母组成
		CustomID string `json:"custom_id,omitempty"`
		// GroupPath 导入后企业所在分组。分组为空的企业会放在根分组下。仅针对新导入企业生效，不会修改已导入企业的分组。
		GroupPath string `json:"group_path,omitempty"`
	} `json:"contact_list"` // 上下游联系人列表。这些联系人将会被导入此上下游中，必填
}

var _ bodyer = ReqImportChainContactCorpgroup{}

func (x ReqImportChainContactCorpgroup) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespImportChainContactCorpgroup 批量导入上下游联系人响应
// 文档：https://developer.work.weixin.qq.com/document/path/95821
type RespImportChainContactCorpgroup struct {
	CommonResp
	// Jobid 异步任务id，最大长度为64字节
	Jobid string `json:"jobid"`
}

var _ bodyer = RespImportChainContactCorpgroup{}

func (x RespImportChainContactCorpgroup) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecImportChainContactCorpgroup 批量导入上下游联系人
// 文档：https://developer.work.weixin.qq.com/document/path/95821
func (c *ApiClient) ExecImportChainContactCorpgroup(req ReqImportChainContactCorpgroup) (RespImportChainContactCorpgroup, error) {
	var resp RespImportChainContactCorpgroup
	err := c.executeWXApiPost("/cgi-bin/corpgroup/import_chain_contact", req, &resp, true)
	if err != nil {
		return RespImportChainContactCorpgroup{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespImportChainContactCorpgroup{}, bizErr
	}
	return resp, nil
}
