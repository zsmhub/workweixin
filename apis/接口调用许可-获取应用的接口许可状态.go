package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAppLicenseInfoLicense 获取应用的接口许可状态请求
// 文档：https://developer.work.weixin.qq.com/document/path/95844#获取应用的接口许可状态
type ReqGetAppLicenseInfoLicense struct {
	// Appid 旧的多应用套件中的应用id，新开发者请忽略
	Appid int `json:"appid,omitempty"`
	// Corpid 企业id，必填
	Corpid string `json:"corpid"`
	// SuiteID 套件id，必填
	SuiteID string `json:"suite_id"`
}

var _ bodyer = ReqGetAppLicenseInfoLicense{}

func (x ReqGetAppLicenseInfoLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetAppLicenseInfoLicense 获取应用的接口许可状态响应
// 文档：https://developer.work.weixin.qq.com/document/path/95844#获取应用的接口许可状态
type RespGetAppLicenseInfoLicense struct {
	CommonResp
	// LicenseStatus license检查开启状态。<br/>0:未开启license检查状态（<a href="#38519" rel="nofollow">未迁移的历史授权应用</a>一般是这种状态） <br/>1:已开启license检查状态。若开启且已过试用期，则需要为企业购买license帐号才可以使用
	LicenseStatus int `json:"license_status"`
	TrailInfo     struct {
		// EndTime 接口许可试用到期时间。若企业多次安装卸载同一个第三方应用，以第一次安装的时间为试用期开始时间，第一次安装完90天后为结束试用时间
		EndTime int `json:"end_time"`
		// StartTime 接口许可试用开始时间
		StartTime int `json:"start_time"`
	} `json:"trail_info"` // 应用license试用期信息。仅当license_status为1时返回该字段
}

var _ bodyer = RespGetAppLicenseInfoLicense{}

func (x RespGetAppLicenseInfoLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAppLicenseInfoLicense 获取应用的接口许可状态
// 文档：https://developer.work.weixin.qq.com/document/path/95844#获取应用的接口许可状态
func (c *ApiClient) ExecGetAppLicenseInfoLicense(req ReqGetAppLicenseInfoLicense) (RespGetAppLicenseInfoLicense, error) {
	var resp RespGetAppLicenseInfoLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_app_license_info", req, &resp, true)
	if err != nil {
		return RespGetAppLicenseInfoLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAppLicenseInfoLicense{}, bizErr
	}
	return resp, nil
}
