package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSimplelistDepartment 获取子部门ID列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/95406#获取子部门ID列表
type ReqSimplelistDepartment struct {
	// ID 部门id。获取指定部门及其下的子部门（以及子部门的子部门等等，递归）。 如果不填，默认获取全量组织架构
	ID int `json:"id"`
}

var _ urlValuer = ReqSimplelistDepartment{}

func (x ReqSimplelistDepartment) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespSimplelistDepartment 获取子部门ID列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/95406#获取子部门ID列表
type RespSimplelistDepartment struct {
	DepartmentID []struct {
		// ID 创建的部门id
		ID int `json:"id"`
		// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)。
		Order int `json:"order"`
		// Parentid 父部门id。根部门为1。
		Parentid int `json:"parentid"`
	} `json:"department_id"` // 部门列表数据。
	CommonResp
}

var _ bodyer = RespSimplelistDepartment{}

func (x RespSimplelistDepartment) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSimplelistDepartment 获取子部门ID列表
// 文档：https://developer.work.weixin.qq.com/document/path/95406#获取子部门ID列表
func (c *ApiClient) ExecSimplelistDepartment(req ReqSimplelistDepartment) (RespSimplelistDepartment, error) {
	var resp RespSimplelistDepartment
	err := c.executeWXApiGet("/cgi-bin/department/simplelist", req, &resp, true)
	if err != nil {
		return RespSimplelistDepartment{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSimplelistDepartment{}, bizErr
	}
	return resp, nil
}
