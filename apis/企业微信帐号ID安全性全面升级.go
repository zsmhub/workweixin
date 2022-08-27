package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCorpidToOpencorpidService 1.4 corpid转换请求
// 文档：https://developer.work.weixin.qq.com/document/path/95327#1.4 corpid转换
type ReqCorpidToOpencorpidService struct {
	// Corpid 待获取的企业ID，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqCorpidToOpencorpidService{}

func (x ReqCorpidToOpencorpidService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCorpidToOpencorpidService 1.4 corpid转换响应
// 文档：https://developer.work.weixin.qq.com/document/path/95327#1.4 corpid转换
type RespCorpidToOpencorpidService struct {
	CommonResp
	OpenCorpid string `json:"open_corpid"`
}

var _ bodyer = RespCorpidToOpencorpidService{}

func (x RespCorpidToOpencorpidService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCorpidToOpencorpidService 1.4 corpid转换
// 文档：https://developer.work.weixin.qq.com/document/path/95327#1.4 corpid转换
func (c *ApiClient) ExecCorpidToOpencorpidService(req ReqCorpidToOpencorpidService) (RespCorpidToOpencorpidService, error) {
	var resp RespCorpidToOpencorpidService
	err := c.executeWXApiPost("/cgi-bin/service/corpid_to_opencorpid", req, &resp, true)
	if err != nil {
		return RespCorpidToOpencorpidService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCorpidToOpencorpidService{}, bizErr
	}
	return resp, nil
}

// ReqUseridToOpenuseridBatch 2.4 userid的转换请求
// 文档：https://developer.work.weixin.qq.com/document/path/95327#2.4 userid的转换
type ReqUseridToOpenuseridBatch struct {
	// UseridList 获取到的成员ID，必填
	UseridList []string `json:"userid_list"`
}

var _ bodyer = ReqUseridToOpenuseridBatch{}

func (x ReqUseridToOpenuseridBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUseridToOpenuseridBatch 2.4 userid的转换响应
// 文档：https://developer.work.weixin.qq.com/document/path/95327#2.4 userid的转换
type RespUseridToOpenuseridBatch struct {
	CommonResp
	OpenUseridList []struct {
		Userid     string `json:"userid"`
		OpenUserid string `json:"open_userid"`
	} `json:"open_userid_list"`
	InvalidUseridList []string `json:"invalid_userid_list"`
}

var _ bodyer = RespUseridToOpenuseridBatch{}

func (x RespUseridToOpenuseridBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUseridToOpenuseridBatch 2.4 userid的转换
// 文档：https://developer.work.weixin.qq.com/document/path/95327#2.4 userid的转换
func (c *ApiClient) ExecUseridToOpenuseridBatch(req ReqUseridToOpenuseridBatch) (RespUseridToOpenuseridBatch, error) {
	var resp RespUseridToOpenuseridBatch
	err := c.executeWXApiPost("/cgi-bin/batch/userid_to_openuserid", req, &resp, true)
	if err != nil {
		return RespUseridToOpenuseridBatch{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUseridToOpenuseridBatch{}, bizErr
	}
	return resp, nil
}

// 3.3 转换external_userid
// 文档：https://developer.work.weixin.qq.com/document/path/95327
type ReqGetNewExternalUseridExternalcontact struct {
	// 旧外部联系人id列表，最多不超过1000个，必填
	ExternalUseridList []string `json:"external_userid_list"`
}

var _ bodyer = ReqGetNewExternalUseridExternalcontact{}

func (x ReqGetNewExternalUseridExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetNewExternalUseridExternalcontact struct {
	CommonResp
	Items []struct {
		ExternalUserid    string `json:"external_userid"`
		NewExternalUserid string `json:"new_external_userid"`
	} `json:"items"`
}

var _ bodyer = RespGetNewExternalUseridExternalcontact{}

func (x RespGetNewExternalUseridExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetNewExternalUseridExternalcontact(req ReqGetNewExternalUseridExternalcontact) (RespGetNewExternalUseridExternalcontact, error) {
	var resp RespGetNewExternalUseridExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_new_external_userid", req, &resp, true)
	if err != nil {
		return RespGetNewExternalUseridExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetNewExternalUseridExternalcontact{}, bizErr
	}
	return resp, nil
}

// 3.3 设置迁移完成
// 文档：https://developer.work.weixin.qq.com/document/path/95327
type ReqFinishExternalUseridMigration struct {
	// 企业corpid，必填
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqFinishExternalUseridMigration{}

func (x ReqFinishExternalUseridMigration) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFinishExternalUseridMigration struct {
	CommonResp
	Items []struct {
		ExternalUserid    string `json:"external_userid"`
		NewExternalUserid string `json:"new_external_userid"`
	} `json:"items"`
}

var _ bodyer = RespFinishExternalUseridMigration{}

func (x RespFinishExternalUseridMigration) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFinishExternalUseridMigration(req ReqFinishExternalUseridMigration) (RespFinishExternalUseridMigration, error) {
	var resp RespFinishExternalUseridMigration
	err := c.executeWXApiPost("/cgi-bin/service/externalcontact/finish_external_userid_migration", req, &resp, true)
	if err != nil {
		return RespFinishExternalUseridMigration{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFinishExternalUseridMigration{}, bizErr
	}
	return resp, nil
}

// ReqUnionidToExternalUserid3RdExternalcontact 4.2 unionid查询external_userid请求
// 文档：https://developer.work.weixin.qq.com/document/path/95327#4.2 unionid查询external_userid
type ReqUnionidToExternalUserid3RdExternalcontact struct {
	// 微信用户的unionid，必填
	Unionid string `json:"unionid"`
	// 微信用户的openid，必填
	Openid string `json:"openid"`
	// 需要换取的企业corpid，不填则拉取所有企业
	Corpid string `json:"corpid"`
}

var _ bodyer = ReqUnionidToExternalUserid3RdExternalcontact{}

func (x ReqUnionidToExternalUserid3RdExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUnionidToExternalUserid3RdExternalcontact 4.2 unionid查询external_userid响应
// 文档：https://developer.work.weixin.qq.com/document/path/95327#4.2 unionid查询external_userid
type RespUnionidToExternalUserid3RdExternalcontact struct {
	CommonResp
	ExternalUseridInfo []struct {
		Corpid         string `json:"corpid"`
		ExternalUserid string `json:"external_userid"`
	} `json:"external_userid_info"`
}

var _ bodyer = RespUnionidToExternalUserid3RdExternalcontact{}

func (x RespUnionidToExternalUserid3RdExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUnionidToExternalUserid3RdExternalcontact 4.2 unionid查询external_userid
// 文档：https://developer.work.weixin.qq.com/document/path/95327#4.2 unionid查询external_userid
func (c *ApiClient) ExecUnionidToExternalUserid3RdExternalcontact(req ReqUnionidToExternalUserid3RdExternalcontact) (RespUnionidToExternalUserid3RdExternalcontact, error) {
	var resp RespUnionidToExternalUserid3RdExternalcontact
	err := c.executeWXApiPost("/cgi-bin/service/externalcontact/unionid_to_external_userid_3rd", req, &resp, true)
	if err != nil {
		return RespUnionidToExternalUserid3RdExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUnionidToExternalUserid3RdExternalcontact{}, bizErr
	}
	return resp, nil
}

// 自建应用代开发external_userid转换
// 文档：https://developer.work.weixin.qq.com/document/path/95195
type ReqExternalcontactToServiceExternalUserid struct {
	// 明文externalUserId
	ExternalUserId string `json:"external_userid"`
}

var _ bodyer = ReqExternalcontactToServiceExternalUserid{}

func (x ReqExternalcontactToServiceExternalUserid) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespExternalcontactToServiceExternalUserid struct {
	CommonResp
	// 密文externalUserId
	ExternalUserid string `json:"external_userid"`
}

var _ bodyer = RespExternalcontactToServiceExternalUserid{}

func (x RespExternalcontactToServiceExternalUserid) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文档：https://developer.work.weixin.qq.com/document/path/95195
func (c *ApiClient) ExternalcontactToServiceExternalUserid(req ReqExternalcontactToServiceExternalUserid) (RespExternalcontactToServiceExternalUserid, error) {
	var resp RespExternalcontactToServiceExternalUserid
	err := c.executeWXApiPost("/cgi-bin/externalcontact/to_service_external_userid", req, &resp, true)
	if err != nil {
		return RespExternalcontactToServiceExternalUserid{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespExternalcontactToServiceExternalUserid{}, bizErr
	}
	return resp, nil
}
