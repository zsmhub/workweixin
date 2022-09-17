package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqActiveAccountLicense 激活帐号请求
// 文档：https://developer.work.weixin.qq.com/document/path/95553#激活帐号
type ReqActiveAccountLicense struct {
	// ActiveCode 帐号激活码，必填
	ActiveCode string `json:"active_code"`
	// Corpid 待绑定激活的成员所属企业corpid，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
	// Userid 待绑定激活的企业成员userid 。只支持加密的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqActiveAccountLicense{}

func (x ReqActiveAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespActiveAccountLicense 激活帐号响应
// 文档：https://developer.work.weixin.qq.com/document/path/95553#激活帐号
type RespActiveAccountLicense struct {
	CommonResp
}

var _ bodyer = RespActiveAccountLicense{}

func (x RespActiveAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecActiveAccountLicense 激活帐号
// 文档：https://developer.work.weixin.qq.com/document/path/95553#激活帐号
func (c *ApiClient) ExecActiveAccountLicense(req ReqActiveAccountLicense) (RespActiveAccountLicense, error) {
	var resp RespActiveAccountLicense
	err := c.executeWXApiPost("/cgi-bin/license/active_account", req, &resp, true)
	if err != nil {
		return RespActiveAccountLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespActiveAccountLicense{}, bizErr
	}
	return resp, nil
}

// ReqBatchActiveAccountLicense 批量激活帐号请求
// 文档：https://developer.work.weixin.qq.com/document/path/95553#批量激活帐号
type ReqBatchActiveAccountLicense struct {
	ActiveList []struct {
		// ActiveCode 帐号激活码，必填
		// ActiveCode 帐号激活码，必填
		ActiveCode string `json:"active_code"`
		// Userid 待绑定激活的企业成员userid 。只支持加密的userid，必填
		// Userid 本次激活的企业成员的加密userid，必填
		Userid string `json:"userid"`
	} `json:"active_list"` // 需要激活的帐号列表，必填
	// Corpid 待绑定激活的成员所属企业corpid，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqBatchActiveAccountLicense{}

func (x ReqBatchActiveAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchActiveAccountLicense 批量激活帐号响应
// 文档：https://developer.work.weixin.qq.com/document/path/95553#批量激活帐号
type RespBatchActiveAccountLicense struct {
	ActiveResult []struct {
		ActiveCode string `json:"active_code"`
		Errcode    int    `json:"errcode"`
		Userid     string `json:"userid"`
	} `json:"active_result"`
	CommonResp
}

var _ bodyer = RespBatchActiveAccountLicense{}

func (x RespBatchActiveAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchActiveAccountLicense 批量激活帐号
// 文档：https://developer.work.weixin.qq.com/document/path/95553#批量激活帐号
func (c *ApiClient) ExecBatchActiveAccountLicense(req ReqBatchActiveAccountLicense) (RespBatchActiveAccountLicense, error) {
	var resp RespBatchActiveAccountLicense
	err := c.executeWXApiPost("/cgi-bin/license/batch_active_account", req, &resp, true)
	if err != nil {
		return resp, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return resp, bizErr
	}
	return resp, nil
}
