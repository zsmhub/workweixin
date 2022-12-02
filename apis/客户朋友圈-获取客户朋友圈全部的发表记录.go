package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetMomentListExternalcontact 获取企业全部的发表列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取企业全部的发表列表
type ReqGetMomentListExternalcontact struct {
	// Creator 朋友圈创建人的userid
	Creator string `json:"creator,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// EndTime 朋友圈记录结束时间。Unix时间戳，必填
	EndTime int `json:"end_time"`
	// FilterType 朋友圈类型。0:企业发表 1:个人发表 2:所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type,omitempty"`
	// Limit 返回的最大记录数，整型，最大值20，默认值20，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// StartTime 朋友圈记录开始时间。Unix时间戳，必填
	StartTime int `json:"start_time"`
}

var _ bodyer = ReqGetMomentListExternalcontact{}

func (x ReqGetMomentListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetMomentListExternalcontact 获取企业全部的发表列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取企业全部的发表列表
type RespGetMomentListExternalcontact struct {
	CommonResp
	MomentList []struct {
		// CreateTime 创建时间
		CreateTime int `json:"create_time"`
		// CreateType 朋友圈创建来源。0:企业 1:个人
		CreateType int `json:"create_type"`
		// Creator 朋友圈创建者userid，<a href="#34646" rel="nofollow">企业发表内容到客户的朋友圈</a>接口创建的朋友圈不再返回该字段
		Creator string `json:"creator"`
		Image   []struct {
			// MediaID 图片的media_id列表，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
			// MediaID 视频media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
			MediaID string `json:"media_id"`
		} `json:"image"`
		Link struct {
			// Title 网页链接标题
			Title string `json:"title"`
			// URL 网页链接url
			URL string `json:"url"`
		} `json:"link"`
		Location struct {
			// Latitude 地理位置纬度
			Latitude string `json:"latitude"`
			// Longitude 地理位置经度
			Longitude string `json:"longitude"`
			// Name 地理位置名称
			Name string `json:"name"`
		} `json:"location"`
		// MomentID 朋友圈id
		MomentID string `json:"moment_id"`
		Text     struct {
			// Content 文本消息结构
			Content string `json:"content"`
		} `json:"text"`
		Video struct {
			// MediaID 图片的media_id列表，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
			// MediaID 视频media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
			MediaID string `json:"media_id"`
			// ThumbMediaID 视频封面media_id，可以通过<a href="#10115" rel="nofollow">获取临时素材</a>下载资源
			ThumbMediaID string `json:"thumb_media_id"`
		} `json:"video"`
		// VisibleType 可见范围类型。0:部分可见 1:公开
		VisibleType int `json:"visible_type"`
	} `json:"moment_list"` // 朋友圈列表
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetMomentListExternalcontact{}

func (x RespGetMomentListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentListExternalcontact 获取企业全部的发表列表
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取企业全部的发表列表
func (c *ApiClient) ExecGetMomentListExternalcontact(req ReqGetMomentListExternalcontact) (RespGetMomentListExternalcontact, error) {
	var resp RespGetMomentListExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_moment_list", req, &resp, true)
	if err != nil {
		return RespGetMomentListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentListExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentTaskExternalcontact 获取客户朋友圈企业发表的列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈企业发表的列表
type ReqGetMomentTaskExternalcontact struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// MomentID 朋友圈id,仅支持企业发表的朋友圈id，必填
	MomentID string `json:"moment_id"`
}

var _ bodyer = ReqGetMomentTaskExternalcontact{}

func (x ReqGetMomentTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetMomentTaskExternalcontact 获取客户朋友圈企业发表的列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈企业发表的列表
type RespGetMomentTaskExternalcontact struct {
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		// PublishStatus 成员发表状态。0:未发表 1:已发表
		PublishStatus int `json:"publish_status"`
		// Userid 发表成员用户userid
		Userid string `json:"userid"`
	} `json:"task_list"` // 发表任务列表
}

var _ bodyer = RespGetMomentTaskExternalcontact{}

func (x RespGetMomentTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentTaskExternalcontact 获取客户朋友圈企业发表的列表
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈企业发表的列表
func (c *ApiClient) ExecGetMomentTaskExternalcontact(req ReqGetMomentTaskExternalcontact) (RespGetMomentTaskExternalcontact, error) {
	var resp RespGetMomentTaskExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_moment_task", req, &resp, true)
	if err != nil {
		return RespGetMomentTaskExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentTaskExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentCustomerListExternalcontact 获取客户朋友圈发表时选择的可见范围请求
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表时选择的可见范围
type ReqGetMomentCustomerListExternalcontact struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// MomentID 朋友圈id，必填
	MomentID string `json:"moment_id"`
	// Userid 企业发表成员userid，如果是企业创建的朋友圈，可以通过<a href="#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8" rel="nofollow">获取客户朋友圈企业发表的列表</a>获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetMomentCustomerListExternalcontact{}

func (x ReqGetMomentCustomerListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetMomentCustomerListExternalcontact 获取客户朋友圈发表时选择的可见范围响应
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表时选择的可见范围
type RespGetMomentCustomerListExternalcontact struct {
	CustomerList []struct {
		// ExternalUserid 发送成功的外部联系人userid
		ExternalUserid string `json:"external_userid"`
		// Userid 发表成员用户userid
		Userid string `json:"userid"`
	} `json:"customer_list"` // 成员可见客户列表
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetMomentCustomerListExternalcontact{}

func (x RespGetMomentCustomerListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentCustomerListExternalcontact 获取客户朋友圈发表时选择的可见范围
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表时选择的可见范围
func (c *ApiClient) ExecGetMomentCustomerListExternalcontact(req ReqGetMomentCustomerListExternalcontact) (RespGetMomentCustomerListExternalcontact, error) {
	var resp RespGetMomentCustomerListExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_moment_customer_list", req, &resp, true)
	if err != nil {
		return RespGetMomentCustomerListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentCustomerListExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentSendResultExternalcontact 获取客户朋友圈发表后的可见客户列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表后的可见客户列表
type ReqGetMomentSendResultExternalcontact struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值5000，默认值3000，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// MomentID 朋友圈id，必填
	MomentID string `json:"moment_id"`
	// Userid 企业发表成员userid，如果是企业创建的朋友圈，可以通过<a href="#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8" rel="nofollow">获取客户朋友圈企业发表的列表</a>获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetMomentSendResultExternalcontact{}

func (x ReqGetMomentSendResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetMomentSendResultExternalcontact 获取客户朋友圈发表后的可见客户列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表后的可见客户列表
type RespGetMomentSendResultExternalcontact struct {
	CustomerList []struct {
		// ExternalUserid 成员发送成功的外部联系人userid
		ExternalUserid string `json:"external_userid"`
	} `json:"customer_list"` // 成员发送成功客户列表
	CommonResp
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetMomentSendResultExternalcontact{}

func (x RespGetMomentSendResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentSendResultExternalcontact 获取客户朋友圈发表后的可见客户列表
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈发表后的可见客户列表
func (c *ApiClient) ExecGetMomentSendResultExternalcontact(req ReqGetMomentSendResultExternalcontact) (RespGetMomentSendResultExternalcontact, error) {
	var resp RespGetMomentSendResultExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_moment_send_result", req, &resp, true)
	if err != nil {
		return RespGetMomentSendResultExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentSendResultExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetMomentCommentsExternalcontact 获取客户朋友圈的互动数据请求
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈的互动数据
type ReqGetMomentCommentsExternalcontact struct {
	// MomentID 朋友圈id，必填
	MomentID string `json:"moment_id"`
	// Userid 企业发表成员userid，如果是企业创建的朋友圈，可以通过<a href="#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8" rel="nofollow">获取客户朋友圈企业发表的列表</a>获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetMomentCommentsExternalcontact{}

func (x ReqGetMomentCommentsExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetMomentCommentsExternalcontact 获取客户朋友圈的互动数据响应
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈的互动数据
type RespGetMomentCommentsExternalcontact struct {
	CommentList []struct {
		// CreateTime 评论时间
		// CreateTime 点赞时间
		// CreateTime 评论时间
		// CreateTime 点赞时间
		CreateTime int `json:"create_time"`
		// ExternalUserid 评论的外部联系人userid
		// ExternalUserid 点赞的外部联系人userid
		// ExternalUserid 评论的外部联系人userid
		// ExternalUserid 点赞的外部联系人userid
		ExternalUserid string `json:"external_userid"`
		// Userid 评论的企业成员userid，userid与external_userid不会同时出现
		// Userid 点赞的企业成员userid，userid与external_userid不会同时出现
		// Userid 评论的企业成员userid，userid与external_userid不会同时出现
		// Userid 点赞的企业成员userid，userid与external_userid不会同时出现
		Userid string `json:"userid"`
	} `json:"comment_list"` // 评论列表
	CommonResp
	LikeList []struct {
		CreateTime     int    `json:"create_time"`
		ExternalUserid string `json:"external_userid"`
		Userid         string `json:"userid"`
	} `json:"like_list"` // 点赞列表
}

var _ bodyer = RespGetMomentCommentsExternalcontact{}

func (x RespGetMomentCommentsExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMomentCommentsExternalcontact 获取客户朋友圈的互动数据
// 文档：https://developer.work.weixin.qq.com/document/path/93333#获取客户朋友圈的互动数据
func (c *ApiClient) ExecGetMomentCommentsExternalcontact(req ReqGetMomentCommentsExternalcontact) (RespGetMomentCommentsExternalcontact, error) {
	var resp RespGetMomentCommentsExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_moment_comments", req, &resp, true)
	if err != nil {
		return RespGetMomentCommentsExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMomentCommentsExternalcontact{}, bizErr
	}
	return resp, nil
}
