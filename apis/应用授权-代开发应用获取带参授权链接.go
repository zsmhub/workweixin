package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetCustomizedAuthUrlService 获取带参授权链接请求
// 文档：https://developer.work.weixin.qq.com/document/path/98744#获取带参授权链接
type ReqGetCustomizedAuthUrlService struct {
	// State state值，可以填写a-zA-Z0-9，长度不可超过32个字节。当扫带参二维码授权代开发模版时，<a href="#57310" rel="nofollow">获取企业永久授权码</a>接口会返回该state值，必填
	State string `json:"state"`
	// TemplateidList 代开发自建应用模版ID列表，数量不能超过9个，必填
	TemplateidList []string `json:"templateid_list"`
}

var _ bodyer = ReqGetCustomizedAuthUrlService{}

func (x ReqGetCustomizedAuthUrlService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetCustomizedAuthUrlService 获取带参授权链接响应
// 文档：https://developer.work.weixin.qq.com/document/path/98744#获取带参授权链接
type RespGetCustomizedAuthUrlService struct {
	CommonResp
	// ExpiresIn 有效期（秒）。10天过期。
	ExpiresIn int `json:"expires_in"`
	// QrcodeURL 可用来生成二维码的授权url，需要开发者自行生成为二维码
	QrcodeURL string `json:"qrcode_url"`
}

var _ bodyer = RespGetCustomizedAuthUrlService{}

func (x RespGetCustomizedAuthUrlService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetCustomizedAuthUrlService 获取带参授权链接
// 文档：https://developer.work.weixin.qq.com/document/path/98744#获取带参授权链接
func (c *ApiClient) ExecGetCustomizedAuthUrlService(req ReqGetCustomizedAuthUrlService) (RespGetCustomizedAuthUrlService, error) {
	var resp RespGetCustomizedAuthUrlService
	err := c.executeWXApiPost("/cgi-bin/service/get_customized_auth_url", req, &resp, true)
	if err != nil {
		return RespGetCustomizedAuthUrlService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCustomizedAuthUrlService{}, bizErr
	}
	return resp, nil
}
