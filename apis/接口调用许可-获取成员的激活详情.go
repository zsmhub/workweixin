package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetActiveInfoByUserLicense 获取成员的激活详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/95555#获取成员的激活详情
type ReqGetActiveInfoByUserLicense struct {
	// Corpid 企业corpid，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
	// Userid 待查询员工的userid，只支持加密的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetActiveInfoByUserLicense{}

func (x ReqGetActiveInfoByUserLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetActiveInfoByUserLicense 获取成员的激活详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/95555#获取成员的激活详情
type RespGetActiveInfoByUserLicense struct {
	ActiveInfoList []struct {
		// ActiveCode 帐号激活码
		ActiveCode string `json:"active_code"`
		// ActiveTime 激活时间
		ActiveTime int `json:"active_time"`
		// ExpireTime 过期时间
		ExpireTime int `json:"expire_time"`
		// Type 帐号类型:1:基础帐号，2:互通帐号
		Type int `json:"type"`
		// Userid 帐号绑定激活的成员userid。返回加密的userid
		Userid string `json:"userid"`
	} `json:"active_info_list"` // 帐号列表，同一个userid最多有两个帐号:一个基础帐号与一个互通帐号
	// ActiveStatus 帐号激活状态。0:未激活、 1:已激活
	ActiveStatus int `json:"active_status"`
	CommonResp
}

var _ bodyer = RespGetActiveInfoByUserLicense{}

func (x RespGetActiveInfoByUserLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetActiveInfoByUserLicense 获取成员的激活详情
// 文档：https://developer.work.weixin.qq.com/document/path/95555#获取成员的激活详情
func (c *ApiClient) ExecGetActiveInfoByUserLicense(req ReqGetActiveInfoByUserLicense) (RespGetActiveInfoByUserLicense, error) {
	var resp RespGetActiveInfoByUserLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_active_info_by_user", req, &resp, true)
	if err != nil {
		return RespGetActiveInfoByUserLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetActiveInfoByUserLicense{}, bizErr
	}
	return resp, nil
}
