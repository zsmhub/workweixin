package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

type DateValue struct { // 日期/日期+时间控件（control参数为Date）
	Type       string `json:"type"`        // 时间展示类型：day-日期；hour-日期+时间 ，和对应模板控件属性一致
	STimestamp string `json:"s_timestamp"` // 时间戳-字符串类型，在此填写日期/日期+时间控件的选择值，以此为准
}
type SelectorValue struct { // 单选/多选控件（control参数为Selector）
	Type    string     `json:"type"` // 选择方式：single-单选；multi-多选
	Options []struct { // 多选选项，多选属性的选择控件允许输入多个
		Key string `json:"key"` // 选项key，可通过“获取审批模板详情”接口获得
	} `json:"options"`
}
type MembersValue struct { // 成员控件（control参数为Contact，且value参数为members）
	Userid string `json:"userid"` // 所选成员的userid
	Name   string `json:"name"`   // 成员名
}
type DepartmentsValue struct { // 部门控件（control参数为Contact，且value参数为departments）
	OpenapiId string `json:"openapi_id"` // 所选部门id
	Name      string `json:"name"`       // 所选部门名
}
type FilesValue struct { // 附件控件（control参数为File，且value参数为files）
	FileId string `json:"file_id"` // 文件id，该id为临时素材上传接口返回的的media_id，注：提单后将作为单据内容转换为长期文件存储；目前一个审批申请单，全局仅支持上传6个附件，否则将失败。
}
type TableValue struct { // 明细控件（control参数为Table）
	List []TableValueList `json:"list"` // 子明细列表，在此填写子明细的所有子控件的值，子控件的数据结构同一般控件
}

