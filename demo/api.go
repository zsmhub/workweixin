package demo

import (
	"fmt"

	"github.com/lovego/workweixin"
	"github.com/lovego/workweixin/apis"
)

// 调用企微sdk示例
func ApiMain() {
	if err := initApiHandler(); err != nil {
		fmt.Println(err)
	}

	// 调用企业授权信息sdk
	resp, err := workweixin.Sdk.ThirdAppClient.ExecGetAuthInfoService(apis.ReqGetAuthInfoService{
		AuthCorpid:    TestAuthCorpId,
		PermanentCode: TestAuthCorpPermanentCode,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

// 企微API客户端初始化
func initApiHandler() error {
	// 初始化企微sdk参数
	workweixin.Sdk.InitOptions(apis.Options{
		DcsToken:          DcsTokenByRedis{},
		DcsAppSuiteTicket: DcsAppSuiteTicketByRedis{},
		// GetThirdAppAuthCorpFunc:      GetThirdAppAuthCorpToSdk,
		// GetCustomizedAppAuthCorpFunc: GetCustomizedAppAuthCorpToSdk,
		Logger: Logger{},
	})

	// 服务商API客户端初始化
	workweixin.Sdk.NewProviderApiClient(CorpId, CorpProviderSecret)

	// 第三方应用API客户端初始化
	suiteTicket := TestSuiteTicket
	workweixin.Sdk.NewThirdAppApiClient(CorpId, AppSuiteId, AppSuiteSecret, suiteTicket)

	// 自建应用代开发API客户端初始化
	// customizedTicket := "xxx"
	// workweixin.Sdk.NewCustomizedApiClient(CorpId, CustomizedAppSuiteId, CustomizedAppSuiteSecret, customizedTicket)

	return nil
}
