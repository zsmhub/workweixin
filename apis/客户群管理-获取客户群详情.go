package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqGetGroupchat 获取客户群详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/92707#获取客户群详情
type ReqGetGroupchat struct {
	// ChatID 客户群ID，必填
	ChatID string `json:"chat_id"`
	// NeedName 是否需要返回群成员的名字<code>group_chat.member_list.name</code>。0-不返回；1-返回。默认不返回
	NeedName int `json:"need_name"`
}

var _ bodyer = ReqGetGroupchat{}

func (x ReqGetGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetGroupchat 获取客户群详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/92707#获取客户群详情
type RespGetGroupchat struct {
	CommonResp
	GroupChat struct {
		AdminList []struct {
			Userid string `json:"userid"`
		} `json:"admin_list"`
		ChatID     string `json:"chat_id"`
		CreateTime int    `json:"create_time"`
		MemberList []struct {
			GroupNickname string `json:"group_nickname"`
			Invitor       struct {
				Userid string `json:"userid"`
			} `json:"invitor"`
			JoinScene int    `json:"join_scene"`
			JoinTime  int    `json:"join_time"`
			Name      string `json:"name"`
			Type      int    `json:"type"`
			Unionid   string `json:"unionid"`
			Userid    string `json:"userid"`
		} `json:"member_list"`
		Name   string `json:"name"`
		Notice string `json:"notice"`
		Owner  string `json:"owner"`
	} `json:"group_chat"`
}

var _ bodyer = RespGetGroupchat{}

func (x RespGetGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetGroupchat 获取客户群详情
// 文档：https://developer.work.weixin.qq.com/document/path/92707#获取客户群详情
func (c *ApiClient) ExecGetGroupchat(req ReqGetGroupchat) (RespGetGroupchat, error) {
	var resp RespGetGroupchat
	err := c.executeWXApiPost("/cgi-bin/externalcontact/groupchat/get", req, &resp, true)
	if err != nil {
		return RespGetGroupchat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetGroupchat{}, bizErr
	}
	return resp, nil
}
