package workweixin

import (
	"errors"
	"fmt"
	"github.com/zsmhub/workweixin/apis"
	"github.com/zsmhub/workweixin/callbacks"
	"sync"
)

// sdk 调用入口
type sdk struct {
	mutex sync.RWMutex

	// callback
	ProviderCallback  *callbacks.CallBackHandler // 服务商回调解析 [配置位置：服务商后台 -> 应用管理 -> 通用开发参数 -> 系统事件接收URL]
	AppSuiteCallback  *callbacks.CallBackHandler // 第三方应用回调解析 [配置位置：服务商后台 -> 应用管理 -> 网页应用 -> 数据回调URL和指令回调URL]
	MiniSuiteCallback *callbacks.CallBackHandler // 第三方小程序回调解析，如果第三方应用需要使用发微信小程序功能，需要先配置一个第三方小程序绑定微信小程序

	// api
	ProviderClient  *apis.ApiClient            // 服务商客户端，用于服务商级别的接口调用，比如登录授权、推广二维码等
	ThirdAppClient  *apis.ApiClient            // 第三方应用客户端，用于获取第三方应用的预授权码，获取授权企业信息等
	AuthCorpsClient map[string]*apis.ApiClient // 授权企业客户端，用于操作授权企业相关接口，如通讯录管理，消息推送等
}

var Sdk = &sdk{mutex: sync.RWMutex{}, AuthCorpsClient: make(map[string]*apis.ApiClient)}

// 设置服务商回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewProviderCallbackHandler(token, encodingAESKey string) (err error) {
	s.ProviderCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)
	return
}

// 设置第三方应用回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewAppSuiteCallbackHandler(token, encodingAESKey string) (err error) {
	s.AppSuiteCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)
	return
}

// 设置第三方小程序回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewMiniSuiteCallbackHandler(token, encodingAESKey string) (err error) {
	s.MiniSuiteCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)
	return
}

// 服务商API客户端初始化
func (s *sdk) NewProviderApiClient(corpId, corpProviderSecret string) {
	s.ProviderClient = apis.NewProviderApiClient(corpId, corpProviderSecret)
	return
}

// 第三方应用API客户端初始化
func (s *sdk) NewThirdAppApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket string) {
	s.ThirdAppClient = apis.NewThirdAppApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket)
	return
}

// 授权企业API客户端初始化，非单体服务则需使用广播通知多个pod做处理
func (s *sdk) NewAuthCorpApiClient(corpId, companyPermanentCode string, agentId int) error {
	if s.ThirdAppClient == nil {
		return errors.New("sdk.ThirdAppClient还未初始化")
	}
	apiClient := apis.NewAuthCorpApiClient(corpId, companyPermanentCode, agentId, s.ThirdAppClient)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.AuthCorpsClient[corpId] = apiClient
	return nil
}

// 获取授权企业客户端
func (s *sdk) GetAuthCorpAPPClient(corpId string) (*apis.ApiClient, error) {
	s.mutex.RLock()
	if v, ok := s.AuthCorpsClient[corpId]; ok {
		s.mutex.RUnlock()
		return v, nil
	}
	s.mutex.RUnlock()

	// todo 此处可以自行加一段代码，从数据库取数
	/*dbData := db.GetCorpData(corpId)
	if dbData.Id > 0 {
		if !dbData.AuthStatus {
			return nil, fmt.Errorf("该企业已取消授权：%s", corpId)
		}
		if err := s.NewAuthCorpApiClient(dbData.CorpId, dbData.PermanentCode, dbData.AgentId); err != nil {
			return nil, err
		}
		s.mutex.RLock()
		if v, ok := s.AuthCorpsClient[corpId]; ok {
			s.mutex.RUnlock()
			return v, nil
		}
		s.mutex.RUnlock()
	} else {
		// 企业不存在或数据库请求异常
		return nil, fmt.Errorf("该企业不存在：%s", corpId)
	}*/

	return nil, errors.New(fmt.Sprintf("corpid不存在：%s", corpId))
}

// 移除授权企业，非单体服务则需使用广播通知多个pod做处理
func (s *sdk) RemoveAuthCorp(corpId string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.AuthCorpsClient, corpId)
}
