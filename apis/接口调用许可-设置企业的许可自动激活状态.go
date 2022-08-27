package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSetAutoActiveStatusLicense 设置企业的许可自动激活状态请求
// 文档：https://developer.work.weixin.qq.com/document/path/95873#设置企业的许可自动激活状态
type ReqSetAutoActiveStatusLicense struct {
	// AutoActiveStatus 许可自动激活状态。0:关闭，1:打开，必填
	AutoActiveStatus int `json:"auto_active_status"`
	// Corpid 企业corpid，要求服务商为企业购买过接口许可，购买指支付完成，购买并退款成功包括在内，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqSetAutoActiveStatusLicense{}

func (x ReqSetAutoActiveStatusLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSetAutoActiveStatusLicense 设置企业的许可自动激活状态响应
// 文档：https://developer.work.weixin.qq.com/document/path/95873#设置企业的许可自动激活状态
type RespSetAutoActiveStatusLicense struct {
	CommonResp
}

var _ bodyer = RespSetAutoActiveStatusLicense{}

func (x RespSetAutoActiveStatusLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSetAutoActiveStatusLicense 设置企业的许可自动激活状态
// 文档：https://developer.work.weixin.qq.com/document/path/95873#设置企业的许可自动激活状态
func (c *ApiClient) ExecSetAutoActiveStatusLicense(req ReqSetAutoActiveStatusLicense) (RespSetAutoActiveStatusLicense, error) {
	var resp RespSetAutoActiveStatusLicense
	err := c.executeWXApiPost("/cgi-bin/license/set_auto_active_status", req, &resp, true)
	if err != nil {
		return RespSetAutoActiveStatusLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSetAutoActiveStatusLicense{}, bizErr
	}
	return resp, nil
}
