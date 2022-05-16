package apis

import (
	"context"
	"errors"
	"github.com/cenkalti/backoff/v4"
	"log"
	"sync"
	"time"
)

// 分布式access_token：获取和设置access_token的值，自行实现该接口的具体逻辑，比如使用redis方案
type DcsToken interface {
	Get(cacheKey string) TokenInfo                                     // 获取access_token
	Set(cacheKey string, tokenInfo TokenInfo, ttl time.Duration) error // 设置access_token，ttl：缓存生存时间
	Lock(cacheKey string, ttl time.Duration) bool                      // 加锁，返回成功或失败
	Unlock(cacheKey string) error                                      // 释放锁
}

// access_token
type TokenInfo struct {
	Token       string        `json:"token"`        // access_token/jsapi_ticket
	ExpiresIn   time.Duration `json:"expires_in"`   // 过期时间
	LastRefresh time.Time     `json:"last_refresh"` // 上次刷新access_token时间
}

type token struct {
	mutex *sync.RWMutex
	TokenInfo
	getTokenFunc  func() (TokenInfo, error)
	dcsToken      DcsToken
	tokenCacheKey string
}

/*** access_token 模块 -begin- ***/
func (c *ApiClient) GetToken() (token string, err error) {
	token = c.accessToken.getToken()
	if token == "" {
		err = errors.New("access_token获取失败，可能原因：配置有误/没有在服务商后台设置IP白名单/企业取消授权")
		return
	}
	return
}

// GetJSAPITicket 获取 JSAPI_ticket
func (c *ApiClient) GetJSAPITicket() (string, error) {
	return c.jsapiTicket.getToken(), nil
}

// GetJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *ApiClient) GetJSAPITicketAgentConfig() (string, error) {
	return c.jsapiTicketAgentConfig.getToken(), nil
}

// 更新suite_ticket
func (c *ApiClient) RefreshSuiteTicket(ticket string, ttl time.Duration) {
	if c.dcsAppSuiteTicket != nil {
		c.dcsAppSuiteTicket.Set(c.dcsSuiteTicketCacheKey, ticket, ttl)
	} else {
		c.AppSuiteTicket = ticket
	}
}

