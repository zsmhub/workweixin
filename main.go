package workweixin

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lovego/workweixin/apis"
	"github.com/lovego/workweixin/callbacks"
)

// sdk 调用入口
var Sdk = &sdk{
	mutex:                     sync.RWMutex{},
	ThirdAuthCorpsClient:      make(map[string]*apis.ApiClient),
	CustomizedAuthCorpsClient: make(map[string]*apis.ApiClient),
	SelfAuthCorpsClient:       make(map[string]*apis.ApiClient),
}

type sdk struct {
	mutex sync.RWMutex

	// 服务商的回调处理和Api客户端
	ProviderCallback *callbacks.CallBackHandler // 服务商回调解析 [配置位置：服务商后台 -> 应用管理 -> 通用开发参数 -> 系统事件接收URL]
	ProviderClient   *apis.ApiClient            // 服务商客户端，用于服务商级别的接口调用，比如登录授权、推广二维码等

	// 第三方应用的回调处理和Api客户端
	ThirdAppCallback     *callbacks.CallBackHandler // 第三方应用回调解析 [配置位置：服务商后台 -> 应用管理 -> 网页应用 -> 数据回调URL和指令回调URL]
	ThirdAppClient       *apis.ApiClient            // 第三方应用客户端，用于获取第三方应用的预授权码，获取授权企业信息等
	ThirdAuthCorpsClient map[string]*apis.ApiClient // 第三方应用授权企业客户端，用于操作授权企业相关接口，如通讯录管理，消息推送等

	// 自建应用代开发的回调处理和Api客户端
	CustomizedTemplateCallback *callbacks.CallBackHandler // 自建应用代开发模板回调解析
	CustomizedAppClient        *apis.ApiClient            // 自建应用代开发客户端，用于获取第三方应用的预授权码，获取授权企业信息等
	CustomizedAuthCorpsClient  map[string]*apis.ApiClient // 代发开应用授权企业客户端，用于操作授权企业相关接口，如通讯录管理，消息推送等

	// 自建应用API客户端
	SelfAuthCorpsClient map[string]*apis.ApiClient // 自建应用客户端，用于操作企业相关接口，如通讯录管理，消息推送等

	// 第三方小程序的回调处理
	ThirdMiniCallback *callbacks.CallBackHandler // 第三方小程序回调解析，如果第三方应用需要使用发微信小程序功能，需要先配置一个第三方小程序绑定微信小程序

	apis.Options
}

// 初始化企微sdk参数
func (s *sdk) InitOptions(opts apis.Options) {
	s.Options = opts
}

// 服务商：设置服务商回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewProviderCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.ProviderCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}

// 服务商：API客户端初始化
func (s *sdk) NewProviderApiClient(corpId, corpProviderSecret string) {
	s.ProviderClient = apis.NewProviderApiClient(corpId, corpProviderSecret, s.Options)
}

// 第三方应用：设置第三方应用回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewThirdAppCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.ThirdAppCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}

// 第三方应用：API客户端初始化
func (s *sdk) NewThirdAppApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket string) {
	s.ThirdAppClient = apis.NewThirdAppApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket, s.Options)
}

// 第三方应用：授权企业API客户端初始化
func (s *sdk) NewThirdAuthCorpApiClient(corpId, companyPermanentCode string, agentId int) error {
	if s.ThirdAppClient == nil {
		return errors.New("Sdk.ThirdAppClient还未初始化")
	}
	apiClient := apis.NewAuthCorpApiClient(corpId, companyPermanentCode, agentId, s.ThirdAppClient, s.Options)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ThirdAuthCorpsClient[corpId] = apiClient
	return nil
}

// 第三方应用：获取授权企业客户端
func (s *sdk) GetThirdAuthCorpApiClient(corpId string) (*apis.ApiClient, error) {
	// 从数据库或缓存读取企业数据
	if s.GetThirdAppAuthCorpFunc != nil {
		corp, err := s.GetThirdAppAuthCorpFunc(corpId, s.ThirdAppClient.AppSuiteId)
		if err != nil {
			return nil, err
		} else if corp.PermanentCode != "" {
			if err := s.NewThirdAuthCorpApiClient(corpId, corp.PermanentCode, corp.AgentId); err != nil {
				return nil, err
			}
		}
	}

	s.mutex.RLock()
	if v, ok := s.ThirdAuthCorpsClient[corpId]; ok {
		s.mutex.RUnlock()
		return v, nil
	}
	s.mutex.RUnlock()

	return nil, fmt.Errorf("第三方应用：corpid不存在：%s", corpId)
}

