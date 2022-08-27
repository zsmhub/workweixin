package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListDepartment 获取部门列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/90344#获取部门列表
type ReqListDepartment struct {
	// ID 部门id。获取指定部门及其下的子部门（以及子部门的子部门等等，递归）。 如果不填，默认获取全量组织架构
	ID int `json:"id"`
}

var _ urlValuer = ReqListDepartment{}

func (x ReqListDepartment) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListDepartment 获取部门列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/90344#获取部门列表
type RespListDepartment struct {
	Department []struct {
		// DepartmentLeader 部门负责人的UserID；第三方仅通讯录应用可获取
		DepartmentLeader []string `json:"department_leader"`
		// ID 创建的部门id
		ID int `json:"id"`
		// Name 部门名称，代开发自建应用需要管理员授权才返回；此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回name，返回的name字段使用id代替，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示部门名称
		Name string `json:"name"`
		// NameEn 英文名称，此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回该字段
		NameEn string `json:"name_en"`
		// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
		Order int `json:"order"`
		// Parentid 父部门id。根部门为1
		Parentid int `json:"parentid"`
	} `json:"department"` // 部门列表数据。
	CommonResp
}

var _ bodyer = RespListDepartment{}

func (x RespListDepartment) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListDepartment 获取部门列表
// 文档：https://developer.work.weixin.qq.com/document/path/90344#获取部门列表
func (c *ApiClient) ExecListDepartment(req ReqListDepartment) (RespListDepartment, error) {
	var resp RespListDepartment
	err := c.executeWXApiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return RespListDepartment{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListDepartment{}, bizErr
	}
	return resp, nil
}
