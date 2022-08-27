package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetTemplateDetail 获取审批模板详情
// 文档：https://developer.work.weixin.qq.com/document/path/92631#获取审批模板详情
type ReqGetTemplateDetail struct {
	// 服务商审批模板的唯一标识id。可在“获取审批单据详情”、“审批状态变化回调通知”中获得，也可在服务商后台-应用管理-审批模板的模板编辑页面中获得。
	TemplateId string `json:"template_id"`
}

var _ bodyer = ReqGetTemplateDetail{}

func (x ReqGetTemplateDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetTemplateDetail 获取审批模板详情
// 文档：https://developer.work.weixin.qq.com/document/path/92631#获取审批模板详情
type Property struct { // 模板控件属性，包含了模板内控件的各种属性信息
	Control string `json:"control"` // 控件类型：Text-文本；Textarea-多行文本；Number-数字；Money-金额；Date-日期/日期+时间；Selector-单选/多选；Contact-成员/部门；Tips-说明文字；File-附件；Table-明细；Attendance-假勤控件；Vacation-请假控件
	Id      string `json:"id"`      // 控件id
	Title   []struct { // 控件名称，若配置了多语言则会包含中英文的控件名称，默认为zh_CN中文
		Text string `json:"text"`
		Lang string `json:"lang"`
	} `json:"title"`
	Placeholder []struct { // 控件说明，向申请者展示的控件填写说明，若配置了多语言则会包含中英文的控件说明，默认为zh_CN中文
		Text string `json:"text"`
		Lang string `json:"lang"`
	} `json:"placeholder"`
	Require int `json:"require"`
	UnPrint int `json:"un_print"`
}

type Date struct {
	Type string `json:"type"`
}

type Selector struct {
	Type    string `json:"type"`
	Options []struct {
		Key   string `json:"key"`
		Value []struct {
			Text string `json:"text"`
			Lang string `json:"lang"`
		} `json:"value"`
	} `json:"options"`
}

type Contact struct {
	Type string `json:"type"`
	Mode string `json:"mode"`
}

type Table struct {
	Children []struct {
		Property struct {
			Control string `json:"control"`
			Id      string `json:"id"`
			Title   []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"title"`
			Placeholder []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"placeholder"`
			Require int `json:"require"`
			UnPrint int `json:"un_print"`
		} `json:"property"`
	} `json:"children"`
	StatField []interface{} `json:"stat_field"`
}

type Attendance struct {
	DateRange struct {
		Type string `json:"type"`
	} `json:"date_range"`
	Type int `json:"type"`
}
type VacationList struct {
	Item []struct {
		Id   int `json:"id"`
		Name []struct {
			Text string `json:"text"`
			Lang string `json:"lang"`
		} `json:"name"`
	} `json:"item"`
}

type RespGetTemplateDetail struct {
	CommonResp
	TemplateNames []struct { // 模板名称，若配置了多语言则会包含中英文的模板名称，默认为zh_CN中文
		Text string `json:"text"`
		Lang string `json:"lang"`
	} `json:"template_names"`
	TemplateContent struct { // 模板控件信息
		Controls []struct {
			Property Property `json:"property"`
			Config   struct {
				Date         Date         `json:"date"`
				Selector     Selector     `json:"selector"`
				Contact      Contact      `json:"contact"`
				Table        Table        `json:"table"`
				Attendance   Attendance   `json:"attendance"`
				VacationList VacationList `json:"vacation_list"`
			} `json:"config"`
		} `json:"controls"`
	} `json:"template_content"`
}

var _ bodyer = RespGetTemplateDetail{}

func (x RespGetTemplateDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetTemplateDetail 获取审批模板详情
// 文档：https://developer.work.weixin.qq.com/document/path/92631#获取审批模板详情
func (c *ApiClient) ExecGetTemplateDetail(req ReqGetTemplateDetail) (RespGetTemplateDetail, error) {
	var resp RespGetTemplateDetail
	err := c.executeWXApiPost("/cgi-bin/oa/gettemplatedetail", req, &resp, true)
	if err != nil {
		return RespGetTemplateDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetTemplateDetail{}, bizErr
	}
	return resp, nil
}
