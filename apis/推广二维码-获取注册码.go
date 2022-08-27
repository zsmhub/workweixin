package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetRegisterCodeService 获取注册码请求
// 文档：https://developer.work.weixin.qq.com/document/path/90581#获取注册码
type ReqGetRegisterCodeService struct {
	// AdminMobile 管理员手机号
	AdminMobile string `json:"admin_mobile"`
	// AdminName 管理员姓名
	AdminName string `json:"admin_name"`
	// CorpName 企业名称
	CorpName string `json:"corp_name"`
	// FollowUser 跟进人的userid，必须是服务商所在企业的成员。若配置该值，则由该注册码创建的企业，在服务商管理后台，该企业的报备记录会自动标注跟进人员为指定成员
	FollowUser string `json:"follow_user"`
	// State 用户自定义的状态值。只支持英文字母和数字，最长为128字节。若指定该参数， 接口 <a href="#11729/查询注册状态">查询注册状态</a> 及 <a href="#11729/注册完成回调事件">注册完成回调事件</a> 会相应返回该字段值
	State string `json:"state"`
	// TemplateID 推广包ID，最长为128个字节。在“<a href="https://open.work.weixin.qq.com/wwopen/developer">服务商管理端</a>-应用管理-推广二维码”，创建的推广码详情可查看。，必填
	TemplateID string `json:"template_id"`
}

var _ bodyer = ReqGetRegisterCodeService{}

func (x ReqGetRegisterCodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetRegisterCodeService 获取注册码响应
// 文档：https://developer.work.weixin.qq.com/document/path/90581#获取注册码
type RespGetRegisterCodeService struct {
	CommonResp
	ExpiresIn    int    `json:"expires_in"`
	RegisterCode string `json:"register_code"`
}

var _ bodyer = RespGetRegisterCodeService{}

func (x RespGetRegisterCodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetRegisterCodeService 获取注册码
// 文档：https://developer.work.weixin.qq.com/document/path/90581#获取注册码
func (c *ApiClient) ExecGetRegisterCodeService(req ReqGetRegisterCodeService) (RespGetRegisterCodeService, error) {
	var resp RespGetRegisterCodeService
	err := c.executeWXApiPost("/cgi-bin/service/get_register_code", req, &resp, true)
	if err != nil {
		return RespGetRegisterCodeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetRegisterCodeService{}, bizErr
	}
	return resp, nil
}
