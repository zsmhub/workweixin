package demo

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/zsmhub/workweixin"
    "github.com/zsmhub/workweixin/callbacks"
)

// 企微回调示例

func CallbackMain() {
    if err := CallbackRepo.InitCallbackHandler(); err != nil {
        fmt.Println(err)
        return
    }

    e := echo.New()

    // 第三方服务商回调，系统事件接收
    e.GET("/callback/corp", func(c echo.Context) error {
        workweixin.Sdk.ProviderCallback.EchoTestHandler(c.Response().Writer, c.Request())
        return nil
    })
    e.POST("/callback/corp", func(c echo.Context) error {
        _ = CallbackRepo.HandleCorpPostRequest(c)
        return c.String(200, "success")
    })

    // 第三方应用回调，数据回调和指令回调
    e.GET("/callback/app", func(c echo.Context) error {
        workweixin.Sdk.AppSuiteCallback.EchoTestHandler(c.Response().Writer, c.Request())
        return nil
    })
    e.POST("/callback/app", func(c echo.Context) error {
        _ = CallbackRepo.HandleAppPostRequest(c)
        return c.String(200, "success")
    })

    e.Logger.Fatal(e.Start(":1323"))
}

type callbackRepo struct{}

var CallbackRepo = new(callbackRepo)

// 企微回调设置初始化
func (callbackRepo) InitCallbackHandler() error {
    // 服务商回调解析
    if err := workweixin.Sdk.NewProviderCallbackHandler(CorpCallbackToken, CorpCallbackEncodingAESKey); err != nil {
        return err
    }

    // 第三方应用回调解析
    if err := workweixin.Sdk.NewAppSuiteCallbackHandler(AppSuiteCallbackToken, AppSuiteCallbackEncodingAESKey); err != nil {
        return err
    }

    return nil
}

// 服务商回调，系统事件接收
func (callbackRepo) HandleCorpPostRequest(c echo.Context) error {
    msg, err := workweixin.Sdk.ProviderCallback.GetCallBackMsg(c.Request())
    if err != nil {
        return err
    }

    go func() {
        defer recover()

        switch msg.MsgType {

        case callbacks.MessageTypeThird: // 第三方应用回调
            switch msg.EventType {

            case callbacks.InfoTypeAgreeExternalUseridMigration:
                corpId := msg.AuthCorpId
                fmt.Printf("同意授权转换external_userid事件：%s\n", corpId)

            }

        }
    }()

    return nil
}

// 第三方应用回调，数据回调和指令回调
func (r callbackRepo) HandleAppPostRequest(c echo.Context) error {
    msg, err := workweixin.Sdk.AppSuiteCallback.GetCallBackMsg(c.Request())
    if err != nil {
        return err
    }

    go func() {
        defer recover()

        switch msg.MsgType {

        case callbacks.MessageTypeThird: // 第三方应用回调
            r.handleMessageTypeThirdEvent(msg)

        }
    }()

    return nil
}

func (callbackRepo) handleMessageTypeThirdEvent(msg callbacks.CallbackMessage) {
    switch msg.EventType {

    case callbacks.InfoTypeSuiteTicket: // 每十分钟推送一次suite_ticket
        extras, ok := msg.Extras.(callbacks.ThirdSuiteTicket)
        if !ok {
            return
        }
        ticket := extras.SuiteTicket.Text

        // todo 如果是用k8s部署，可使用广播更新方案通知其他pod

        fmt.Printf("收到suite_ticked: ", ticket)

    }
}