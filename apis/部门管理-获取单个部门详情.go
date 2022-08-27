package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetDepartment 获取单个部门详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/95407#获取单个部门详情
type ReqGetDepartment struct {
	// ID 部门id，必填
	ID int `json:"id"`
}

var _ urlValuer = ReqGetDepartment{}

func (x ReqGetDepartment) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetDepartment 获取单个部门详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/95407#获取单个部门详情
type RespGetDepartment struct {
	Department struct {
		// DepartmentLeader 部门负责人的UserID，返回在应用可见范围内的部门负责人列表；第三方仅通讯录应用或者授权了“组织架构信息-应用可获取企业的部门组织架构信息-部门负责人”的第三方应用可获取
		DepartmentLeader []string `json:"department_leader"`
		// ID 部门id
		ID int `json:"id"`
		// Name 部门名称，代开发自建应用需要管理员授权才返回；第三方不可获取，需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示部门名称
		Name string `json:"name"`
		// NameEn 部门英文名称，代开发自建应用需要管理员授权才返回；第三方不可获取，需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示部门名称
		NameEn string `json:"name_en"`
		// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
		Order int `json:"order"`
		// Parentid 父部门id。根部门为1。
		Parentid int `json:"parentid"`
	} `json:"department"` // 部门详情。
	CommonResp
}

var _ bodyer = RespGetDepartment{}

func (x RespGetDepartment) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetDepartment 获取单个部门详情
// 文档：https://developer.work.weixin.qq.com/document/path/95407#获取单个部门详情
func (c *ApiClient) ExecGetDepartment(req ReqGetDepartment) (RespGetDepartment, error) {
	var resp RespGetDepartment
	err := c.executeWXApiGet("/cgi-bin/department/get", req, &resp, true)
	if err != nil {
		return RespGetDepartment{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetDepartment{}, bizErr
	}
	return resp, nil
}
