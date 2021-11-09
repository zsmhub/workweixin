package apis

import (
	"bytes"
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSentMessage 接口定义请求
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90372#接口定义

type ReqSentMessageCardTextBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}
type ReqSentMessageCardNewsArticleBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlImg      string `json:"picurl"`
}
type ReqSentMessageCardNewsBody struct {
	Articles []ReqSentMessageCardNewsArticleBody `json:"articles"`
}
type ReqSentMessageCard struct {
	ToUser                 string                     `json:"touser"`
	ToParty                string                     `json:"toparty"`
	ToTag                  string                     `json:"totag"`
	MsgType                string                     `json:"msgtype"`
	AgentId                int                        `json:"agentid"`
	EnableIdTrans          string                     `json:"enable_id_trans"`
	EnableDuplicateCheck   string                     `json:"enable_duplicate_check"`
	DuplicateCheckInterval string                     `json:"duplicate_check_interval"`
	Textcard               ReqSentMessageCardTextBody `json:"textcard"`
	News                   ReqSentMessageCardNewsBody `json:"news"`
}

var _ bodyer = ReqSentMessageCard{}

func (x ReqSentMessageCard) intoBody() ([]byte, error) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(x)
	if err != nil {
		panic(err)
	}
	return byteBuf.Bytes(), nil
}

// RespSentMessage 接口定义响应
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90372#接口定义

type RespSentMessageCard struct {
	CommonResp
}

// execSentMessage 接口定义
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90372#接口定义
func (c *ApiClient) ExecSentMessageCard(req ReqSentMessageCard) (RespSentMessageCard, error) {
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
