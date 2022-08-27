package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchShareActiveCodeLicense 分配激活码给下游企业请求
// 文档：https://developer.work.weixin.qq.com/document/path/96059#分配激活码给下游企业
type ReqBatchShareActiveCodeLicense struct {
	// FromCorpid 上游企业corpid。支持明文或者密文的corpid，必填
	FromCorpid string `json:"from_corpid"`
	ShareList  []struct {
		// ActiveCode 分享的激活码，必填
		ActiveCode string `json:"active_code"`
	} `json:"share_list"` // 分享的接口许可列表。单次分享激活码不可超过1000个，累计分享给同一下游企业的激活码总数不可超过上下游通讯录中该下游企业人数的2倍，必填
	// ToCorpid 下游企业corpid。支持明文或者密文的corpid，必填
	ToCorpid string `json:"to_corpid"`
}

var _ bodyer = ReqBatchShareActiveCodeLicense{}

func (x ReqBatchShareActiveCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchShareActiveCodeLicense 分配激活码给下游企业响应
// 文档：https://developer.work.weixin.qq.com/document/path/96059#分配激活码给下游企业
type RespBatchShareActiveCodeLicense struct {
	// Errcode 错误码说明
	CommonResp
	Errmsg      string `json:"errmsg"`
	ShareResult []struct {
		// ActiveCode 分享的激活码
		ActiveCode string `json:"active_code"`
		// Errcode 错误码说明
		CommonResp
		Errmsg string `json:"errmsg"`
	} `json:"share_result"` // 分享结果
}

var _ bodyer = RespBatchShareActiveCodeLicense{}

func (x RespBatchShareActiveCodeLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchShareActiveCodeLicense 分配激活码给下游企业
// 文档：https://developer.work.weixin.qq.com/document/path/96059#分配激活码给下游企业
func (c *ApiClient) ExecBatchShareActiveCodeLicense(req ReqBatchShareActiveCodeLicense) (RespBatchShareActiveCodeLicense, error) {
	var resp RespBatchShareActiveCodeLicense
	err := c.executeWXApiPost("/cgi-bin/license/batch_share_active_code", req, &resp, true)
	if err != nil {
		return RespBatchShareActiveCodeLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchShareActiveCodeLicense{}, bizErr
	}
	return resp, nil
}
