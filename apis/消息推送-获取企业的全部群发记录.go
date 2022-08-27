package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetGroupmsgListV2Externalcontact 获取群发记录列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发记录列表
type ReqGetGroupmsgListV2Externalcontact struct {
	// ChatType 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群，必填
	ChatType string `json:"chat_type"`
	// Creator 群发任务创建人企业账号id
	Creator string `json:"creator"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// EndTime 群发任务记录结束时间，必填
	EndTime int `json:"end_time"`
	// FilterType 创建人类型。0:企业发表 1:个人发表 2:所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type"`
	// Limit 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取默认值
	Limit int `json:"limit"`
	// StartTime 群发任务记录开始时间，必填
	StartTime int `json:"start_time"`
}

var _ bodyer = ReqGetGroupmsgListV2Externalcontact{}

func (x ReqGetGroupmsgListV2Externalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetGroupmsgListV2Externalcontact 获取群发记录列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发记录列表
type RespGetGroupmsgListV2Externalcontact struct {
	CommonResp
	GroupMsgList []struct {
		Attachments []struct {
			File struct {
				// MediaID 图片的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 视频的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 文件的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 图片的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 视频的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 文件的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 图片的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 视频的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				// MediaID 文件的media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
				MediaID string `json:"media_id"`
			} `json:"file"`
			Image struct {
				MediaID string `json:"media_id"`
				// PicURL 图片的url，与图片的media_id不能共存优先吐出media_id
				PicURL string `json:"pic_url"`
			} `json:"image"`
			Link struct {
				// Desc 图文消息的描述，最多512个<strong>字节</strong>
				Desc string `json:"desc"`
				// Picurl 图文消息封面的url
				Picurl string `json:"picurl"`
				// Title 图文消息标题
				// Title 小程序消息标题，最多64个<strong>字节</strong>
				Title string `json:"title"`
				// URL 图文消息的链接
				URL string `json:"url"`
			} `json:"link"`
			Miniprogram struct {
				// Appid 小程序appid，必须是<strong>关联到企业的小程序应用</strong>
				Appid string `json:"appid"`
				// Page 小程序page路径
				Page       string `json:"page"`
				PicMediaID string `json:"pic_media_id"`
				// Title 图文消息标题
				// Title 小程序消息标题，最多64个<strong>字节</strong>
				Title string `json:"title"`
			} `json:"miniprogram"`
			// Msgtype 值必须是image
			// Msgtype 值必须是link
			// Msgtype 值必须是miniprogram
			// Msgtype 值必须是video
			// Msgtype 值必须是file
			Msgtype string `json:"msgtype"`
			Video   struct {
				MediaID string `json:"media_id"`
			} `json:"video"`
		} `json:"attachments"`
		// CreateTime 创建时间
		CreateTime string `json:"create_time"`
		// CreateType 群发消息创建来源。0:企业 1:个人
		CreateType int `json:"create_type"`
		// Creator 群发消息创建者userid，<a href="#15836" rel="nofollow">API接口</a>创建的群发消息不返回该字段
		Creator string `json:"creator"`
		// Msgid 企业群发消息的id，可用于<a href="" rel="nofollow">获取企业群发成员执行结果</a>
		Msgid string `json:"msgid"`
		Text  struct {
			// Content 消息文本内容，最多4000个<strong>字节</strong>
			Content string `json:"content"`
		} `json:"text"`
	} `json:"group_msg_list"` // 群发记录列表
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetGroupmsgListV2Externalcontact{}

func (x RespGetGroupmsgListV2Externalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetGroupmsgListV2Externalcontact 获取群发记录列表
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发记录列表
func (c *ApiClient) ExecGetGroupmsgListV2Externalcontact(req ReqGetGroupmsgListV2Externalcontact) (RespGetGroupmsgListV2Externalcontact, error) {
	var resp RespGetGroupmsgListV2Externalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_groupmsg_list_v2", req, &resp, true)
	if err != nil {
		return RespGetGroupmsgListV2Externalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetGroupmsgListV2Externalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetGroupmsgTaskExternalcontact 获取群发成员发送任务列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发成员发送任务列表
type ReqGetGroupmsgTaskExternalcontact struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit"`
	// Msgid 群发消息的id，通过<a href="#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8" rel="nofollow">获取群发记录列表</a>接口返回，必填
	Msgid string `json:"msgid"`
}

var _ bodyer = ReqGetGroupmsgTaskExternalcontact{}

func (x ReqGetGroupmsgTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetGroupmsgTaskExternalcontact 获取群发成员发送任务列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发成员发送任务列表
type RespGetGroupmsgTaskExternalcontact struct {
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		// SendTime 发送时间，未发送时不返回
		SendTime int `json:"send_time"`
		// Status 发送状态:0-未发送 2-已发送
		Status int `json:"status"`
		// Userid 企业服务人员的userid
		Userid string `json:"userid"`
	} `json:"task_list"` // 群发成员发送任务列表
}

var _ bodyer = RespGetGroupmsgTaskExternalcontact{}

func (x RespGetGroupmsgTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetGroupmsgTaskExternalcontact 获取群发成员发送任务列表
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取群发成员发送任务列表
func (c *ApiClient) ExecGetGroupmsgTaskExternalcontact(req ReqGetGroupmsgTaskExternalcontact) (RespGetGroupmsgTaskExternalcontact, error) {
	var resp RespGetGroupmsgTaskExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_groupmsg_task", req, &resp, true)
	if err != nil {
		return RespGetGroupmsgTaskExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetGroupmsgTaskExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetGroupmsgSendResultExternalcontact 获取企业群发成员执行结果请求
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取企业群发成员执行结果
type ReqGetGroupmsgSendResultExternalcontact struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
	Msgid  string `json:"msgid"`
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetGroupmsgSendResultExternalcontact{}

func (x ReqGetGroupmsgSendResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetGroupmsgSendResultExternalcontact 获取企业群发成员执行结果响应
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取企业群发成员执行结果
type RespGetGroupmsgSendResultExternalcontact struct {
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
	SendList   []struct {
		// ChatID 外部客户群id，群发消息到客户不返回该字段
		ChatID string `json:"chat_id"`
		// ExternalUserid 外部联系人userid，群发消息到企业的客户群不返回该字段
		ExternalUserid string `json:"external_userid"`
		// SendTime 发送时间，发送状态为1时返回
		SendTime int `json:"send_time"`
		// Status 发送状态:0-未发送 1-已发送 2-因客户不是好友导致发送失败 3-因客户已经收到其他群发消息导致发送失败
		Status int `json:"status"`
		// Userid 企业服务人员的userid
		Userid string `json:"userid"`
	} `json:"send_list"` // 群成员发送结果列表
}

var _ bodyer = RespGetGroupmsgSendResultExternalcontact{}

func (x RespGetGroupmsgSendResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetGroupmsgSendResultExternalcontact 获取企业群发成员执行结果
// 文档：https://developer.work.weixin.qq.com/document/path/93338#获取企业群发成员执行结果
func (c *ApiClient) ExecGetGroupmsgSendResultExternalcontact(req ReqGetGroupmsgSendResultExternalcontact) (RespGetGroupmsgSendResultExternalcontact, error) {
	var resp RespGetGroupmsgSendResultExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_groupmsg_send_result", req, &resp, true)
	if err != nil {
		return RespGetGroupmsgSendResultExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetGroupmsgSendResultExternalcontact{}, bizErr
	}
	return resp, nil
}
