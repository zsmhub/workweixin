package demo

import (
    "fmt"
    "github.com/zsmhub/workweixin"
    "github.com/zsmhub/workweixin/apis"
)

// 调用企微sdk示例

func ApiMain() {
    if err := initApiHandler(); err != nil {
        fmt.Println(err)
    }

    // 调用企业授权信息sdk
    resp, err := workweixin.Sdk.ThirdAppClient.ExecGetAuthInfoService(apis.ReqGetAuthInfoService{
        AuthCorpid:    AuthCorpId,
        PermanentCode: AuthPermanentCode,
    })
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(resp)
}

// 企微API客户端初始化
func initApiHandler() error {
    // 服务商API客户端初始化
    workweixin.Sdk.NewProviderApiClient(CorpId, CorpProviderSecret)

    // 第三方应用API客户端初始化
    workweixin.Sdk.NewThirdAppApiClient(CorpId, AppSuiteId, AppSuiteSecret, SuiteTicket)

    // 授权企业API客户端初始化（此段代码也可以考虑注释掉，由用户主动触发）
    // if err := workweixin.Sdk.NewAuthCorpApiClient(AuthCorpId, AuthPermanentCode, AuthAgentId); err != nil {
    // 	return err
    // }

    return nil
}
