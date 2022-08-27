package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqGetByUserBatch 批量获取客户详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/93010#批量获取客户详情
type ReqGetByUserBatch struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit int `json:"limit"`
	// UseridList 企业成员的userid列表，字符串类型，最多支持100个，必填
	UseridList []string `json:"userid_list"`
}

var _ bodyer = ReqGetByUserBatch{}

func (x ReqGetByUserBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetByUserBatch 批量获取客户详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/93010#批量获取客户详情
type RespGetByUserBatch struct {
	CommonResp
	ExternalContactList []struct {
		ExternalContact struct {
			Avatar          string `json:"avatar"`
			CorpFullName    string `json:"corp_full_name"`
			CorpName        string `json:"corp_name"`
			ExternalProfile struct {
				ExternalAttr []struct {
					Miniprogram struct {
						Appid    string `json:"appid"`
						Pagepath string `json:"pagepath"`
						Title    string `json:"title"`
					} `json:"miniprogram"`
					Name string `json:"name"`
					Text struct {
						Value string `json:"value"`
					} `json:"text"`
					Type int `json:"type"`
					Web  struct {
						Title string `json:"title"`
						URL   string `json:"url"`
					} `json:"web"`
				} `json:"external_attr"`
			} `json:"external_profile"`
			ExternalUserid string `json:"external_userid"`
			Gender         int    `json:"gender"`
			Name           string `json:"name"`
			Position       string `json:"position"`
			Type           int    `json:"type"`
			Unionid        string `json:"unionid"`
		} `json:"external_contact"`
		FollowInfo struct {
			AddWay         int      `json:"add_way"`
			Createtime     int      `json:"createtime"`
			Description    string   `json:"description"`
			OperUserid     string   `json:"oper_userid"`
			Remark         string   `json:"remark"`
			RemarkCorpName string   `json:"remark_corp_name"`
			RemarkMobiles  []string `json:"remark_mobiles"`
			State          string   `json:"state"`
			TagID          []string `json:"tag_id"`
			Userid         string   `json:"userid"`
			WechatChannels struct {
				Nickname string `json:"nickname"`
			} `json:"wechat_channels"`
		} `json:"follow_info"`
	} `json:"external_contact_list"`
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetByUserBatch{}

func (x RespGetByUserBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetByUserBatch 批量获取客户详情
// 文档：https://developer.work.weixin.qq.com/document/path/93010#批量获取客户详情
func (c *ApiClient) ExecGetByUserBatch(req ReqGetByUserBatch) (RespGetByUserBatch, error) {
	var resp RespGetByUserBatch
	err := c.executeWXApiPost("/cgi-bin/externalcontact/batch/get_by_user", req, &resp, true)
	if err != nil {
		return RespGetByUserBatch{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetByUserBatch{}, bizErr
	}
	return resp, nil
}
