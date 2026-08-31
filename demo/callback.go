package demo

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lovego/workweixin"
	"github.com/lovego/workweixin/callbacks"
)

// 企微回调事件示例
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
		workweixin.Sdk.ThirdAppCallback.EchoTestHandler(c.Response().Writer, c.Request())
		return nil
	})
	e.POST("/callback/app", func(c echo.Context) error {
		_ = CallbackRepo.HandleAppPostRequest(c)
		return c.String(200, "success")
	})

	// 第三方小程序回调，数据回调和指令回调
	e.GET("/callback/mini", func(c echo.Context) error {
		workweixin.Sdk.ThirdMiniCallback.EchoTestHandler(c.Response().Writer, c.Request())
		return nil
	})
	e.POST("/callback/mini", func(c echo.Context) error {
		_ = CallbackRepo.HandleMiniPostRequest(c)
		return c.String(200, "success")
	})

	// 自建应用代开发回调，数据回调和指令回调
	e.GET("/callback/customized", func(c echo.Context) error {
		workweixin.Sdk.CustomizedTemplateCallback.EchoTestHandler(c.Response().Writer, c.Request())
		return nil
	})
	e.POST("/callback/customized", func(c echo.Context) error {
		_ = CallbackRepo.HandleCustomizedPostRequest(c)
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
	if err := workweixin.Sdk.NewThirdAppCallbackHandler(AppSuiteCallbackToken, AppSuiteCallbackEncodingAESKey); err != nil {
		return err
	}

	// 第三方小程序回调解析
	if err := workweixin.Sdk.NewThirdMiniCallbackHandler(MiniSuiteCallbackToken, MiniSuiteCallbackEncodingAESKey); err != nil {
		return err
	}

	// 自建应用代开发回调解析
	if err := workweixin.Sdk.NewCustomizedTemplateCallbackHandler(CustomizedCallbackToken, CustomizedCallbackEncodingAESKey); err != nil {
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
		defer func() {
			_ = recover()
		}()

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
	msg, err := workweixin.Sdk.ThirdAppCallback.GetCallBackMsg(c.Request())
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = recover()
		}()

		switch msg.MsgType {

		case callbacks.MessageTypeThird: // 第三方应用回调
			r.handleMessageTypeThirdEvent(msg)

		}
	}()

	return nil
}

// 第三方小程序回调，数据回调和指令回调
func (r callbackRepo) HandleMiniPostRequest(c echo.Context) error {
	msg, err := workweixin.Sdk.ThirdMiniCallback.GetCallBackMsg(c.Request())
	if err != nil {
		return err
	}

	fmt.Printf("%v", msg)

	return nil
}

// 自建应用代开发回调，数据回调和指令回调
func (r callbackRepo) HandleCustomizedPostRequest(c echo.Context) error {
	msg, err := workweixin.Sdk.CustomizedTemplateCallback.GetCallBackMsg(c.Request())
	if err != nil {
		return err
	}

	fmt.Printf("%v", msg)

	return nil
}

func (r callbackRepo) handleMessageTypeThirdEvent(msg callbacks.CallbackMessage) {
	switch msg.EventType {

	case callbacks.InfoTypeSuiteTicket: // 每十分钟推送一次suite_ticket
		extras, ok := msg.Extras.(callbacks.ThirdSuiteTicket)
		if !ok {
			fmt.Println("suite_ticket get failed")
			return
		}

		ticket := extras.SuiteTicket.Text
		workweixin.Sdk.ThirdAppClient.RefreshSuiteTicket(ticket, time.Hour)

		// todo: 此处可将 suite_ticket 保存进数据库

		fmt.Println("收到suite_ticked: ", ticket)

	}
}
