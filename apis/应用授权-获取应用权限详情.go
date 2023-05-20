package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// 文档：https://developer.work.weixin.qq.com/document/path/99052#获取应用权限详情
type ReqGetPermissionsAgent struct{}

var _ bodyer = ReqGetPermissionsAgent{}

func (x ReqGetPermissionsAgent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文档：https://developer.work.weixin.qq.com/document/path/99052#获取应用权限详情
type RespGetPermissionsAgent struct {
	CommonResp
	AppPermissions []string `json:"app_permissions"`
}

var _ bodyer = RespGetPermissionsAgent{}

func (x RespGetPermissionsAgent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文档：https://developer.work.weixin.qq.com/document/path/99052#获取应用权限详情
func (c *ApiClient) ExecGetPermissionsAgent(req ReqGetPermissionsAgent) (RespGetPermissionsAgent, error) {
	var resp RespGetPermissionsAgent
	err := c.executeWXApiPost("/cgi-bin/agent/get_permissions", req, &resp, true)
	if err != nil {
		return RespGetPermissionsAgent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetPermissionsAgent{}, bizErr
	}
	return resp, nil
}
