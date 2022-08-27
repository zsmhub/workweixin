package apis

import (
	"encoding/json"
	"net/url"
)

// ReqUploadMedia 上传临时素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/90389#上传临时素材
type ReqUploadMedia struct {
	// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
	Type  string `json:"type"`
	Media *Media `json:"media"`
}

var _ mediaUploader = ReqUploadMedia{}
var _ urlValuer = ReqUploadMedia{}

func (x ReqUploadMedia) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		if v == nil {
			continue
		}
		ret.Add(k, StrVal(v))
	}
	return ret
}

func (x ReqUploadMedia) getMedia() *Media {
	return x.Media
}

// RespUploadMedia 上传临时素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/90389#上传临时素材
type RespUploadMedia struct {
	CommonResp
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

var _ bodyer = RespUploadMedia{}

func (x RespUploadMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUploadMedia 上传临时素材
// 文档：https://developer.work.weixin.qq.com/document/path/90389#上传临时素材
func (c *ApiClient) ExecUploadMedia(req ReqUploadMedia) (RespUploadMedia, error) {
	var resp RespUploadMedia
	err := c.executeWXApiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
	if err != nil {
		return RespUploadMedia{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUploadMedia{}, bizErr
	}
	return resp, nil
}
