package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetPaymentInfoExternalpay 获取收款项目的商户单号请求
// 文档：https://developer.work.weixin.qq.com/document/path/95936#获取收款项目的商户单号
type ReqGetPaymentInfoExternalpay struct {
	// PaymentID 收款项目单号。在<a href="#39955" rel="nofollow">发起对外收款</a>中返回。，必填
	PaymentID string `json:"payment_id"`
}

var _ bodyer = ReqGetPaymentInfoExternalpay{}

func (x ReqGetPaymentInfoExternalpay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetPaymentInfoExternalpay 获取收款项目的商户单号响应
// 文档：https://developer.work.weixin.qq.com/document/path/95936#获取收款项目的商户单号
type RespGetPaymentInfoExternalpay struct {
	BillList []struct {
		OutTradeNo string `json:"out_trade_no"`
	} `json:"bill_list"`
	CommonResp
}

var _ bodyer = RespGetPaymentInfoExternalpay{}

func (x RespGetPaymentInfoExternalpay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetPaymentInfoExternalpay 获取收款项目的商户单号
// 文档：https://developer.work.weixin.qq.com/document/path/95936#获取收款项目的商户单号
func (c *ApiClient) ExecGetPaymentInfoExternalpay(req ReqGetPaymentInfoExternalpay) (RespGetPaymentInfoExternalpay, error) {
	var resp RespGetPaymentInfoExternalpay
	err := c.executeWXApiPost("/cgi-bin/externalpay/get_payment_info", req, &resp, true)
	if err != nil {
		return RespGetPaymentInfoExternalpay{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetPaymentInfoExternalpay{}, bizErr
	}
	return resp, nil
}
