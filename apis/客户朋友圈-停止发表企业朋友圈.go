package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求


// ReqCancelMomentTaskExternalcontact 停止发表企业朋友圈请求
// 文档：https://developer.work.weixin.qq.com/document/path/97612#停止发表企业朋友圈
type ReqCancelMomentTaskExternalcontact struct {
	// MomentID 朋友圈id，可通过<a href="#25254/%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8" rel="nofollow">获取客户朋友圈企业发表的列表</a>接口获取朋友圈企业发表的列表，必填
	MomentID string `json:"moment_id"`
}


var _ urlValuer = ReqCancelMomentTaskExternalcontact{}

func (x ReqCancelMomentTaskExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}


// RespCancelMomentTaskExternalcontact 停止发表企业朋友圈响应
// 文档：https://developer.work.weixin.qq.com/document/path/97612#停止发表企业朋友圈
type RespCancelMomentTaskExternalcontact struct {
	CommonResp
}

var _ bodyer = RespCancelMomentTaskExternalcontact{}

func (x RespCancelMomentTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCancelMomentTaskExternalcontact 停止发表企业朋友圈
// 文档：https://developer.work.weixin.qq.com/document/path/97612#停止发表企业朋友圈
func (c *ApiClient) ExecCancelMomentTaskExternalcontact(req ReqCancelMomentTaskExternalcontact) (RespCancelMomentTaskExternalcontact, error) {
	var resp RespCancelMomentTaskExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/cancel_moment_task", req, &resp, true)
	if err != nil {
		return RespCancelMomentTaskExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCancelMomentTaskExternalcontact{}, bizErr
	}
	return resp, nil
}
