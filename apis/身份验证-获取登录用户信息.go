package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetLoginInfoService 获取登录用户信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/91125#获取登录用户信息
type ReqGetLoginInfoService struct {
	// AuthCode oauth2.0授权企业微信管理员登录产生的code，最长为512字节。只能使用一次，5分钟未被使用自动过期，必填
	AuthCode string `json:"auth_code"`
}

var _ bodyer = ReqGetLoginInfoService{}

func (x ReqGetLoginInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetLoginInfoService 获取登录用户信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/91125#获取登录用户信息
type RespGetLoginInfoService struct {
	CommonResp
	Usertype int `json:"usertype"`
	UserInfo struct {
		Userid     string `json:"userid"`
		OpenUserid string `json:"open_userid"`
		Name       string `json:"name"`
		Avatar     string `json:"avatar"`
	} `json:"user_info"`
	CorpInfo struct {
		Corpid string `json:"corpid"`
	} `json:"corp_info"`
	Agent []struct {
		Agentid  int `json:"agentid"`
		AuthType int `json:"auth_type"`
	} `json:"agent"`
	AuthInfo struct {
		Department []struct {
			ID       int  `json:"id"`
			Writable bool `json:"writable"`
		} `json:"department"`
	} `json:"auth_info"`
}

var _ bodyer = RespGetLoginInfoService{}

func (x RespGetLoginInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetLoginInfoService 获取登录用户信息
// 文档：https://developer.work.weixin.qq.com/document/path/91125#获取登录用户信息
func (c *ApiClient) ExecGetLoginInfoService(req ReqGetLoginInfoService) (RespGetLoginInfoService, error) {
	var resp RespGetLoginInfoService
	err := c.executeWXApiPost("/cgi-bin/service/get_login_info", req, &resp, true)
	if err != nil {
		return RespGetLoginInfoService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetLoginInfoService{}, bizErr
	}
	return resp, nil
}
