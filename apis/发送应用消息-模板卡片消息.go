package apis

import (
	"bytes"
	"encoding/json"
)

// 模板卡片消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

type (
	SentTemplateCardBody struct {
		CardType              string                              `json:"card_type"`                         // 模板卡片类型，文本通知型卡片填写 "text_notice"
		Source                *TemplateCardSource                 `json:"source,omitempty"`                  // 卡片来源样式信息，不需要来源样式可不填写
		ActionMenu            *TemplateCardActionMenu             `json:"action_menu,omitempty"`             // 卡片右上角更多操作按钮
		TaskId                string                              `json:"task_id,omitempty"`                 // 任务id，同一个应用任务id不能重复，只能由数字、字母和“_-@”组成，最长128字节，填了action_menu字段的话本字段必填
		MainTitle             *TemplateCardMainTitle              `json:"main_title,omitempty"`              // 一级标题
		QuoteArea             *TemplateCardQuoteArea              `json:"quote_area,omitempty"`              // 引用文献样式
		HorizontalContentList []TemplateCardHorizontalContentItem `json:"horizontal_content_list,omitempty"` // 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
		JumpList              []TemplateCardJumpListItem          `json:"jump_list,omitempty"`               // 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
		CardAction            *TemplateCardCardAction             `json:"card_action,omitempty"`             // 整体卡片的点击跳转事件

		EmphasisContent *TemplateCardEmphasisContent `json:"emphasis_content,omitempty"` // 关键数据样式
		SubTitleText    string                       `json:"sub_title_text,omitempty"`   // 二级普通文本，建议不超过160个字，（支持id转译）

		ImageTextArea       *TemplateCardImageTextArea        `json:"image_text_area,omitempty"`       // 左图右文样式，news_notice类型的卡片，card_image和image_text_area两者必填一个字段，不可都不填
		CardImage           *TemplateCardCardImage            `json:"card_image,omitempty"`            // 图片样式，news_notice类型的卡片，card_image和image_text_area两者必填一个字段，不可都不填
		VerticalContentList []TemplateCardVerticalContentItem `json:"vertical_content_list,omitempty"` // 卡片二级垂直内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4

		ButtonSelection *TemplateCardButtonSelection `json:"button_selection,omitempty"` // 下拉式的选择器
		ButtonList      []TemplateCardButtonListItem `json:"button_list,omitempty"`      // 按钮列表，列表长度不超过6

		Checkbox     *TemplateCardCheckbox     `json:"checkbox,omitempty"`      // 选择题样式
		SubmitButton *TemplateCardSubmitButton `json:"submit_button,omitempty"` // 提交按钮样式

		SelectList []TemplateCardSelectListItem `json:"select_list,omitempty"` // 下拉式的选择器列表，multiple_interaction类型的卡片该字段不可为空，一个消息最多支持 3 个选择器
	}

	TemplateCardSource struct {
		IconUrl   string `json:"icon_url,omitempty"` // 来源图片的url，来源图片的尺寸建议为72*72
		Desc      string `json:"desc,omitempty"`     // 来源图片的描述，建议不超过20个字，（支持id转译）
		DescColor int    `json:"desc_color"`         // 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
	}

	TemplateCardActionMenu struct {
		Desc       string                   `json:"desc,omitempty"`        // 更多操作界面的描述
		ActionList []TemplateCardActionItem `json:"action_list,omitempty"` // 操作列表，列表长度取值范围为 [1, 3]
	}
	TemplateCardActionItem struct {
		Text string `json:"text"` // 操作的描述文案
		Key  string `json:"key"`  // 操作key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复
	}

	TemplateCardMainTitle struct {
		Title string `json:"title,omitempty"` // 一级标题，建议不超过36个字，文本通知型卡片本字段非必填，但不可本字段和sub_title_text都不填，（支持id转译）
		Desc  string `json:"desc,omitempty"`  // 标题辅助信息，建议不超过44个字，（支持id转译）
	}

	TemplateCardQuoteArea struct {
		Type      int    `json:"type"`                 // 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
		Url       string `json:"url,omitempty"`        // 点击跳转的url，quote_area.type是1时必填
		Appid     string `json:"appid,omitempty"`      // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，quote_area.type是2时必填
		Pagepath  string `json:"pagepath,omitempty"`   // 点击跳转的小程序的pagepath，quote_area.type是2时选填
		Title     string `json:"title,omitempty"`      // 引用文献样式的标题
		QuoteText string `json:"quote_text,omitempty"` // 引用文献样式的引用文案
	}

	TemplateCardEmphasisContent struct {
		Title string `json:"title,omitempty"` // 关键数据样式的数据内容，建议不超过14个字
		Desc  string `json:"desc,omitempty"`  // 关键数据样式的数据描述内容，建议不超过22个字
	}

	TemplateCardHorizontalContentItem struct {
		Type    int    `json:"type"`               // 链接类型，0或不填代表不是链接，1 代表跳转url，2 代表下载附件，3 代表点击跳转成员详情
		Keyname string `json:"keyname"`            // 二级标题，建议不超过5个字
		Value   string `json:"value,omitempty"`    // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过30个字，（支持id转译）
		Url     string `json:"url,omitempty"`      // 链接跳转的url，horizontal_content_list.type是1时必填
		MediaId string `json:"media_id,omitempty"` // 附件的media_id，horizontal_content_list.type是2时必填
		Userid  string `json:"userid,omitempty"`   // 成员详情的userid，horizontal_content_list.type是3时必填
	}

	TemplateCardJumpListItem struct {
		Type     int    `json:"type"`               // 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
		Title    string `json:"title"`              // 跳转链接样式的文案内容，建议不超过18个字
		Url      string `json:"url,omitempty"`      // 跳转链接的url，jump_list.type是1时必填
		Appid    string `json:"appid,omitempty"`    // 跳转链接的小程序的appid，必须是与当前应用关联的小程序，jump_list.type是2时必填
		Pagepath string `json:"pagepath,omitempty"` // 跳转链接的小程序的pagepath，jump_list.type是2时选填
	}

	TemplateCardCardAction struct {
		Type     int    `json:"type,omitempty"`     // 跳转事件类型，1 代表跳转url，2 代表打开小程序。text_notice卡片模版中该字段取值范围为[1,2]
		Url      string `json:"url,omitempty"`      // 跳转事件的url，card_action.type是1时必填
		Appid    string `json:"appid,omitempty"`    // 跳转事件的小程序的appid，必须是与当前应用关联的小程序，card_action.type是2时必填
		Pagepath string `json:"pagepath,omitempty"` // 跳转事件的小程序的pagepath，card_action.type是2时选填
	}

	TemplateCardImageTextArea struct {
		Type      int    `json:"type"`                 // 左图右文样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
		Url       string `json:"url,omitempty"`        // 点击跳转的url，image_text_area.type是1时必填
		Title     string `json:"title,omitempty"`      // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，image_text_area.type是2时必填
		Desc      string `json:"desc,omitempty"`       // 左图右文样式的描述
		ImagesUrl string `json:"images_url,omitempty"` // 左图右文样式的图片url
	}

	TemplateCardCardImage struct {
		Url         string  `json:"url"`                    // 图片的url
		AspectRatio float64 `json:"aspect_ratio,omitempty"` // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
	}

	TemplateCardVerticalContentItem struct {
		Title string `json:"title"`          // 卡片二级标题，建议不超过38个字
		Desc  string `json:"desc,omitempty"` // 二级普通文本，建议不超过160个字
	}

	TemplateCardButtonSelection struct {
		QuestionKey string                       `json:"question_key"`          // 下拉式的选择器的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
		Title       string                       `json:"title,omitempty"`       // 下拉式的选择器左边的标题
		OptionList  []TemplateCardOptionListItem `json:"option_list"`           // 选项列表，下拉选项不超过 10 个，最少1个
		SelectedId  string                       `json:"selected_id,omitempty"` // 默认选定的id，不填或错填默认第一个
	}
	TemplateCardOptionListItem struct {
		Id   string `json:"id"`   // 下拉式的选择器选项的id，用户提交后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
		Text string `json:"text"` // 下拉式的选择器选项的文案，建议不超过16个字
	}

	TemplateCardButtonListItem struct {
		Type  int    `json:"type"`          // 按钮点击事件类型，0 或不填代表回调点击事件，1 代表跳转url
		Text  string `json:"text"`          // 按钮文案，建议不超过10个字
		Style int    `json:"style"`         // 按钮样式，目前可填1~4，不填或错填默认1
		Key   string `json:"key,omitempty"` // 按钮key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复，button_list.type是0时必填
		Url   string `json:"url,omitempty"` // 跳转事件的url，button_list.type是1时必填
	}

	TemplateCardCheckbox struct {
		QuestionKey string                           `json:"question_key"` // 选择题key值，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
		OptionList  []TemplateCardCheckboxOptionItem `json:"option_list"`  // 选项list，选项个数不超过 20 个，最少1个
		Mode        int                              `json:"mode"`         // 选择题模式，单选：0，多选：1，不填默认0
	}
	TemplateCardCheckboxOptionItem struct {
		Id        string `json:"id"`         // 选项id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
		Text      string `json:"text"`       // 选项文案描述，建议不超过17个字
		IsChecked bool   `json:"is_checked"` // 该选项是否要默认选中
	}
	TemplateCardSubmitButton struct {
		Text string `json:"text"` // 按钮文案，建议不超过10个字，不填默认为提交
		Key  string `json:"key"`  // 提交按钮的key，会产生回调事件将本参数作为EventKey返回，最长支持1024字节
	}

	TemplateCardSelectListItem struct {
		QuestionKey string                             `json:"question_key"`          // 下拉式的选择器题目的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节，不可重复
		Title       string                             `json:"title,omitempty"`       // 下拉式的选择器上面的title
		SelectedId  string                             `json:"selected_id,omitempty"` // 默认选定的id，不填或错填默认第一个
		OptionList  []TemplateCardSelectListOptionItem `json:"option_list"`           // 选项列表，下拉选项不超过 10 个，最少1个
	}
	TemplateCardSelectListOptionItem struct {
		Id   string `json:"id"`   // 下拉式的选择器选项的id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
		Text string `json:"text"` // 下拉式的选择器选项的文案，建议不超过16个字
	}
)

