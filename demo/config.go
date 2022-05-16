package demo

// 企微配置文件

// 企业微信服务商信息
var (
	CorpId                     = "xxx" // 通用开发参数 >> CorpId
	CorpProviderSecret         = "xxx" // 通用开发参数 >> ProviderSecret
	CorpCallbackToken          = "xxx" // 通用开发参数 >> Token
	CorpCallbackEncodingAESKey = "xxx" // 通用开发参数 >> EncodingAESKey
)

// 企业微信第三方应用配置
var (
	AppSuiteId                     = "xxx" // 唯一身份标识
	AppSuiteSecret                 = "xxx" // 调用身份密钥
	AppSuiteCallbackToken          = "xxx" // 回调 token
	AppSuiteCallbackEncodingAESKey = "xxx" // 回调加解密 AES KEY
)

// 企业微信第三方小程序配置
var (
	MiniSuiteId                     = "xxx" // 唯一身份标识
	MiniSuiteSecret                 = "xxx" // 调用身份密钥
	MiniSuiteCallbackToken          = "xxx" // 回调 token
	MiniSuiteCallbackEncodingAESKey = "xxx" // 回调加解密 AES KEY
)

// 企业微信代开发配置
var (
	CustomizedAppSuiteId             = "xxx" // 代开发唯一身份标识
	CustomizedAppSuiteSecret         = "xxx" // 代开发调用身份密钥
	CustomizedCallbackToken          = "xxx" // 代开发回调 token
	CustomizedCallbackEncodingAESKey = "xxx" // 代开发回调加解密 AES KEY
)

// 调试数据配置
var (
	TestSuiteTicket           = "xxx"   // suite_ticket，10分钟刷新一次
	TestAuthCorpId            = "xxx"   // 授权企业corpid
	TestAuthCorpPermanentCode = "xxx"   // 授权企业永久授权码
	TestAuthCorpAgentId       = 1000001 // 授权企业agentid
)
