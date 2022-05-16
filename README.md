## 企业微信第三方服务商 Go SDK

Go语言实现企业微信sdk，a sensible Work Weixin SDK for Go。

以第三方服务商角度整理的sdk，集成了第三方应用sdk和自建应用代开发的sdk，支持一键生成新的sdk，使用简单，扩展灵活。

- 用缓存方案实现分布式 access_token/jsapi_ticket，保证在多个服务中只有一个服务能成功调用企微API请求 access_token/jsapi_ticket，减少API调用次数和服务重启需要重新获取的情况
- 用缓存方案实现读取/更新suite_ticket，保证多个服务能读取到最新的suite_ticket（suite_ticket每十分钟更新一次）
- 获取授权企业ApiClient时，支持自定义闭包从数据库等读取企业数据，eg: Sdk.GetThirdAuthCorpApiClient
- 提供Logger interface：用于自行实现日志记录器，便于收集日志

[点击查看企业微信第三方应用开发博文](https://zsmhub.github.io/post/%E5%AE%9E%E6%88%98%E6%A1%88%E4%BE%8B/%E4%BC%81%E4%B8%9A%E5%BE%AE%E4%BF%A1%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%94%E7%94%A8%E5%BC%80%E5%8F%91/)

### 一键生成sdk代码命令

> 注意：部分复杂的页面需要手动整理下sdk，如消息推送>发送应用消息接口。

- 生成api代码（tip: 生成GET方式的接口，请求参数的数据类型需要手动调整下）

    `make api doc=https://developer.work.weixin.qq.com/document/path/90600`

- 生成callback代码

    `make callback doc=https://developer.work.weixin.qq.com/document/path/92277`


### sdk调用示例

**强烈建议去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/workweixin/tree/master/demo)

#### 回调sdk调用示例

```go
// 企微回调设置初始化
func InitCallbackHandler() error {
    // 服务商回调解析
    if err := workweixin.Sdk.NewProviderCallbackHandler(config.CorpCallbackToken, config.CorpCallbackEncodingAESKey); err != nil {
        return err
    }

    // 第三方应用回调解析【可选】
    if err := workweixin.Sdk.NewThirdAppCallbackHandler(config.AppSuiteCallbackToken, config.AppSuiteCallbackEncodingAESKey); err != nil {
        return err
    }

    // 第三方小程序回调解析【可选】
    if err := workweixin.Sdk.NewMiniSuiteCaNewThirdMiniCallbackHandlerllbackHandler(config.MiniSuiteCallbackToken, config.MiniSuiteCallbackEncodingAESKey); err != nil {
        return err
    }

    // 自建应用代开发回调解析【可选】
    if err := workweixin.Sdk.NewCustomizedTemplateCallbackHandler(config.CustomizedCallbackToken, config.CustomizedCallbackEncodingAESKey); err != nil {
        return err
    }

    return nil
}

// 服务商-解析并获取回调信息
workweixin.Sdk.ProviderCallback.GetCallBackMsg(r *http.Request)

// 第三方应用-解析并获取回调信息
workweixin.Sdk.ThirdAppCallback.GetCallBackMsg(r *http.Request)

// 第三方小程序-解析并获取回调信息
workweixin.Sdk.ThirdMiniCallback.GetCallBackMsg(r *http.Request)

// 自建应用代开发--解析并获取回调信息
workweixin.Sdk.CustomizedTemplateCallback.GetCallBackMsg(r *http.Request)

// 第三方应用回调完整示例
func HandleAppPostRequest(c echo.Context) error {
    msg, err := workweixin.Sdk.ThirdAppCallback.GetCallBackMsg(c.Request())
    if err != nil {
        return err
    }

    switch msg.MsgType {

    case callbacks.MessageTypeThird: // 第三方应用回调
        switch msg.EventType {

            case callbacks.InfoTypeSuiteTicket: // 每十分钟推送一次suite_ticket
                extras, ok := msg.Extras.(callbacks.ThirdSuiteTicket)
                if !ok {
                    return errors.New("suite_ticket get failed")
                }

                ticket := extras.SuiteTicket.Text
                workweixin.Sdk.ThirdAppClient.RefreshSuiteTicket(ticket, time.Hour)

                // todo: 此处可将 suite_ticket 保存进数据库

                return nil


    }

    return nil
}
```

