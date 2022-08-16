package apis

import (
	"encoding/json"
)

// 文档：https://developer.work.weixin.qq.com/document/path/93476
type OwnerFilterP struct {
	UseridList []string `json:"userid_list"`
}

type ReqGroupchatStatistic struct {
	DayBeginTime int64        `json:"day_begin_time"`
	DayEndTime   int64        `json:"day_end_time"`
	OwnerFilter  OwnerFilterP `json:"owner_filter"`
}

var _ bodyer = ReqGroupchatStatistic{}

func (x ReqGroupchatStatistic) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type RespGroupchatStatistic struct {
	CommonResp
	Items []struct {
		StatTime int64 `json:"stat_time"`
		Data     struct {
			NewChatCnt            uint64 `json:"new_chat_cnt"`
			ChatTotal             uint64 `json:"chat_total"`
			ChatHasMsg            uint64 `json:"chat_has_msg"`
			NewMemberCnt          uint64 `json:"new_member_cnt"`
			MemberTotal           uint64 `json:"member_total"`
			MemberHasMsg          uint64 `json:"member_has_msg"`
			MsgTotal              uint64 `json:"msg_total"`
			MigrateTraineeChatCnt uint64 `json:"migrate_trainee_chat_cnt"`
		} `json:"data"`
	} `json:"items"`
}

var _ bodyer = RespGroupchatStatistic{}

func (x RespGroupchatStatistic) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGroupchatStatistic(req ReqGroupchatStatistic) (RespGroupchatStatistic, error) {
	var resp RespGroupchatStatistic
	err := c.executeWXApiPost("/cgi-bin/externalcontact/groupchat/statistic_group_by_day", req, &resp, true)
	if err != nil {
		return RespGroupchatStatistic{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGroupchatStatistic{}, bizErr
	}

	return resp, nil
}
