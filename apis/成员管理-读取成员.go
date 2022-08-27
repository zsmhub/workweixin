package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetUser 读取成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/90332#读取成员
type ReqGetUser struct {
	// Userid 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节，必填
	Userid string `json:"userid"`
}

var _ urlValuer = ReqGetUser{}

func (x ReqGetUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetUser 读取成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/90332#读取成员
type RespGetUser struct {
	Address    string `json:"address"`
	Alias      string `json:"alias"`
	Avatar     string `json:"avatar"`
	Department []int  `json:"department"`
	Email      string `json:"email"`
	CommonResp
	Extattr struct {
		Attrs []struct {
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Type int `json:"type"`
			Web  struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"web"`
		} `json:"attrs"`
	} `json:"extattr"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalAttr []struct {
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Type int `json:"type"`
			Web  struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"web"`
		} `json:"external_attr"`
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status"`
		} `json:"wechat_channels"`
	} `json:"external_profile"`
	Gender         string `json:"gender"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	MainDepartment int    `json:"main_department"`
	Mobile         string `json:"mobile"`
	Name           string `json:"name"`
	OpenUserid     string `json:"open_userid"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	QrCode         string `json:"qr_code"`
	Status         int    `json:"status"`
	Telephone      string `json:"telephone"`
	ThumbAvatar    string `json:"thumb_avatar"`
	Userid         string `json:"userid"`
}

var _ bodyer = RespGetUser{}

func (x RespGetUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUser 读取成员
// 文档：https://developer.work.weixin.qq.com/document/path/90332#读取成员
func (c *ApiClient) ExecGetUser(req ReqGetUser) (RespGetUser, error) {
	var resp RespGetUser
	err := c.executeWXApiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return RespGetUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUser{}, bizErr
	}
	return resp, nil
}
