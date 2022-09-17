package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchTransferLicenseLicense 帐号继承请求
// 文档：https://developer.work.weixin.qq.com/document/path/95673#帐号继承
type (
	ReqBatchTransferLicenseLicense struct {
		// Corpid 待绑定激活的成员所属企业corpid，只支持加密的corpid，必填
		Corpid       string             `json:"corpid"`
		TransferList []TransferListItem `json:"transfer_list"`
	}

	TransferListItem struct {
		// HandoverUserid 转移成员的userid。只支持加密的userid，必填
		HandoverUserid string `json:"handover_userid"`
		// TakeoverUserid 接收成员的userid。只支持加密的userid，必填
		TakeoverUserid string `json:"takeover_userid"`
	}
)

var _ bodyer = ReqBatchTransferLicenseLicense{}

func (x ReqBatchTransferLicenseLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchTransferLicenseLicense 帐号继承响应
// 文档：https://developer.work.weixin.qq.com/document/path/95673#帐号继承
type RespBatchTransferLicenseLicense struct {
	CommonResp
	TransferResult []struct {
		Errcode int `json:"errcode"`
		// HandoverUserid 转移成员加密的userid
		HandoverUserid string `json:"handover_userid"`
		// TakeoverUserid 接收成员加密的userid
		TakeoverUserid string `json:"takeover_userid"`
	} `json:"transfer_result"`
}

var _ bodyer = RespBatchTransferLicenseLicense{}

func (x RespBatchTransferLicenseLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchTransferLicenseLicense 帐号继承
// 文档：https://developer.work.weixin.qq.com/document/path/95673#帐号继承
func (c *ApiClient) ExecBatchTransferLicenseLicense(req ReqBatchTransferLicenseLicense) (RespBatchTransferLicenseLicense, error) {
	var resp RespBatchTransferLicenseLicense
	err := c.executeWXApiPost("/cgi-bin/license/batch_transfer_license", req, &resp, true)
	if err != nil {
		return resp, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return resp, bizErr
	}
	return resp, nil
}
