package apis

import (
    "encoding/json"

    "fmt"
    "net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqGetExternalcontact 请求
// 文档：https://open.work.weixin.qq.com/api/doc/90001/90143/92265#获取客户详情

type ReqGetExternalcontact struct {
    // Cursor 上次请求返回的<code>next_cursor</code>
    Cursor string `json:"cursor"`
    // ExternalUserid 外部联系人的<code>userid</code>，注意不是企业成员的帐号，必填
    ExternalUserid string `json:"external_userid"`
}

var _ urlValuer = ReqGetExternalcontact{}

func (x ReqGetExternalcontact) intoURLValues() url.Values {
    var ret url.Values = make(map[string][]string)

    var vals map[string]interface{}
    jsonBytes, _ := json.Marshal(x)
    _ = json.Unmarshal(jsonBytes, &vals)

    for k, v := range vals {
        ret.Add(k, fmt.Sprintf("%v", v))
    }
    return ret
}

// RespGetExternalcontact 响应
// 文档：https://open.work.weixin.qq.com/api/doc/90001/90143/92265#获取客户详情

type RespGetExternalcontact struct {
    CommonResp
    ExternalContact struct {
        Avatar          string `json:"avatar"`
        CorpFullName    string `json:"corp_full_name"`
        CorpName        string `json:"corp_name"`
        ExternalProfile struct {
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
        } `json:"external_profile"`
        ExternalUserid string `json:"external_userid"`
        Gender         int    `json:"gender"`
        Name           string `json:"name"`
        Position       string `json:"position"`
        Type           int    `json:"type"`
        Unionid        string `json:"unionid"`
    } `json:"external_contact"`
    FollowUser []struct {
        AddWay         int      `json:"add_way"`
        Createtime     int      `json:"createtime"`
        Description    string   `json:"description"`
        OperUserid     string   `json:"oper_userid"`
        Remark         string   `json:"remark"`
        RemarkCorpName string   `json:"remark_corp_name"`
        RemarkMobiles  []string `json:"remark_mobiles"`
        State          string   `json:"state"`
        Tags           []struct {
            GroupName string `json:"group_name"`
            TagID     string `json:"tag_id"`
            TagName   string `json:"tag_name"`
            Type      int    `json:"type"`
        } `json:"tags"`
        Userid string `json:"userid"`
    } `json:"follow_user"`
    NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetExternalcontact{}

func (x RespGetExternalcontact) intoBody() ([]byte, error) {
    result, err := json.Marshal(x)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// execGetExternalcontact 
// 文档：https://open.work.weixin.qq.com/api/doc/90001/90143/92265#获取客户详情
func (c *ApiClient) ExecGetExternalcontact(req ReqGetExternalcontact) (RespGetExternalcontact, error) {
    var resp RespGetExternalcontact
    err := c.executeWXApiGet("/cgi-bin/externalcontact/get", req, &resp, true)
    if err != nil {
        return RespGetExternalcontact{}, err
    }
    if bizErr := resp.TryIntoErr(); bizErr != nil {
        return RespGetExternalcontact{}, bizErr
    }

    return resp, nil
}