#### api sdk调用示例

```go
import "xxx/workweixin/demo"

// 企微API客户端初始化
func InitApiHandler() error {
    // 初始化企微sdk参数
    workweixin.Sdk.InitOptions(apis.Options{
        DcsToken:                     demo.DcsTokenByRedis{},
        DcsAppSuiteTicket:            demo.DcsAppSuiteTicketByRedis{},
        GetThirdAppAuthCorpFunc:      demo.GetThirdAppAuthCorpToSdk,
        GetCustomizedAppAuthCorpFunc: demo.GetCustomizedAppAuthCorpToSdk,
        Logger:                       demo.Logger{},
    })

    // 服务商API客户端初始化
    workweixin.Sdk.NewProviderApiClient(config.CorpId, config.CorpProviderSecret)

    // 第三方应用API客户端初始化【可选】
    suiteTicket := dao.ConfigDao.GetByUniqueIndex(global.ConfigKeySuiteTicket)
    workweixin.Sdk.NewThirdAppApiClient(config.CorpId, config.AppSuiteId, config.AppSuiteSecret, suiteTicket.V)

    // 自建应用代开发API客户端初始化【可选】
    customizedTicket := dao.ConfigDao.GetByUniqueIndex(global.ConfigKeyCustomizedTicket)
    workweixin.Sdk.NewCustomizedApiClient(config.CorpId, config.CustomizedAppSuiteId, config.CustomizedAppSuiteSecret, customizedTicket.V)

    // 由于本地开发环境和预发布无法接收企微回调事件，故需定时刷新suite_ticket
    if config.IsLocal() || config.IsUat() {
        go func(suiteTicket, customizedTicket model.Config) {
            defer recover()
            ticker := time.NewTicker(10 * time.Minute)
            defer ticker.Stop()
            for {
                // 更新第三方应用ticket
                _ = suiteTicket.DelCache() // 清除本地缓存
                suiteTicket = dao.ConfigDao.GetByUniqueIndex(global.ConfigKeySuiteTicket)
                workweixin.Sdk.ThirdAppClient.RefreshSuiteTicket(suiteTicket.V, 30*time.Minute)

                // 更新自建应用代开发ticket
                _ = customizedTicket.DelCache() // 清除本地缓存
                customizedTicket = dao.ConfigDao.GetByUniqueIndex(global.ConfigKeyCustomizedTicket)
                workweixin.Sdk.CustomizedAppClient.RefreshSuiteTicket(customizedTicket.V, 30*time.Minute)

                <-ticker.C
            }
        }(suiteTicket, customizedTicket)
    }

    return nil
}

// 获取企业永久授权码
resp, err := workweixin.Sdk.ThirdAppClient.ExecGetPermanentCodeService(apis.ReqGetPermanentCodeService{AuthCode: authCode})

// 企微 error code 类型强制转换
if err != nil {
    apiError, ok := err.(*apis.ClientError)
    if !ok {
         return nil, errors.New("转换失败，类型有误")
    }
    if apiError.Code == apis.ErrCode60011 {
        return nil, errors.New("无权限访问")
    }
    return nil, err
}

// 推送消息到第三方应用
apiClient, err := workweixin.Sdk.GetThirdAuthCorpApiClient(v.CorpId)
if err != nil {
    fmt.Println(err)
}
reqSentMessageCard := apis.ReqSentMessageCard{
    ToUser:  v.InstallUserId,
    MsgType: "news",
    AgentId: v.AgentId,
    News: apis.ReqSentMessageCardNewsBody{
        Articles: []apis.ReqSentMessageCardNewsArticleBody{
            {
                Title:       "新模块【xxx】已上线",
                Description: "快进入【管理后台】把它配置到你的【侧边栏】中！",
                Url:         "https://xxx",
                UrlImg:      "https://xxx/workbench-config.jpg",
            },
        },
    },
}
if _, err = apiClient.ExecSentMessageCard(reqSentMessageCard); err != nil {
    fmt.Println(err)
}
```

### 注意点

- 如果你发现了sdk中，没有某个回调事件或某个api，可以使用一键生成sdk代码命令生成，然后提交下pr