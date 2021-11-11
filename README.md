## 企业微信第三方应用开发 go sdk

Go语言实现企业微信API，a sensible Work Weixin SDK for Go。

[点击查看企业微信第三方应用开发博文](https://zsmhub.github.io/post/%E5%AE%9E%E6%88%98%E6%A1%88%E4%BE%8B/%E4%BC%81%E4%B8%9A%E5%BE%AE%E4%BF%A1%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%94%E7%94%A8%E5%BC%80%E5%8F%91/)

### 自动生成sdk代码命令

`支持手动生成企业微信新API或新回调代码`

- 生成api代码
    
    `make api doc=https://open.work.weixin.qq.com/api/doc/90001/90143/90600`

- 生成callback代码

    `make callback doc=https://open.work.weixin.qq.com/api/doc/90001/90143/92277`


### sdk 调用示例

#### 回调sdk调用示例

```go
// 企微回调设置初始化
func InitCallbackHandler() error {
    // 服务商回调解析
    if err := workweixin.Sdk.NewProviderCallbackHandler(config.CorpCallbackToken, config.CorpCallbackEncodingAESKey); err != nil {
        return err
    }

    // 第三方应用回调解析
    if err := workweixin.Sdk.NewAppSuiteCallbackHandler(config.AppSuiteCallbackToken, config.AppSuiteCallbackEncodingAESKey); err != nil {
        return err
    }

    // 第三方小程序回调解析【可选】
    if err := workweixin.Sdk.NewMiniSuiteCallbackHandler(config.MiniSuiteCallbackToken, config.MiniSuiteCallbackEncodingAESKey); err != nil {
        return err
    }

    return nil
}

// 服务商-解析并获取回调信息
workweixin.Sdk.ProviderCallback.GetCallBackMsg(r *http.Request)

// 第三方应用-解析并获取回调信息
workweixin.Sdk.AppSuiteCallback.GetCallBackMsg(r *http.Request)

// 第三方小程序-解析并获取回调信息
workweixin.Sdk.MiniSuiteCallback.GetCallBackMsg(r *http.Request)

// 第三方应用回调完整示例
func HandleAppPostRequest(c echo.Context) error {
    msg, err := workweixin.Sdk.AppSuiteCallback.GetCallBackMsg(c.Request())
    if err != nil {
        return err
    }

    switch msg.MsgType {

    case callbacks.MessageTypeThird: // 第三方应用回调
        switch msg.EventType {
        
            case callbacks.InfoTypeSuiteTicket: // 每十分钟推送一次suite_ticket
                ticket := msg.Extras.(callbacks.ThirdSuiteTicket).SuiteTicket.Text
                workweixin.Sdk.ThirdAppClient.RefreshSuiteTicket(ticket)

                retryer := backoff.WithContext(backoff.NewExponentialBackOff(), context.Background())
                err := backoff.Retry(func() error {
                    // todo 将 suite_ticket 存储在数据库或其他地方
                    return nil
                }, retryer)

                return err

    }

    return nil
}
```

#### api sdk调用示例

```go
// 企微API客户端初始化
func InitApiHandler() error {
    // 服务商API客户端初始化
    workweixin.Sdk.NewProviderApiClient(config.CorpId, config.CorpProviderSecret)

    // 第三方应用API客户端初始化
    suiteTicket := "xxx" // 从数据库等地方获取已得到的suite_ticket
    workweixin.Sdk.NewThirdAppApiClient(config.CorpId, config.AppSuiteId, config.AppSuiteSecret, suiteTicket)

    // 授权企业API客户端初始化
    authCorpList := xxx.GetAuthCorpList() // 从数据库获取已授权企业
    for _, corp := range authCorpList {
        workweixin.Sdk.NewAuthCorpApiClient(corp.CorpId, corp.PermanentCode, workweixin.Sdk.ThirdAppClient)
    }

    return nil
}

// 获取企业永久授权码
resp, err := workweixin.Sdk.ThirdAppClient.ExecGetPermanentCodeService(apis.ReqGetPermanentCodeService{AuthCode: authCode})

// 企微 error code 处理
if err != nil {
    apiError, _ := err.(*apis.ClientError)
    if apiError.Code == apis.ErrCode60011 {
        return nil, errors.New("无权限访问")
    }
    return nil, err
}

// 推送消息到第三方应用
apiClient, err := workweixin.Sdk.GetAuthCorpAPPClient(v.CorpId)
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

### 如果部署到K8S多个副本

1. 可以将 access_token 的刷新代码抽离出来，用一个定时任务单独处理，并存在缓存中，提供给多个K8S副本读取 access_token。
2. 项目运行中，会有新的企业安装我们的第三方应用，此时多个K8S副本需要同步新的企业数据并实例化，避免 sdk 调用失败。