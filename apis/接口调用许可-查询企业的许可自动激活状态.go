package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAutoActiveStatusLicense 查询企业的许可自动激活状态请求
// 文档：https://developer.work.weixin.qq.com/document/path/95874#查询企业的许可自动激活状态
type ReqGetAutoActiveStatusLicense struct {
	// Corpid 查询的企业corpid。要求服务商为企业购买过接口许可才有查询结果。，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqGetAutoActiveStatusLicense{}

func (x ReqGetAutoActiveStatusLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetAutoActiveStatusLicense 查询企业的许可自动激活状态响应
// 文档：https://developer.work.weixin.qq.com/document/path/95874#查询企业的许可自动激活状态
type RespGetAutoActiveStatusLicense struct {
	// AutoActiveStatus 许可自动激活状态。0:关闭，1:打开
	AutoActiveStatus int `json:"auto_active_status"`
	CommonResp
}

var _ bodyer = RespGetAutoActiveStatusLicense{}

func (x RespGetAutoActiveStatusLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAutoActiveStatusLicense 查询企业的许可自动激活状态
// 文档：https://developer.work.weixin.qq.com/document/path/95874#查询企业的许可自动激活状态
func (c *ApiClient) ExecGetAutoActiveStatusLicense(req ReqGetAutoActiveStatusLicense) (RespGetAutoActiveStatusLicense, error) {
	var resp RespGetAutoActiveStatusLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_auto_active_status", req, &resp, true)
	if err != nil {
		return RespGetAutoActiveStatusLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAutoActiveStatusLicense{}, bizErr
	}
	return resp, nil
}