// 第三方应用：移除授权企业，比如企业取消授权时触发
func (s *sdk) RemoveThirdAuthCorp(corpId string) {
	apiClient, err := s.GetThirdAuthCorpApiClient(corpId)
	if err != nil {
		return
	}

	apiClient.RemoveToken()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.ThirdAuthCorpsClient, corpId)
}

// 自建应用代开发：设置应用回调token和加密key(如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewCustomizedTemplateCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.CustomizedTemplateCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}

// 自建应用代开发：API客户端初始化
func (s *sdk) NewCustomizedApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket string) {
	s.CustomizedAppClient = apis.NewCustomizedApiClient(corpId, appSuiteID, appSuiteSecret, appSuiteTicket, s.Options)
}

// 自建应用代开发：授权企业API客户端初始化
func (s *sdk) NewCustomizedAuthCorpApiClient(corpId, companyPermanentCode string, agentId int) error {
	if s.CustomizedAppClient == nil {
		return errors.New("Sdk.CustomizedAppClient还未初始化")
	}

	apiClient := apis.NewCustomizedAuthCorpApiClient(corpId, companyPermanentCode, agentId, s.CustomizedAppClient, s.Options)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.CustomizedAuthCorpsClient[corpId] = apiClient
	return nil
}

// 自建应用代开发：获取授权企业客户端
func (s *sdk) GetCustomizedAuthCorpApiClient(corpId string) (*apis.ApiClient, error) {
	// 从数据库或缓存读取企业数据
	if s.GetCustomizedAppAuthCorpFunc != nil {
		corp, err := s.GetCustomizedAppAuthCorpFunc(corpId, s.CustomizedAppClient.AppSuiteId)
		if err != nil {
			return nil, err
		} else if corp.PermanentCode != "" {
			if err := s.NewCustomizedAuthCorpApiClient(corpId, corp.PermanentCode, corp.AgentId); err != nil {
				return nil, err
			}
		}
	}

	s.mutex.RLock()
	if v, ok := s.CustomizedAuthCorpsClient[corpId]; ok {
		s.mutex.RUnlock()
		return v, nil
	}
	s.mutex.RUnlock()

	return nil, fmt.Errorf("自建应用代开发：corpid不存在：%s", corpId)
}

// 自建应用代开发：移除授权企业，比如企业取消授权时触发
func (s *sdk) RemoveCustomizedAuthCorp(corpId string) {
	apiClient, err := s.GetCustomizedAuthCorpApiClient(corpId)
	if err != nil {
		return
	}

	apiClient.RemoveToken()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.CustomizedAuthCorpsClient, corpId)
}

// 自建应用：API客户端初始化
func (s *sdk) NewSelfApiClient(corpId string, agentId int, secret string) {
	apiClient := apis.NewSelfApiClient(corpId, agentId, secret, s.Options)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SelfAuthCorpsClient[corpId] = apiClient
}

// 自建应用：获取企业客户端
func (s *sdk) GetSelfAuthCorpApiClient(corpId string) (*apis.ApiClient, error) {
	s.mutex.RLock()
	if v, ok := s.SelfAuthCorpsClient[corpId]; ok {
		s.mutex.RUnlock()
		return v, nil
	}
	s.mutex.RUnlock()

	return nil, fmt.Errorf("自建应用：corpid不存在：%s", corpId)
}

// 自建应用：移除授权企业
func (s *sdk) RemoveSelfAuthCorp(corpId string) {
	apiClient, err := s.GetSelfAuthCorpApiClient(corpId)
	if err != nil {
		return
	}

	apiClient.RemoveToken()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.SelfAuthCorpsClient, corpId)
}

// 第三方小程序：设置第三方小程序回调token和加密key (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewThirdMiniCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.ThirdMiniCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}