type TableValueList struct { // 子明细列表，在此填写子明细的所有子控件的值，子控件的数据结构同一般控件
	Control string `json:"control"`
	Id      string `json:"id"`
	Value   struct {
		Text      string `json:"text"`
		NewNumber string `json:"new_number"`
		NewMoney  string `json:"new_money"`
	} `json:"value"`
}
type VacationValue struct { // 假勤组件-请假组件（control参数为Vacation）
	Selector struct { // 请假类型，所选选项与假期管理关联，为假期管理中的假期类型
		Type    string     `json:"type"` // 选择方式：single-单选；multi-多选，在假勤控件中固定为单选
		Options []struct { // 用户所选选项
			Key   string     `json:"key"` // 选项key，选项的唯一id，可通过“获取审批模板详情”接口获得vacation_list中item的id值
			Value []struct { // 选项值，若配置了多语言则会包含中英文的选项值
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"value"`
		} `json:"options"`
		ExpType int `json:"exp_type"`
	} `json:"selector"`
	Attendance struct { // 假勤组件
		DateRange struct { // 假勤组件时间选择范围
			Type        string `json:"type"`         // 时间展示类型：day-日期；hour-日期+时间
			NewBegin    int    `json:"new_begin"`    // 开始时间戳
			NewEnd      int    `json:"new_end"`      // 结束时间戳
			NewDuration int    `json:"new_duration"` // 请假时长，单位秒
		} `json:"date_range"`
		Type int `json:"type"` // 假勤组件类型：1-请假；3-出差；4-外出；5-加班
	} `json:"attendance"`
}
type AttendanceValue struct { // 假勤组件-出差/外出/加班组件（control参数为Attendance）
	DateRange struct { // 假勤组件时间选择范围
		Type        string `json:"type"`         // 时间展示类型：day-日期；hour-日期+时间
		NewBegin    int    `json:"new_begin"`    // 开始时间戳
		NewEnd      int    `json:"new_end"`      // 结束时间戳
		NewDuration int    `json:"new_duration"` // 请假时长，单位秒
	} `json:"date_range"`
	Type int `json:"type"` // 假勤组件类型：1-请假；3-出差；4-外出；5-加班
}

type ApplyDataContent struct { // 审批申请详情，由多个表单控件及其内容组成，其中包含需要对控件赋值的信息
	Control string `json:"control"` // 控件类型：Text-文本；Textarea-多行文本；Number-数字；Money-金额；Date-日期/日期+时间；Selector-单选/多选；；Contact-成员/部门；Tips-说明文字；File-附件；Table-明细；
	Id      string `json:"id"`      // 控件id：控件的唯一id，可通过“获取审批模板详情”接口获取
	Value   struct {
		Text             string             `json:"text"`
		NewNumber        string             `json:"new_number"`
		NewMoney         string             `json:"new_money"`
		DateValue        DateValue          `json:"date"`
		SelectorValue    DateValue          `json:"selector"`
		MembersValue     []MembersValue     `json:"members"`
		DepartmentsValue []DepartmentsValue `json:"departments"`
		FilesValue       []FilesValue       `json:"files"`
		TableValue       []TableValue       `json:"children"`
		VacationValue    VacationValue      `json:"vacation"`
		AttendanceValue  AttendanceValue    `json:"attendance"`
	} `json:"value"`
}

type SummaryList struct { // 摘要信息，用于显示在审批通知卡片、审批列表的摘要信息，最多3行
	SummaryInfo []SummaryInfo `json:"summary_info"`
}
type SummaryInfo struct { // 摘要行信息，用于定义某一行摘要显示的内容
	Text string `json:"text"` // 摘要行显示文字，用于记录列表和消息通知的显示，不要超过20个字符
	Lang string `json:"lang"` // 摘要行显示语言，中文：zh_CN（注意不是zh-CN），英文：en。
}

type ApplyData struct { // 审批申请数据，可定义审批申请中各个控件的值，其中必填项必须有值，选填项可为空，数据结构同“获取审批申请详情”接口返回值中同名参数“apply_data”
	Contents []ApplyDataContent `json:"contents"` // 审批申请详情，由多个表单控件及其内容组成，其中包含需要对控件赋值的信息
}

type ApplyEventProcess struct {
	NodeList []ApplyEventProcessNodeList `json:"node_list"`
}
type ApplyEventProcessNodeList struct {
	Type   int      `json:"type"`
	ApvRel int      `json:"apv_rel"`
	Userid []string `json:"userid"`
}

// ReqApplyEvent 提交审批申请
// 文档：https://developer.work.weixin.qq.com/document/path/92632#提交审批申请
type ReqApplyEvent struct {
	CreatorUserid       string             `json:"creator_userid"`              // 申请人userid，此审批申请将以此员工身份提交，申请人需在应用可见范围内
	TemplateId          string             `json:"template_id"`                 // 模板的唯一标识id。可在“获取审批单据详情”、“审批状态变化回调通知”中获得，也可在使用“复制/更新模板到企业”接口回调中获得。注：此id为企业内模板的实例id，非服务商后台对应模板的id。暂不支持通过接口提交[打卡补卡][调班]模板审批单。
	UseTemplateApprover int                `json:"use_template_approver"`       // 审批人模式：0-通过接口指定审批人、抄送人（此时approver、notifyer等参数可用）; 1-使用此模板在管理后台设置的审批流程，支持条件审批。
	ChooseDepartment    int                `json:"choose_department,omitempty"` // 提单者提单部门id，不填默认为主部门
	ApplyData           ApplyData          `json:"apply_data"`                  // 审批申请数据，可定义审批申请中各个控件的值，其中必填项必须有值，选填项可为空，数据结构同“获取审批申请详情”接口返回值中同名参数“apply_data”
	SummaryList         []SummaryList      `json:"summary_list"`                // 摘要信息，用于显示在审批通知卡片、审批列表的摘要信息，最多3行
	Process             *ApplyEventProcess `json:"process,omitempty"`           // 新版流程列表，use_template_approver = 0 时必填
}

var _ bodyer = ReqApplyEvent{}

func (x ReqApplyEvent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespApplyEvent 提交审批申请
// 文档：https://developer.work.weixin.qq.com/document/path/92632#提交审批申请
type RespApplyEvent struct {
	CommonResp
	SpNo string `json:"sp_no"` // 表单提交成功后，返回的表单编号
}

var _ bodyer = RespApplyEvent{}

func (x RespApplyEvent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecApplyEvent 提交审批申请
// 文档：https://developer.work.weixin.qq.com/document/path/92632#提交审批申请
func (c *ApiClient) ExecApplyEvent(req ReqApplyEvent) (RespApplyEvent, error) {
	var resp RespApplyEvent
	err := c.executeWXApiPost("/cgi-bin/oa/applyevent", req, &resp, true)
	if err != nil {
		return RespApplyEvent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespApplyEvent{}, bizErr
	}
	return resp, nil
}
