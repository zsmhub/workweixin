package apis

import (
	"encoding/json"
)

// ReqOpenOrderPaytool 创建收款订单
// 文档：https://developer.work.weixin.qq.com/document/path/98045#创建收款订单

type (
	ReqOpenOrderPaytool struct {
		BusinessType       int                `json:"business_type"`
		PayType            int                `json:"pay_type"`
		BankReceiptMediaId string             `json:"bank_receipt_media_id,omitempty"`
		Creator            string             `json:"creator,omitempty"`
		CustomCorpid       string             `json:"custom_corpid,omitempty"`
		ProductList        PaytoolProductList `json:"product_list"`
		NonceStr           string             `json:"nonce_str"`
		Ts                 int                `json:"ts"`
		Sig                string             `json:"sig"`
	}

	PaytoolProductList struct {
		ThirdApp      *PaytoolThirdApp      `json:"third_app,omitempty"`
		CustomizedApp *PaytoolCustomizedApp `json:"customized_app,omitempty"`
		PromotionCase *PaytoolPromotionCase `json:"promotion_case,omitempty"`
	}

	PaytoolThirdApp struct {
		OrderType   int                  `json:"order_type"`
		BuyInfoList []PaytoolBuyInfoList `json:"buy_info_list"`
	}

	PaytoolBuyInfoList struct {
		Suiteid        string               `json:"suiteid"`
		Appid          int                  `json:"appid,omitempty"`
		EditionId      string               `json:"edition_id,omitempty"`
		DurationDays   int                  `json:"duration_days,omitempty"`
		TakeEffectDate string               `json:"take_effect_date,omitempty"`
		UserCount      int                  `json:"user_count,omitempty"`
		DiscountInfo   *PaytoolDiscountInfo `json:"discount_info,omitempty"`
	}

	PaytoolDiscountInfo struct {
		DiscountType    int    `json:"discount_type"`
		DiscountAmount  int    `json:"discount_amount,omitempty"`
		DiscountRatio   int    `json:"discount_ratio,omitempty"`
		DiscountRemarks string `json:"discount_remarks"`
	}

	PaytoolPromotionCase struct {
		OrderType            int                  `json:"order_type"`
		CaseId               string               `json:"case_id"`
		PromotionEditionName string               `json:"promotion_edition_name"`
		DurationDays         int                  `json:"duration_days,omitempty"`
		TakeEffectDate       string               `json:"take_effect_date,omitempty"`
		BuyInfoList          []PaytoolBuyInfoList `json:"buy_info_list"`
	}

	PaytoolCustomizedApp struct {
		OrderType   int                  `json:"order_type"`
		BuyInfoList []PaytoolBuyInfoList `json:"buy_info_list"`
	}
)

var _ bodyer = ReqOpenOrderPaytool{}

func (x ReqOpenOrderPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOpenOrderPaytool struct {
	CommonResp
	OrderId     string `json:"order_id"`
	OrderUrl    string `json:"order_url"`
	OriginPrice int    `json:"origin_price"`
	PaidPrice   int    `json:"paid_price"`
}

var _ bodyer = RespOpenOrderPaytool{}

func (x RespOpenOrderPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOpenOrderPaytool(req ReqOpenOrderPaytool) (RespOpenOrderPaytool, error) {
	var resp RespOpenOrderPaytool
	err := c.executeWXApiPost("/cgi-bin/paytool/open_order", req, &resp, true)
	if err != nil {
		return RespOpenOrderPaytool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOpenOrderPaytool{}, bizErr
	}
	return resp, nil
}
