package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAdminListService 获取应用的管理员列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/90606#获取应用的管理员列表
type ReqGetAdminListService struct {
	// Agentid 授权方安装的应用agentid，必填
	Agentid int `json:"agentid"`
	// AuthCorpid 授权方corpid，必填
	AuthCorpid string `json:"auth_corpid"`
}

var _ bodyer = ReqGetAdminListService{}

func (x ReqGetAdminListService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetAdminListService 获取应用的管理员列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/90606#获取应用的管理员列表
type RespGetAdminListService struct {
	Admin []struct {
		AuthType   int    `json:"auth_type"`
		OpenUserid string `json:"open_userid"`
		Userid     string `json:"userid"`
	} `json:"admin"`
	CommonResp
}

var _ bodyer = RespGetAdminListService{}

func (x RespGetAdminListService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAdminListService 获取应用的管理员列表
// 文档：https://developer.work.weixin.qq.com/document/path/90606#获取应用的管理员列表
func (c *ApiClient) ExecGetAdminListService(req ReqGetAdminListService) (RespGetAdminListService, error) {
	var resp RespGetAdminListService
	err := c.executeWXApiPost("/cgi-bin/service/get_admin_list", req, &resp, true)
	if err != nil {
		return RespGetAdminListService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAdminListService{}, bizErr
	}
	return resp, nil
}
