package apis

import (
	"encoding/json"
)

// 文档：https://developer.work.weixin.qq.com/document/path/92275

type ReqUserBehaviorData struct {
	Userid    []string `json:"userid"`
	Partyid   []int    `json:"partyid"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

var _ bodyer = ReqUserBehaviorData{}

func (x ReqUserBehaviorData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespUserBehaviorData struct {
	CommonResp
	BehaviorData []struct {
		StatTime            int64   `json:"stat_time"`
		ChatCnt             uint64  `json:"chat_cnt"`
		MessageCnt          uint64  `json:"message_cnt"`
		ReplyPercentage     float64 `json:"reply_percentage"`
		AvgReplyTime        uint64  `json:"avg_reply_time"`
		NegativeFeedbackCnt uint64  `json:"negative_feedback_cnt"`
		NewApplyCnt         uint64  `json:"new_apply_cnt"`
		NewContactCnt       uint64  `json:"new_contact_cnt"`
	} `json:"behavior_data"`
}

var _ bodyer = RespUserBehaviorData{}

func (x RespUserBehaviorData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

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
