package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqProlongTryService 延长试用期请求
// 文档：https://developer.work.weixin.qq.com/document/path/91913#延长试用期
type ReqProlongTryService struct {
	// Appid 仅旧套件需要填此参数
	Appid int `json:"appid,omitempty"`
	// BuyerCorpid 购买方corpid，必填
	BuyerCorpid string `json:"buyer_corpid"`
	// ProlongDays 延长天数，必填
	ProlongDays int `json:"prolong_days"`
}

var _ bodyer = ReqProlongTryService{}

func (x ReqProlongTryService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespProlongTryService 延长试用期响应
// 文档：https://developer.work.weixin.qq.com/document/path/91913#延长试用期
type RespProlongTryService struct {
	CommonResp
	// TryEndTime 延长后的试用到期时间（UNIX时间戳）
	TryEndTime int `json:"try_end_time"`
}

var _ bodyer = RespProlongTryService{}

func (x RespProlongTryService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecProlongTryService 延长试用期
// 文档：https://developer.work.weixin.qq.com/document/path/91913#延长试用期
func (c *ApiClient) ExecProlongTryService(req ReqProlongTryService) (RespProlongTryService, error) {
	var resp RespProlongTryService
	err := c.executeWXApiPost("/cgi-bin/service/prolong_try", req, &resp, true)
	if err != nil {
		return RespProlongTryService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProlongTryService{}, bizErr
	}
	return resp, nil
}
