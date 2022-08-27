package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSetoneuserquotaVacation 修改成员假期余额请求
// 文档：https://developer.work.weixin.qq.com/document/path/94213#修改成员假期余额
type ReqSetoneuserquotaVacation struct {
	Leftduration int    `json:"leftduration"`
	Remarks      string `json:"remarks"`
	TimeAttr     int    `json:"time_attr"`
	Userid       string `json:"userid"`
	VacationID   int    `json:"vacation_id"`
}

var _ bodyer = ReqSetoneuserquotaVacation{}

func (x ReqSetoneuserquotaVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSetoneuserquotaVacation 修改成员假期余额响应
// 文档：https://developer.work.weixin.qq.com/document/path/94213#修改成员假期余额
type RespSetoneuserquotaVacation struct {
	CommonResp
}

var _ bodyer = RespSetoneuserquotaVacation{}

func (x RespSetoneuserquotaVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSetoneuserquotaVacation 修改成员假期余额
// 文档：https://developer.work.weixin.qq.com/document/path/94213#修改成员假期余额
func (c *ApiClient) ExecSetoneuserquotaVacation(req ReqSetoneuserquotaVacation) (RespSetoneuserquotaVacation, error) {
	var resp RespSetoneuserquotaVacation
	err := c.executeWXApiPost("/cgi-bin/oa/vacation/setoneuserquota", req, &resp, true)
	if err != nil {
		return RespSetoneuserquotaVacation{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSetoneuserquotaVacation{}, bizErr
	}
	return resp, nil
}
