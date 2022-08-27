package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSetSessionInfoService 设置授权配置请求
// 文档：https://developer.work.weixin.qq.com/document/path/90602#设置授权配置
type ReqSetSessionInfoService struct {
	// PreAuthCode 预授权码，必填
	PreAuthCode string `json:"pre_auth_code"`
	SessionInfo struct {
		// Appid 允许进行授权的应用id，如1、2、3， 不填或者填空数组都表示允许授权套件内所有应用（仅旧的多应用套件可传此参数，新开发者可忽略）
		Appid []int `json:"appid,omitempty"`
		// AuthType 授权类型:0 正式授权， 1 测试授权。 默认值为0。<font data-color="red" color="red">注意，请确保应用在正式发布后的授权类型为“正式授权”<font></font></font>
		AuthType int `json:"auth_type,omitempty"`
	} `json:"session_info"` // 本次授权过程中需要用到的会话信息，必填
}

var _ bodyer = ReqSetSessionInfoService{}

func (x ReqSetSessionInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSetSessionInfoService 设置授权配置响应
// 文档：https://developer.work.weixin.qq.com/document/path/90602#设置授权配置
type RespSetSessionInfoService struct {
	CommonResp
}

var _ bodyer = RespSetSessionInfoService{}

func (x RespSetSessionInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSetSessionInfoService 设置授权配置
// 文档：https://developer.work.weixin.qq.com/document/path/90602#设置授权配置
func (c *ApiClient) ExecSetSessionInfoService(req ReqSetSessionInfoService) (RespSetSessionInfoService, error) {
	var resp RespSetSessionInfoService
	err := c.executeWXApiPost("/cgi-bin/service/set_session_info", req, &resp, true)
	if err != nil {
		return RespSetSessionInfoService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSetSessionInfoService{}, bizErr
	}
	return resp, nil
}
