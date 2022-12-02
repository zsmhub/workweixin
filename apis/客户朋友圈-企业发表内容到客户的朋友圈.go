package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqAddMomentTaskExternalcontact 创建发表任务请求
// 文档：https://developer.work.weixin.qq.com/document/path/95094#创建发表任务
type ReqAddMomentTaskExternalcontact struct {
	Attachments []struct {
		Image struct {
			// MediaID 图片的素材id。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 图片链接封面，普通图片:总像素不超过1555200。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 视频的素材id，未填写报错&#34;invalid msg&#34;。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 图片的素材id。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 图片链接封面，普通图片:总像素不超过1555200。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 视频的素材id，未填写报错&#34;invalid msg&#34;。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 图片的素材id。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 图片链接封面，普通图片:总像素不超过1555200。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			// MediaID 视频的素材id，未填写报错&#34;invalid msg&#34;。可通过<a href="#34657" rel="nofollow">上传附件资源</a>接口获得，必填
			MediaID string `json:"media_id"`
		} `json:"image,omitempty"` // 图片消息附件。普通图片:总像素不超过1555200。图片大小不超过10M。最多支持传入9个；超过9个报错&#39;invalid attachments size&#39;
		Link struct {
			MediaID string `json:"media_id"`
			// Title 图文消息标题，最多64个<strong>字节</strong>
			Title string `json:"title,omitempty"`
			// URL 图文消息链接，必填
			URL string `json:"url"`
		} `json:"link,omitempty"` // 图文消息附件。只支持1个；若超过1个报错&#39;invalid attachments size&#39;
		// Msgtype 附件类型，可选image、link或者video，必填
		Msgtype string `json:"msgtype"`
		Video   struct {
			MediaID string `json:"media_id"`
		} `json:"video,omitempty"` // 视频消息附件。最长不超过30S，最大不超过10MB。只支持1个；若超过1个报错&#39;invalid attachments size&#39;
	} `json:"attachments,omitempty"` // 附件，不能与text.content同时为空，最多支持9个图片类型，或者1个视频，或者1个链接。类型只能三选一，若传了不同类型，报错&#39;invalid attachments msgtype&#39;
	Text struct {
		// Content 消息文本内容，不能与附件同时为空，最多支持传入2000个字符，若超出长度报错&#39;invalid text size&#39;
		Content string `json:"content,omitempty"`
	} `json:"text,omitempty"` // 文本消息
	VisibleRange struct {
		ExternalContactList struct {
			// TagList 可见到该朋友圈的客户标签列表
			TagList []string `json:"tag_list,omitempty"`
		} `json:"external_contact_list,omitempty"` // 可见到该朋友圈的客户列表，详见下文的“可见范围说明”
		SenderList struct {
			// DepartmentList 发表任务的执行者部门列表
			DepartmentList []int `json:"department_list,omitempty"`
			// UserList 发表任务的执行者用户列表，最多支持10万个
			UserList []string `json:"user_list,omitempty"`
		} `json:"sender_list,omitempty"` // 发表任务的执行者列表，详见下文的“可见范围说明”
	} `json:"visible_range,omitempty"` // 指定的发表范围；若未指定，则表示执行者为应用可见范围内所有成员
}

var _ bodyer = ReqAddMomentTaskExternalcontact{}

func (x ReqAddMomentTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespAddMomentTaskExternalcontact 创建发表任务响应
// 文档：https://developer.work.weixin.qq.com/document/path/95094#创建发表任务
type RespAddMomentTaskExternalcontact struct {
	CommonResp
	// Jobid 异步任务id，最大长度为64字节，24小时有效；可使用<a href="#%E8%8E%B7%E5%8F%96%E5%8F%91%E8%A1%A8%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BB%BB%E5%8A%A1%E7%BB%93%E6%9E%9C" rel="nofollow">获取发表朋友圈任务结果</a>查询任务状态
	Jobid string `json:"jobid"`
}

var _ bodyer = RespAddMomentTaskExternalcontact{}

func (x RespAddMomentTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddMomentTaskExternalcontact 创建发表任务
// 文档：https://developer.work.weixin.qq.com/document/path/95094#创建发表任务
func (c *ApiClient) ExecAddMomentTaskExternalcontact(req ReqAddMomentTaskExternalcontact) (RespAddMomentTaskExternalcontact, error) {
	var resp RespAddMomentTaskExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/add_moment_task", req, &resp, true)
	if err != nil {
		return RespAddMomentTaskExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddMomentTaskExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentTaskResultExternalcontact 获取任务创建结果请求
// 文档：https://developer.work.weixin.qq.com/document/path/95094#获取任务创建结果
type ReqGetMomentTaskResultExternalcontact struct {
	// Jobid 异步任务id，最大长度为64字节，由<a href="#%E5%88%9B%E5%BB%BA%E5%8F%91%E8%A1%A8%E5%86%85%E5%AE%B9%E5%88%B0%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BB%BB%E5%8A%A1" rel="nofollow">创建发表内容到客户朋友圈任务</a>接口获取，必填
	Jobid string `json:"jobid"`
}

var _ urlValuer = ReqGetMomentTaskResultExternalcontact{}

func (x ReqGetMomentTaskResultExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetMomentTaskResultExternalcontact 获取任务创建结果响应
// 文档：https://developer.work.weixin.qq.com/document/path/95094#获取任务创建结果
type RespGetMomentTaskResultExternalcontact struct {
	CommonResp
	// 任务状态，整型，1表示开始创建任务，2表示正在创建任务中，3表示创建任务已完成
	Status int `json:"status"`
	// 操作类型，字节串，此处固定为add_moment_task
	Type string `json:"type"`
	// 详细的处理结果。当任务完成后此字段有效
	Result struct {
		Errcode           int    `json:"errcode"`
		Errmsg            string `json:"errmsg"`
		MomentID          string `json:"moment_id"`
		InvalidSenderList struct {
			UserList       []string `json:"user_list"`
			DepartmentList []int    `json:"department_list"`
		} `json:"invalid_sender_list"`
		InvalidExternalContactList struct {
			TagList []string `json:"tag_list"`
		} `json:"invalid_external_contact_list"`
	} `json:"result"`
}

var _ bodyer = RespGetMomentTaskResultExternalcontact{}

func (x RespGetMomentTaskResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentTaskResultExternalcontact 获取任务创建结果
// 文档：https://developer.work.weixin.qq.com/document/path/95094#获取任务创建结果
func (c *ApiClient) ExecGetMomentTaskResultExternalcontact(req ReqGetMomentTaskResultExternalcontact) (RespGetMomentTaskResultExternalcontact, error) {
	var resp RespGetMomentTaskResultExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get_moment_task_result", req, &resp, true)
	if err != nil {
		return RespGetMomentTaskResultExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentTaskResultExternalcontact{}, bizErr
	}
	return resp, nil
}
