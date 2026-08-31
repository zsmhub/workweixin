package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lovego/workweixin/apis"
)

var ctx = context.Background()

var redisDb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// 实现企微 access_token 的 redis 存取方案
type DcsTokenByRedis struct{}

var _ apis.DcsToken = DcsTokenByRedis{}

func (DcsTokenByRedis) Get(cacheKey string) apis.TokenInfo {
	var tokenInfo apis.TokenInfo
	result, err := redisDb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		_ = json.Unmarshal(result, &tokenInfo)
	} else if err != redis.Nil {
		log.Println("获取 access_token 失败：", err)
	}

	return tokenInfo
}

func (DcsTokenByRedis) Set(cacheKey string, tokenInfo apis.TokenInfo, ttl time.Duration) error {
	data, _ := json.Marshal(tokenInfo)
	err := redisDb.Set(ctx, cacheKey, string(data), ttl).Err()
	if err != nil {
		log.Println("保存 access_token 失败:", err)
	}
	return err
}

func (DcsTokenByRedis) Del(cacheKey string) error {
	return redisDb.Del(ctx, cacheKey).Err()
}

func (DcsTokenByRedis) Lock(cacheKey string, ttl time.Duration) bool {
	if ok, _ := redisDb.SetNX(ctx, cacheKey, 1, ttl).Result(); ok {
		return true
	}
	return false
}

func (DcsTokenByRedis) Unlock(cacheKey string) error {
	return redisDb.Del(ctx, cacheKey).Err()
}

// 实现企微 suite_ticket 的 redis 存取方案
type DcsAppSuiteTicketByRedis struct{}

var _ apis.DcsAppSuiteTicket = DcsAppSuiteTicketByRedis{}

func (DcsAppSuiteTicketByRedis) Get(cacheKey string) string {
	suiteTicket, err := redisDb.Get(ctx, cacheKey).Result()
	if err != nil {
		return ""
	}
	return suiteTicket
}

func (DcsAppSuiteTicketByRedis) Set(cacheKey, suiteTicket string, ttl time.Duration) {
	err := redisDb.Set(ctx, cacheKey, suiteTicket, ttl).Err()
	if err != nil {
		log.Println("保存 suite_ticket 失败:", err)
	}
}

// 日志记录器，可按需改造
type Logger struct{}

var _ apis.Logger = Logger{}

func (Logger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (Logger) Infof(template string, args ...interface{}) {
	fmt.Printf(template, args...)
}

func (Logger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (Logger) Errorf(template string, args ...interface{}) {
	fmt.Printf(template, args...)
}

// 获取第三方应用的企业数据
func GetThirdAppAuthCorpToSdk(corpId, appSuiteId string) (apis.AuthCorp, error) {
	// todo
	return apis.AuthCorp{PermanentCode: "xxx", AgentId: 1}, nil
}

// 获取自建应用代开发的企业数据
func GetCustomizedAppAuthCorpToSdk(corpId, appSuiteId string) (apis.AuthCorp, error) {
	// todo
	return apis.AuthCorp{PermanentCode: "xxx", AgentId: 2}, nil
}