type ReqSentMessageCardTemplateCard struct {
	ToUser                 string               `json:"touser,omitempty"`                   // 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToParty                string               `json:"toparty,omitempty"`                  // 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag                  string               `json:"totag,omitempty"`                    // 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	MsgType                string               `json:"msgtype"`                            // 消息类型，如text/image/file
	AgentId                int                  `json:"agentid"`                            // 企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
	EnableIdTrans          int                  `json:"enable_id_trans,omitempty"`          // 表示是否开启id转译，0表示否，1表示是，默认0
	EnableDuplicateCheck   int                  `json:"enable_duplicate_check,omitempty"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval int                  `json:"duplicate_check_interval,omitempty"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	TemplateCard           SentTemplateCardBody `json:"template_card"`
}

var _ bodyer = ReqSentMessageCardTemplateCard{}

func (x ReqSentMessageCardTemplateCard) intoBody() ([]byte, error) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(x)
	if err != nil {
		panic(err)
	}
	body := byteBuf.Bytes()
	return body, nil
}

// 发送应用消息-模板卡片消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义
func (c *ApiClient) ExecSentMessageCardTemplateCard(req ReqSentMessageCardTemplateCard) (RespSentMessageCard, error) {
	var resp RespSentMessageCard
	err := c.executeWXApiPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return RespSentMessageCard{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSentMessageCard{}, bizErr
	}
	return resp, nil
}
