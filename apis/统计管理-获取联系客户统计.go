package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUserBehaviorData 获取「联系客户统计」数据请求
// 文档：https://developer.work.weixin.qq.com/document/path/92275#获取「联系客户统计」数据
type ReqUserBehaviorData struct {
	// EndTime 数据结束时间，必填
	EndTime int64 `json:"end_time"`
	// Partyid 部门ID列表，最多100个
	Partyid []int `json:"partyid,omitempty"`
	// StartTime 数据起始时间，必填
	StartTime int64 `json:"start_time"`
	// Userid 成员ID列表，最多100个
	Userid []string `json:"userid,omitempty"`
}

var _ bodyer = ReqUserBehaviorData{}

func (x ReqUserBehaviorData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUserBehaviorData 获取「联系客户统计」数据响应
// 文档：https://developer.work.weixin.qq.com/document/path/92275#获取「联系客户统计」数据
type RespUserBehaviorData struct {
	BehaviorData []struct {
		// AvgReplyTime 平均首次回复时长，单位为分钟，即客户主动发起聊天后，成员在一个自然日内首次回复的时长间隔为首次回复时长，所有聊天的首次回复总时长/已回复的聊天总数即为平均首次回复时长，不包括群聊，<strong>仅在确有聊天时返回</strong>。
		AvgReplyTime uint64 `json:"avg_reply_time"`
		// ChatCnt 聊天总数， 成员有主动发送过消息的单聊总数。
		ChatCnt uint64 `json:"chat_cnt"`
		// MessageCnt 发送消息数，成员在单聊中发送的消息总数。
		MessageCnt uint64 `json:"message_cnt"`
		// NegativeFeedbackCnt 删除/拉黑成员的客户数，即将成员删除或加入黑名单的客户数。
		NegativeFeedbackCnt uint64 `json:"negative_feedback_cnt"`
		// NewApplyCnt 发起申请数，成员通过「搜索手机号」、「扫一扫」、「从微信好友中添加」、「从群聊中添加」、「添加共享、分配给我的客户」、「添加单向、双向删除好友关系的好友」、「从新的联系人推荐中添加」等渠道主动向客户发起的好友申请数量。
		NewApplyCnt uint64 `json:"new_apply_cnt"`
		// NewContactCnt 新增客户数，成员新添加的客户数量。
		NewContactCnt uint64 `json:"new_contact_cnt"`
		// ReplyPercentage 已回复聊天占比，浮点型，客户主动发起聊天后，成员在一个自然日内有回复过消息的聊天数/客户主动发起的聊天数比例，不包括群聊，<strong>仅在确有聊天时返回。</strong>
		ReplyPercentage float64 `json:"reply_percentage"`
		// StatTime 数据日期，为当日0点的时间戳
		StatTime int64 `json:"stat_time"`
	} `json:"behavior_data"`
	CommonResp
}

var _ bodyer = RespUserBehaviorData{}

func (x RespUserBehaviorData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUserBehaviorData 获取「联系客户统计」数据
// 文档：https://developer.work.weixin.qq.com/document/path/92275#获取「联系客户统计」数据
func (c *ApiClient) ExecUserBehaviorData(req ReqUserBehaviorData) (RespUserBehaviorData, error) {
	var resp RespUserBehaviorData
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_user_behavior_data", req, &resp, true)
	if err != nil {
		return RespUserBehaviorData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUserBehaviorData{}, bizErr
	}
	return resp, nil
}
