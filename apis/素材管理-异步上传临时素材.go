package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUploadByUrlMedia 生成异步上传任务请求
// 文档：https://developer.work.weixin.qq.com/document/path/96219#生成异步上传任务
type ReqUploadByUrlMedia struct {
	// Filename 文件名，标识文件展示的名称。比如，使用该media_id发消息时，展示的文件名由该字段控制。<br/>不超过128字节。必填
	Filename string `json:"filename"`
	// Md5 文件md5。对比从url下载下来的文件md5是否一致。<br/>不超过32字节。必填
	Md5 string `json:"md5"`
	// Scene 场景值。1-客户联系入群欢迎语素材（目前仅支持1）。<br/>注意:每个场景值有对应的使用范围，详见上面的「使用场景说明」，必填
	Scene int `json:"scene"`
	// Type 媒体文件类型。目前仅支持video-视频，file-普通文件<br/>不超过32字节。必填
	Type string `json:"type"`
	// URL 文件cdn url。url要求支持Range分块下载<br/>不超过1024字节。必填
	URL string `json:"url"`
}

var _ bodyer = ReqUploadByUrlMedia{}

func (x ReqUploadByUrlMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUploadByUrlMedia 生成异步上传任务响应
// 文档：https://developer.work.weixin.qq.com/document/path/96219#生成异步上传任务
type RespUploadByUrlMedia struct {
	CommonResp
	// 任务id。可通过此jobid查询结果
	Jobid string `json:"jobid"`
}

var _ bodyer = RespUploadByUrlMedia{}

func (x RespUploadByUrlMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUploadByUrlMedia 生成异步上传任务
// 文档：https://developer.work.weixin.qq.com/document/path/96219#生成异步上传任务
func (c *ApiClient) ExecUploadByUrlMedia(req ReqUploadByUrlMedia) (RespUploadByUrlMedia, error) {
	var resp RespUploadByUrlMedia
	err := c.executeWXApiPost("/cgi-bin/media/upload_by_url", req, &resp, true)
	if err != nil {
		return RespUploadByUrlMedia{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUploadByUrlMedia{}, bizErr
	}
	return resp, nil
}

// ReqGetUploadByUrlResultMedia 查询异步任务结果请求
// 文档：https://developer.work.weixin.qq.com/document/path/96219#查询异步任务结果
type ReqGetUploadByUrlResultMedia struct {
	// Jobid 任务id。最长为128字节，<strong>60分钟内有效</strong>，必填
	Jobid string `json:"jobid"`
}

var _ bodyer = ReqGetUploadByUrlResultMedia{}

func (x ReqGetUploadByUrlResultMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetUploadByUrlResultMedia 查询异步任务结果响应
// 文档：https://developer.work.weixin.qq.com/document/path/96219#查询异步任务结果
type RespGetUploadByUrlResultMedia struct {
	CommonResp
	Detail struct {
		CommonResp
		// CreatedAt 媒体文件创建的时间戳。当status为<code>2</code>时返回。
		CreatedAt string `json:"created_at"`
		// MediaID 媒体文件上传后获取的唯一标识，3天内有效。当status为<code>2</code>时返回。
		MediaID string `json:"media_id"`
	} `json:"detail"` // 结果明细
	// Status 任务状态。1-处理中，2-完成，3-异常失败
	Status int `json:"status"`
}

var _ bodyer = RespGetUploadByUrlResultMedia{}

func (x RespGetUploadByUrlResultMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUploadByUrlResultMedia 查询异步任务结果
// 文档：https://developer.work.weixin.qq.com/document/path/96219#查询异步任务结果
func (c *ApiClient) ExecGetUploadByUrlResultMedia(req ReqGetUploadByUrlResultMedia) (RespGetUploadByUrlResultMedia, error) {
	var resp RespGetUploadByUrlResultMedia
	err := c.executeWXApiPost("/cgi-bin/media/get_upload_by_url_result", req, &resp, true)
	if err != nil {
		return RespGetUploadByUrlResultMedia{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUploadByUrlResultMedia{}, bizErr
	}
	return resp, nil
}
