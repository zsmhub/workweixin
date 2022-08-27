package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAppQrcodeService 获取应用二维码请求
// 文档：https://developer.work.weixin.qq.com/document/path/95430#获取应用二维码
type ReqGetAppQrcodeService struct {
	// Appid 第三方应用id，单应用不需要该参数，多应用旧套件才需要传该参数。若不传默认为1
	Appid int `json:"appid,omitempty"`
	// ResultType 结果返回方式，默认为返回二维码图片buffer。1:二维码图片buffer，2:二维码图片url
	ResultType int `json:"result_type,omitempty"`
	// State state值，用于区分不同的安装渠道，可以填写a-zA-Z0-9，长度不可超过32个字节，默认为空。扫应用带参二维码授权安装后，获取企业永久授权码接口会返回该state值
	State string `json:"state,omitempty"`
	// Style 二维码样式选项，默认为不带说明外框小尺寸。0:带说明外框的二维码，适合于实体物料，1:带说明外框的二维码，适合于屏幕类，2:不带说明外框（小尺寸），3:不带说明外框（中尺寸），4:不带说明外框（大尺寸）。具体样式与服务商管理端获取到的应用二维码样式一一对应，参见下文二维码样式说明
	Style int `json:"style,omitempty"`
	// SuiteID 第三方应用id（即ww或wx开头的suiteid），必填
	SuiteID string `json:"suite_id"`
}

var _ bodyer = ReqGetAppQrcodeService{}

func (x ReqGetAppQrcodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetAppQrcodeService 获取应用二维码响应
// 文档：https://developer.work.weixin.qq.com/document/path/95430#获取应用二维码
type RespGetAppQrcodeService struct {
	CommonResp
	Qrcode string `json:"qrcode"`
}

var _ bodyer = RespGetAppQrcodeService{}

func (x RespGetAppQrcodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAppQrcodeService 获取应用二维码
// 文档：https://developer.work.weixin.qq.com/document/path/95430#获取应用二维码
func (c *ApiClient) ExecGetAppQrcodeService(req ReqGetAppQrcodeService) (RespGetAppQrcodeService, error) {
	req.ResultType = 2 // 结果返回方式，固定传2，返回二维码图片url
	var resp RespGetAppQrcodeService
	err := c.executeWXApiPost("/cgi-bin/service/get_app_qrcode", req, &resp, true)
	if err != nil {
		return RespGetAppQrcodeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAppQrcodeService{}, bizErr
	}
	return resp, nil
}