// 获取服务商的provider_access_token
func (c *ApiClient) getProviderAccessToken() (TokenInfo, error) {
	req := ReqGetProviderTokenService{
		Corpid:         c.CorpId,
		ProviderSecret: c.CorpProviderSecret,
	}
	get, err := c.ExecGetProviderTokenService(req)
	if err != nil {
		c.logger.Errorf("provider_access_token: ReqGetProviderTokenService=%+v, err=%+v\n", req, err)
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.ProviderAccessToken, ExpiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// 第三方应用的suite_access_token
func (c *ApiClient) getSuiteToken() (TokenInfo, error) {
	if c.dcsAppSuiteTicket != nil {
		if tmp := c.dcsAppSuiteTicket.Get(c.dcsSuiteTicketCacheKey); tmp != "" {
			c.AppSuiteTicket = tmp
		}
	}
	if c.AppSuiteTicket == "" {
		return TokenInfo{}, errors.New("服务商的" + c.accessTokenName + "缺失，app_suite_id:" + c.AppSuiteId)
	}
	req := ReqGetSuiteTokenService{
		SuiteID:     c.AppSuiteId,
		SuiteSecret: c.AppSuiteSecret,
		SuiteTicket: c.AppSuiteTicket,
	}
	get, err := c.ExecGetSuiteTokenService(req)
	if err != nil {
		c.logger.Errorf(c.accessTokenName+": ReqGetSuiteTokenService=%+v, err=%+v\n", req, err)
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.SuiteAccessToken, ExpiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// 获取第三方应用授权企业access_token
func (c *ApiClient) getAuthCorpToken() (TokenInfo, error) {
	if c.CompanyPermanentCode == "" {
		return TokenInfo{}, errors.New("企业永久授权码不存在，corp_id:" + c.CorpId)
	}
	get, err := c.ExecGetCorpTokenService(ReqGetCorpTokenService{
		AuthCorpid:    c.CorpId,
		PermanentCode: c.CompanyPermanentCode,
	})
	if err != nil {
		apiError, ok := err.(*ClientError)
		if ok {
			if apiError.Code == ErrCode2000002 || apiError.Code == ErrCode301007 || apiError.Code == ErrCode40084 { // 企业已注销，但要等15天后才会收到企业取消授权事件
				return TokenInfo{}, nil
			}
			c.logger.Errorf("corp_access_token1: corp_id=%s, permanent_code=%s, err=%+v\n", c.CorpId, c.CompanyPermanentCode, apiError)
		} else {
			c.logger.Errorf("corp_access_token2: corp_id=%s, permanent_code=%s, err=%+v\n", c.CorpId, c.CompanyPermanentCode, err)
		}
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.AccessToken, ExpiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// 获取自建应用代开发授权企业access_token
func (c *ApiClient) getCustomizedAuthCorpToken() (TokenInfo, error) {
	if c.CompanyPermanentCode == "" {
		return TokenInfo{}, errors.New("代开发企业永久授权码不存在，corp_id:" + c.CorpId)
	}
	get, err := c.ExecGetCustomizedCorpTokenService(ReqGetCustomizedCorpTokenService{
		Corpid:     c.CorpId,
		Corpsecret: c.CompanyPermanentCode,
	})
	if err != nil {
		apiError, ok := err.(*ClientError)
		if ok {
			if apiError.Code == ErrCode2000002 || apiError.Code == ErrCode301007 || apiError.Code == ErrCode40084 { // 企业已注销，但要等15天后才会收到企业取消授权事件
				return TokenInfo{}, nil
			}
			c.logger.Errorf("customized_corp_access_token1: corp_id=%s, permanent_code=%s, err=%+v\n", c.CorpId, c.CompanyPermanentCode, apiError)
		} else {
			c.logger.Errorf("customized_corp_access_token2: corp_id=%s, permanent_code=%s, err=%+v\n", c.CorpId, c.CompanyPermanentCode, err)
		}
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.AccessToken, ExpiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *ApiClient) getJSAPITicket() (TokenInfo, error) {
	get, err := c.ExecGetJSAPITicket(JsAPITicketReq{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.Ticket, ExpiresIn: time.Duration(get.ExpiresInSecs) * time.Second}, nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *ApiClient) getJSAPITicketAgentConfig() (TokenInfo, error) {
	get, err := c.ExecGetJSAPITicketAgentConfig(JsAPITicketAgentConfigReq{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.Ticket, ExpiresIn: time.Duration(get.ExpiresInSecs) * time.Second}, nil
}

func (t *token) setGetTokenFunc(f func() (TokenInfo, error)) {
	t.getTokenFunc = f
}

func (t *token) getToken() string {
	if err := Retry(t.syncToken); err != nil {
		log.Printf("retry getting access Token failed, err=%+v\n", err)
	}

	t.mutex.RLock()
	tokenToUse := t.Token
	t.mutex.RUnlock()

	return tokenToUse
}

func (t *token) syncToken() error {
	var tokenInfo TokenInfo

	if t.dcsToken != nil {
		if t.Token == "" {
			tokenInfo = t.dcsToken.Get(t.tokenCacheKey)
		}
		if tokenInfo.Token == "" || tokenInfo.LastRefresh.Unix()+3600 < time.Now().Unix() {
			lockCacheKey := t.tokenCacheKey + "#lock"

			// 抢锁
			if ok := t.dcsToken.Lock(lockCacheKey, 10*time.Second); ok {
				defer func() {
					_ = t.dcsToken.Unlock(lockCacheKey)
				}()

				get, err := t.getTokenFunc()
				if err != nil {
					return err
				}

				tokenInfo.Token = get.Token
				tokenInfo.ExpiresIn = get.ExpiresIn
				tokenInfo.LastRefresh = time.Now()

				if err := t.dcsToken.Set(t.tokenCacheKey, tokenInfo, time.Hour*2); err != nil {
					return err
				}
			} else {
				// 抢锁失败则等待
				time.Sleep(time.Second * 2)
				tokenInfo = t.dcsToken.Get(t.tokenCacheKey)
			}
		}
	} else {
		// 兼容方案
		if t.Token == "" || t.LastRefresh.Unix()+3600 < time.Now().Unix() {
			get, err := t.getTokenFunc()
			if err != nil {
				return err
			}
			tokenInfo.Token = get.Token
			tokenInfo.ExpiresIn = get.ExpiresIn
			tokenInfo.LastRefresh = time.Now()
		}
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Token = tokenInfo.Token
	t.ExpiresIn = tokenInfo.ExpiresIn
	t.LastRefresh = tokenInfo.LastRefresh

	return nil
}

func Retry(o backoff.Operation) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()
	retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	return backoff.Retry(o, retryer)
}
