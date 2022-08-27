package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

type OwnerFilterP struct {
	// UseridList 群主ID列表。最多100个，必填
	UseridList []string `json:"userid_list"`
}

// ReqStatisticGroupchat 按群主聚合的方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按群主聚合的方式
type ReqStatisticGroupchat struct {
	// DayBeginTime 起始日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围:昨天至前180天。，必填
	DayBeginTime int `json:"day_begin_time"`
	// DayEndTime 结束日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围:昨天至前180天。<br/><strong>如果不填，默认同 day_begin_time（即默认取一天的数据）</strong>
	DayEndTime int `json:"day_end_time,omitempty"`
	// Limit 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
	Limit int `json:"limit,omitempty"`
	// Offset 分页，偏移量, 默认为0
	Offset int `json:"offset,omitempty"`
	// OrderAsc 是否升序。0-否；1-是。默认降序
	OrderAsc int `json:"order_asc,omitempty"`
	// OrderBy 排序方式。<br/>1 - 新增群的数量<br/>2 - 群总数<br/>3 - 新增群人数<br/>4 - 群总人数<br/><br/>默认为1
	OrderBy     int          `json:"order_by,omitempty"`
	OwnerFilter OwnerFilterP `json:"owner_filter"` // 群主过滤。<br/>如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017），必填
}

var _ bodyer = ReqStatisticGroupchat{}

func (x ReqStatisticGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespStatisticGroupchat 按群主聚合的方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按群主聚合的方式
type RespStatisticGroupchat struct {
	CommonResp
	Items []struct {
		Data struct {
			// ChatHasMsg 截至当天有发过消息的客户群数量
			ChatHasMsg int `json:"chat_has_msg"`
			// ChatTotal 截至当天客户群总数量
			ChatTotal int `json:"chat_total"`
			// MemberHasMsg 截至当天有发过消息的群成员数
			MemberHasMsg int `json:"member_has_msg"`
			// MemberTotal 截至当天客户群总人数
			MemberTotal int `json:"member_total"`
			// MigrateTraineeChatCnt 截至当天新增迁移群数(仅教培行业返回)
			MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
			// MsgTotal 截至当天客户群消息总数
			MsgTotal int `json:"msg_total"`
			// NewChatCnt 新增客户群数量
			NewChatCnt int `json:"new_chat_cnt"`
			// NewMemberCnt 客户群新增群人数。
			NewMemberCnt int `json:"new_member_cnt"`
		} `json:"data"` // 详情
		// Owner 群主ID
		Owner string `json:"owner"`
	} `json:"items"` // 记录列表。表示某个群主所拥有的客户群的统计数据
	// NextOffset 当前分页的下一个offset。当next_offset和total相等时，说明已经取完所有
	NextOffset int `json:"next_offset"`
	// Total 命中过滤条件的记录总个数
	Total int `json:"total"`
}

var _ bodyer = RespStatisticGroupchat{}

func (x RespStatisticGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecStatisticGroupchat 按群主聚合的方式
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按群主聚合的方式
func (c *ApiClient) ExecStatisticGroupchat(req ReqStatisticGroupchat) (RespStatisticGroupchat, error) {
	var resp RespStatisticGroupchat
	err := c.executeWXApiPost("/cgi-bin/externalcontact/groupchat/statistic", req, &resp, true)
	if err != nil {
		return RespStatisticGroupchat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespStatisticGroupchat{}, bizErr
	}
	return resp, nil
}

// ReqGroupchatStatistic 按自然日聚合的方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按自然日聚合的方式
type ReqGroupchatStatistic struct {
	// DayBeginTime 起始日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围:昨天至前180天。，必填
	DayBeginTime int64 `json:"day_begin_time"`
	// DayEndTime 结束日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围:昨天至前180天。<br/><strong>如果不填，默认同 day_begin_time（即默认取一天的数据）</strong>
	DayEndTime  int64        `json:"day_end_time,omitempty"`
	OwnerFilter OwnerFilterP `json:"owner_filter"` // 群主过滤。<br/>如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017），必填
}

var _ bodyer = ReqGroupchatStatistic{}

func (x ReqGroupchatStatistic) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGroupchatStatistic 按自然日聚合的方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按自然日聚合的方式
type RespGroupchatStatistic struct {
	CommonResp
	Items []struct {
		Data struct {
			// ChatHasMsg 截至当天有发过消息的客户群数量
			ChatHasMsg uint64 `json:"chat_has_msg"`
			// ChatTotal 截至当天客户群总数量
			ChatTotal uint64 `json:"chat_total"`
			// MemberHasMsg 截至当天有发过消息的群成员数
			MemberHasMsg uint64 `json:"member_has_msg"`
			// MemberTotal 截至当天客户群总人数
			MemberTotal uint64 `json:"member_total"`
			// MigrateTraineeChatCnt 截至当天新增迁移群数(仅教培行业返回)
			MigrateTraineeChatCnt uint64 `json:"migrate_trainee_chat_cnt"`
			// MsgTotal 截至当天客户群消息总数
			MsgTotal uint64 `json:"msg_total"`
			// NewChatCnt 新增客户群数量
			NewChatCnt uint64 `json:"new_chat_cnt"`
			// NewMemberCnt 客户群新增群人数。
			NewMemberCnt uint64 `json:"new_member_cnt"`
		} `json:"data"` // 详情
		// StatTime 数据日期，为当日0点的时间戳
		StatTime int64 `json:"stat_time"`
	} `json:"items"` // 记录列表。表示某个自然日客户群的统计数据
}

var _ bodyer = RespGroupchatStatistic{}

func (x RespGroupchatStatistic) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGroupchatStatistic 按自然日聚合的方式
// 文档：https://developer.work.weixin.qq.com/document/path/93476#按自然日聚合的方式
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
