package demo

// 企微配置文件

// 企微第三方服务商配置
var (
    CorpId                     = "xxx" // 通用开发参数 >> CorpId
    CorpProviderSecret         = "xxx" // 通用开发参数 >> ProviderSecret
    CorpCallbackToken          = "xxx" // 通用开发参数 >> Token
    CorpCallbackEncodingAESKey = "xxx" // 通用开发参数 >> EncodingAESKey
)

// 企微第三方应用配置
var (
    AppSuiteId                     = "xxx" // 唯一身份标识
    AppSuiteSecret                 = "xxx" // 调用身份密钥
    AppSuiteCallbackToken          = "xxx" // 回调 token
    AppSuiteCallbackEncodingAESKey = "xxx" // 回调加解密 AES KEY
)

var (
    SuiteTicket       = "xxx"   // suite_ticket，10分钟刷新一次x
    AuthCorpId        = "xxx"   // 授权企业corpid
    AuthPermanentCode = "xxx"   // 授权企业永久授权码
    AuthAgentId       = 1000001 // 授权企业agentid
)