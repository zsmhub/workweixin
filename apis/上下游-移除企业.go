package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqRemoveCorpCorp 移除企业请求
// 文档：https://developer.work.weixin.qq.com/document/path/95822#移除企业
type ReqRemoveCorpCorp struct {
	// ChainID 上下游id，必填
	ChainID string `json:"chain_id"`
	// Corpid 需要移除的下游企业corpid
	Corpid string `json:"corpid,omitempty"`
	// PendingCorpid 需要移除的未加入下游企业corpid，corpid和pending_corpid至少填一个，都填corpid生效
	PendingCorpid string `json:"pending_corpid,omitempty"`
}

var _ bodyer = ReqRemoveCorpCorp{}

func (x ReqRemoveCorpCorp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespRemoveCorpCorp 移除企业响应
// 文档：https://developer.work.weixin.qq.com/document/path/95822#移除企业
type RespRemoveCorpCorp struct {
	CommonResp
}

var _ bodyer = RespRemoveCorpCorp{}

func (x RespRemoveCorpCorp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecRemoveCorpCorp 移除企业
// 文档：https://developer.work.weixin.qq.com/document/path/95822#移除企业
func (c *ApiClient) ExecRemoveCorpCorp(req ReqRemoveCorpCorp) (RespRemoveCorpCorp, error) {
	var resp RespRemoveCorpCorp
	err := c.executeWXApiPost("/cgi-bin/corpgroup/corp/remove_corp", req, &resp, true)
	if err != nil {
		return RespRemoveCorpCorp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespRemoveCorpCorp{}, bizErr
	}
	return resp, nil
}
