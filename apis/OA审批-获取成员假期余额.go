package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuservacationquotaVacation 获取成员假期余额请求
// 文档：https://developer.work.weixin.qq.com/document/path/94212#获取成员假期余额

type ReqGetuservacationquotaVacation struct {
	Userid string `json:"userid"`
}

var _ bodyer = ReqGetuservacationquotaVacation{}

func (x ReqGetuservacationquotaVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetuservacationquotaVacation 获取成员假期余额响应
// 文档：https://developer.work.weixin.qq.com/document/path/94212#获取成员假期余额

type RespGetuservacationquotaVacation struct {
	CommonResp
	Lists []struct {
		Assignduration int    `json:"assignduration"`
		ID             int    `json:"id"`
		Leftduration   int    `json:"leftduration"`
		Usedduration   int    `json:"usedduration"`
		Vacationname   string `json:"vacationname"`
	} `json:"lists"`
}

var _ bodyer = RespGetuservacationquotaVacation{}

func (x RespGetuservacationquotaVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// execGetuservacationquotaVacation 获取成员假期余额
// 文档：https://developer.work.weixin.qq.com/document/path/94212#获取成员假期余额
func (c *ApiClient) ExecGetuservacationquotaVacation(req ReqGetuservacationquotaVacation) (RespGetuservacationquotaVacation, error) {
	var resp RespGetuservacationquotaVacation
	err := c.executeWXApiPost("/cgi-bin/oa/vacation/getuservacationquota", req, &resp, true)
	if err != nil {
		return RespGetuservacationquotaVacation{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuservacationquotaVacation{}, bizErr
	}
	return resp, nil
}
