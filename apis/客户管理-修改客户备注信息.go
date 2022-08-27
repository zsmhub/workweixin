package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqRemarkExternalcontact 修改客户备注信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/92694#修改客户备注信息
type ReqRemarkExternalcontact struct {
	// Description 此用户对外部联系人的描述，最多150个字符
	Description string `json:"description,omitempty"`
	// ExternalUserid 外部联系人userid，必填
	ExternalUserid string `json:"external_userid"`
	// Remark 此用户对外部联系人的备注，最多20个字符
	Remark string `json:"remark,omitempty"`
	// RemarkCompany 此用户对外部联系人备注的所属公司名称，最多20个字符
	RemarkCompany string `json:"remark_company,omitempty"`
	// RemarkMobiles 此用户对外部联系人备注的手机号
	RemarkMobiles []string `json:"remark_mobiles,omitempty"`
	// RemarkPicMediaid 备注图片的mediaid，
	RemarkPicMediaid string `json:"remark_pic_mediaid,omitempty"`
	// Userid 企业成员的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqRemarkExternalcontact{}

func (x ReqRemarkExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespRemarkExternalcontact 修改客户备注信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/92694#修改客户备注信息
type RespRemarkExternalcontact struct {
	CommonResp
}

var _ bodyer = RespRemarkExternalcontact{}

func (x RespRemarkExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecRemarkExternalcontact 修改客户备注信息
// 文档：https://developer.work.weixin.qq.com/document/path/92694#修改客户备注信息
func (c *ApiClient) ExecRemarkExternalcontact(req ReqRemarkExternalcontact) (RespRemarkExternalcontact, error) {
	var resp RespRemarkExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/remark", req, &resp, true)
	if err != nil {
		return RespRemarkExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespRemarkExternalcontact{}, bizErr
	}
	return resp, nil
}
