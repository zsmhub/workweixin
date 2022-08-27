package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetBillListExternalpay 获取对外收款记录请求
// 文档：https://developer.work.weixin.qq.com/document/path/93727#获取对外收款记录
type ReqGetBillListExternalpay struct {
	// BeginTime 收款记录开始时间，必填
	BeginTime int `json:"begin_time"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// EndTime 收款记录结束时间，必填
	EndTime int `json:"end_time"`
	// Limit 返回的最大记录数，整型，最大值1000
	Limit int `json:"limit,omitempty"`
	// PayeeUserid 企业收款成员userid，不填则为全部成员
	PayeeUserid string `json:"payee_userid,omitempty"`
}

var _ bodyer = ReqGetBillListExternalpay{}

func (x ReqGetBillListExternalpay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetBillListExternalpay 获取对外收款记录响应
// 文档：https://developer.work.weixin.qq.com/document/path/93727#获取对外收款记录
type RespGetBillListExternalpay struct {
	BillList []struct {
		CommodityList []struct {
			// Amount 商品数量
			Amount int `json:"amount"`
			// Description 商品描述
			Description string `json:"description"`
		} `json:"commodity_list"` // 商品信息详情列表,商品信息结构参考commodity
		// ExternalUserid 付款人的userid
		ExternalUserid string `json:"external_userid"`
		// MchID 收款商户号id
		MchID string `json:"mch_id"`
		// OutTradeNo 商户单号
		OutTradeNo string `json:"out_trade_no"`
		// PayTime 交易时间
		PayTime int `json:"pay_time"`
		// PayeeUserid 收款人企业内账号userid
		PayeeUserid string `json:"payee_userid"`
		PayerInfo   struct {
			// Address 联系地址
			Address string `json:"address"`
			// Name 联系人姓名
			Name string `json:"name"`
			// Phone 联系人手机号
			Phone string `json:"phone"`
		} `json:"payer_info"`
		// PaymentType 收款方式。0:在聊天中收款 1:收款码收款 2:在直播间收款 3:用产品图册收款
		PaymentType int `json:"payment_type"`
		RefundList  []struct {
			// OutRefundNo 退款单号
			OutRefundNo string `json:"out_refund_no"`
			// RefundComment 退款备注
			RefundComment string `json:"refund_comment"`
			// RefundFee 退款金额
			RefundFee int `json:"refund_fee"`
			// RefundReqtime 退款发起时间
			RefundReqtime int `json:"refund_reqtime"`
			// RefundStatus 退款状态。0:已申请退款；1:退款处理中；2:退款成功；3:退款关闭；4:退款异常；5:审批中；6:审批失败；7:审批取消
			RefundStatus int `json:"refund_status"`
			// RefundUserid 退款发起人ID
			RefundUserid string `json:"refund_userid"`
		} `json:"refund_list"` // 退款单据详情列表，退款单据详情参考refund
		// Remark 收款备注
		Remark string `json:"remark"`
		// TotalFee 收款总金额，单位为分
		TotalFee int `json:"total_fee"`
		// TotalRefundFee 退款总金额
		TotalRefundFee int `json:"total_refund_fee"`
		// TradeState 交易状态。1:已完成 3:已完成有退款
		TradeState int `json:"trade_state"`
		// TransactionID 交易单号
		TransactionID string `json:"transaction_id"`
	} `json:"bill_list"` // 交易单详情列表
	CommonResp
	// NextCursor 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetBillListExternalpay{}

func (x RespGetBillListExternalpay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetBillListExternalpay 获取对外收款记录
// 文档：https://developer.work.weixin.qq.com/document/path/93727#获取对外收款记录
func (c *ApiClient) ExecGetBillListExternalpay(req ReqGetBillListExternalpay) (RespGetBillListExternalpay, error) {
	var resp RespGetBillListExternalpay
	err := c.executeWXApiPost("/cgi-bin/externalpay/get_bill_list", req, &resp, true)
	if err != nil {
		return RespGetBillListExternalpay{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetBillListExternalpay{}, bizErr
	}
	return resp, nil
}
