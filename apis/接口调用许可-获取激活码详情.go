package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetActiveInfoByCodeLicense 获取激活码详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/95552#获取激活码详情
type ReqGetActiveInfoByCodeLicense struct {
	// ActiveCode 激活码，必填
	ActiveCode string `json:"active_code"`
	// Corpid 要查询的企业的corpid，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqGetActiveInfoByCodeLicense{}

func (x ReqGetActiveInfoByCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetActiveInfoByCodeLicense 获取激活码详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/95552#获取激活码详情
type RespGetActiveInfoByCodeLicense struct {
	ActiveInfo struct {
		// ActiveCode 帐号激活码
		ActiveCode string `json:"active_code"`
		// ActiveTime 首次激活绑定用户的时间，未激活则不返回该字段
		ActiveTime int `json:"active_time"`
		// CreateTime 创建时间，订单支付成功后立即创建。激活码必须在创建时间后的365天内激活。
		CreateTime int `json:"create_time"`
		// ExpireTime 过期时间。为首次激活绑定的时间加上购买时长。未激活则不返回该字段
		ExpireTime int `json:"expire_time"`
		MergeInfo  struct {
			// FromActiveCode 激活码激活userid时，若userid原来已经绑定了一个激活码，则会返回该字段
			FromActiveCode string `json:"from_active_code"`
			// ToActiveCode 该激活码合并到的新激活码信息
			ToActiveCode string `json:"to_active_code"`
		} `json:"merge_info"` // 合并信息，合并的激活码或者被合并的激活码才返回该字段。功能灰度中，目前仅通过api新续期过的才一定返回
		// Status 帐号状态:1:未绑定，2:已绑定且有效，3:已过期，4:待转移，5:已合并（激活码本身激活了userid，后续使用新的激活码重新激活了该userid，则该码变为已合并状态。若被合并时，该激活码未过期则合并后会重置expire_time为合并时间。若被合并时，激活码已过期则不重置expire_time）
		Status int `json:"status"`
		// Type 帐号类型:1:基础帐号，2:互通帐号
		Type int `json:"type"`
		// Userid 帐号绑定激活的企业成员userid，未激活则不返回该字段。返回加密的userid
		Userid string `json:"userid"`
	} `json:"active_info"` // 帐号码信息
	CommonResp
}

var _ bodyer = RespGetActiveInfoByCodeLicense{}

func (x RespGetActiveInfoByCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetActiveInfoByCodeLicense 获取激活码详情
// 文档：https://developer.work.weixin.qq.com/document/path/95552#获取激活码详情
func (c *ApiClient) ExecGetActiveInfoByCodeLicense(req ReqGetActiveInfoByCodeLicense) (RespGetActiveInfoByCodeLicense, error) {
	var resp RespGetActiveInfoByCodeLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_active_info_by_code", req, &resp, true)
	if err != nil {
		return RespGetActiveInfoByCodeLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetActiveInfoByCodeLicense{}, bizErr
	}
	return resp, nil
}

// ReqBatchGetActiveInfoByCodeLicense 批量获取激活码详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/95552#批量获取激活码详情
type ReqBatchGetActiveInfoByCodeLicense struct {
	// ActiveCodeList 激活码列表，最多不超过1000个，必填
	ActiveCodeList []string `json:"active_code_list"`
	// Corpid 要查询的企业的corpid，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqBatchGetActiveInfoByCodeLicense{}

func (x ReqBatchGetActiveInfoByCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchGetActiveInfoByCodeLicense 批量获取激活码详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/95552#批量获取激活码详情
type RespBatchGetActiveInfoByCodeLicense struct {
	ActiveInfoList []struct {
		// ActiveCode 帐号激活码
		ActiveCode string `json:"active_code"`
		// ActiveTime 首次激活绑定用户的时间，未激活则不返回该字段
		ActiveTime int `json:"active_time"`
		// CreateTime 创建时间，订单支付成功后立即创建。激活码必须在创建时间后的365天内激活
		CreateTime int `json:"create_time"`
		// ExpireTime 过期时间。为首次激活绑定的时间加上购买时长。未激活则不返回该字段
		ExpireTime int `json:"expire_time"`
		MergeInfo  struct {
			// FromActiveCode 合并信息，合并的激活码或者被合并的激活码才返回该字段。功能灰度中，目前仅通过api新续期过的才一定返回
			FromActiveCode string `json:"from_active_code"`
			// ToActiveCode 该激活码合并到的新激活码信息
			ToActiveCode string `json:"to_active_code"`
		} `json:"merge_info"` // 合并信息。功能灰度中，目前仅通过api新续期过的才一定返回
		// Status 帐号状态:1:未绑定，2:已绑定且有效，3:已过期，4:待转移，5:已合并（激活码本身激活了userid，后续使用新的激活码重新激活了该userid，则该码变为已合并状态。若被合并时，该激活码未过期则合并后会重置expire_time为合并时间。若被合并时，激活码已过期则不重置expire_time）
		Status int `json:"status"`
		// Type 帐号类型:1:基础帐号，2:互通帐号
		Type int `json:"type"`
		// Userid 帐号绑定激活的企业成员userid，未激活则不返回该字段。返回加密的userid
		Userid string `json:"userid"`
	} `json:"active_info_list"` // 帐号码信息列表
	CommonResp
	InvalidActiveCodeList []string `json:"invalid_active_code_list"`
}

var _ bodyer = RespBatchGetActiveInfoByCodeLicense{}

func (x RespBatchGetActiveInfoByCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchGetActiveInfoByCodeLicense 批量获取激活码详情
// 文档：https://developer.work.weixin.qq.com/document/path/95552#批量获取激活码详情
func (c *ApiClient) ExecBatchGetActiveInfoByCodeLicense(req ReqBatchGetActiveInfoByCodeLicense) (RespBatchGetActiveInfoByCodeLicense, error) {
	var resp RespBatchGetActiveInfoByCodeLicense
	err := c.executeWXApiPost("/cgi-bin/license/batch_get_active_info_by_code", req, &resp, true)
	if err != nil {
		return RespBatchGetActiveInfoByCodeLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchGetActiveInfoByCodeLicense{}, bizErr
	}
	return resp, nil
}
