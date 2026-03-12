package apis

import (
	"encoding/json"
	"net/url"
)

// ReqUploadImg 上传图片
// 文档：https://developer.work.weixin.qq.com/document/path/90392
type ReqUploadImg struct {
	Media *Media `json:"media"`
}

var _ mediaUploader = ReqUploadImg{}
var _ urlValuer = ReqUploadImg{}

func (x ReqUploadImg) intoURLValues() url.Values {
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

func (x ReqUploadImg) getMedia() *Media {
	return x.Media
}

// RespUploadImg 上传图片响应
// 文档：https://developer.work.weixin.qq.com/document/path/90392
type RespUploadImg struct {
	CommonResp
	URL string `json:"url"`
}

var _ bodyer = RespUploadImg{}

func (x RespUploadImg) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUploadImg 上传图片
// 文档：https://developer.work.weixin.qq.com/document/path/90392
func (c *ApiClient) ExecUploadImg(req ReqUploadImg) (RespUploadImg, error) {
	var resp RespUploadImg
	err := c.executeWXApiMediaUpload("/cgi-bin/media/uploadimg", req, &resp, true)
	if err != nil {
		return RespUploadImg{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUploadImg{}, bizErr
	}
	return resp, nil
}
