package apis

// 推广二维码方式安装应用：获取应用安装链接
// 文档：https://open.work.weixin.qq.com/api/doc/90001/90143/90578
func (c *ApiClient) ExecGetRegisterUrl(registerCode string) string {
	return "https://open.work.weixin.qq.com/3rdservice/wework/register?register_code=" + registerCode
}
