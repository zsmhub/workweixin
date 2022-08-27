package apis

import (
	"encoding/json"
)

// 文档：https://developer.work.weixin.qq.com/document/path/95153

type ReqGetCorpQualification struct{}

var _ bodyer = ReqGetCorpQualification{}

func (x ReqGetCorpQualification) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCorpQualification struct {
	CommonResp
	WechatChannelsBinding bool `json:"wechat_channels_binding"` // 当企业具有绑定成功的视频号时，返回true，否则返回false。
}

var _ bodyer = RespGetCorpQualification{}

func (x RespGetCorpQualification) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) GetCorpQualification(req ReqGetCorpQualification) (RespGetCorpQualification, error) {
	var resp RespGetCorpQualification
	err := c.executeWXApiPost("/cgi-bin/kf/get_corp_qualification", req, &resp, true)
	if err != nil {
		return RespGetCorpQualification{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCorpQualification{}, bizErr
	}
	return resp, nil
}
