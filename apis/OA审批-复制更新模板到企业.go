package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCopyTemplateApproval 复制/更新模板到企业请求
// 文档：https://developer.work.weixin.qq.com/document/path/92630#复制/更新模板到企业
type ReqCopyTemplateApproval struct {
	// 服务商审批模板的唯一标识id。可在“获取审批单据详情”、“审批状态变化回调通知”中获得，也可在服务商后台-应用管理-审批模板的模板编辑页面中获得。
	OpenTemplateId string `json:"open_template_id"`
}

var _ bodyer = ReqCopyTemplateApproval{}

func (x ReqCopyTemplateApproval) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCopyTemplateApproval 复制/更新模板到企业响应
// 文档：https://developer.work.weixin.qq.com/document/path/92630#复制/更新模板到企业
type RespCopyTemplateApproval struct {
	CommonResp
	// 服务商审批模板的唯一标识id。可在“获取审批单据详情”、“审批状态变化回调通知”中获得，也可在服务商后台-应用管理-审批模板的模板编辑页面中获得。
	TemplateId string `json:"template_id"`
}

var _ bodyer = RespCopyTemplateApproval{}

func (x RespCopyTemplateApproval) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCopyTemplateApproval 复制/更新模板到企业
// 文档：https://developer.work.weixin.qq.com/document/path/92630#复制/更新模板到企业
func (c *ApiClient) ExecCopyTemplateApproval(req ReqCopyTemplateApproval) (RespCopyTemplateApproval, error) {
	var resp RespCopyTemplateApproval
	err := c.executeWXApiPost("/cgi-bin/oa/approval/copytemplate", req, &resp, true)
	if err != nil {
		return RespCopyTemplateApproval{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCopyTemplateApproval{}, bizErr
	}
	return resp, nil
}
