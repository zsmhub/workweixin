package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetApprovalDetailOa 获取审批申请详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
type ReqGetApprovalDetailOa struct {
	// SpNo 审批单编号。，必填
	SpNo int `json:"sp_no"`
}

var _ bodyer = ReqGetApprovalDetailOa{}

func (x ReqGetApprovalDetailOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetApprovalDetailOa 获取审批申请详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
type RespGetApprovalDetailOa struct {
	CommonResp
	Info struct {
		ApplyData struct {
			Contents []struct {
				Control string `json:"control"`
				ID      string `json:"id"`
				Title   []struct {
					Lang string `json:"lang"`
					// Text 文本内容，即申请人在此控件填写的文本内容
					Text string `json:"text"`
				} `json:"title"`
				Value struct {
					Children    []interface{} `json:"children"`    // 明细内容，一个明细控件可能包含多个子明细
					Departments []interface{} `json:"departments"` // 部门内容，即申请人在此控件选择的部门，多选模式下可能有多个
					Files       []interface{} `json:"files"`       // 文件内容，即申请人在此控件上传的文件内容，可能有多个
					Members     []interface{} `json:"members"`     // 成员内容，即申请人在此控件选择的成员，多选模式下可能有多个
					StatField   []interface{} `json:"stat_field"`
					// Text 文本内容，即申请人在此控件填写的文本内容
					Text string        `json:"text"`
					Tips []interface{} `json:"tips"`
				} `json:"value"`
			} `json:"contents"`
		} `json:"apply_data"`
		ApplyTime int `json:"apply_time"`
		Applyer   struct {
			Partyid string `json:"partyid"`
			Userid  string `json:"userid"`
		} `json:"applyer"`
		Comments []struct {
			CommentUserInfo struct {
				Userid string `json:"userid"`
			} `json:"commentUserInfo"`
			Commentcontent string   `json:"commentcontent"`
			Commentid      string   `json:"commentid"`
			Commenttime    int      `json:"commenttime"`
			MediaID        []string `json:"media_id"`
		} `json:"comments"`
		Notifyer []struct {
			Userid string `json:"userid"`
		} `json:"notifyer"`
		SpName   string `json:"sp_name"`
		SpNo     string `json:"sp_no"`
		SpRecord []struct {
			Approverattr int `json:"approverattr"`
			Details      []struct {
				Approver struct {
					Userid string `json:"userid"`
				} `json:"approver"`
				MediaID  []interface{} `json:"media_id"`
				SpStatus int           `json:"sp_status"`
				Speech   string        `json:"speech"`
				Sptime   int           `json:"sptime"`
			} `json:"details"`
			SpStatus int `json:"sp_status"`
		} `json:"sp_record"`
		SpStatus   int    `json:"sp_status"`
		TemplateID string `json:"template_id"`
	} `json:"info"`
}

var _ bodyer = RespGetApprovalDetailOa{}

func (x RespGetApprovalDetailOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetApprovalDetailOa 获取审批申请详情
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
func (c *ApiClient) ExecGetApprovalDetailOa(req ReqGetApprovalDetailOa) (RespGetApprovalDetailOa, error) {
	var resp RespGetApprovalDetailOa
	err := c.executeWXApiPost("/cgi-bin/oa/getapprovaldetail", req, &resp, true)
	if err != nil {
		return RespGetApprovalDetailOa{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetApprovalDetailOa{}, bizErr
	}
	return resp, nil
}
