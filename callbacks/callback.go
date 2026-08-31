package callbacks

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/lovego/workweixin/internal/encryptor"
	"github.com/lovego/workweixin/internal/envelope"
	"github.com/lovego/workweixin/internal/signature"
)

type callBackUrlVars struct {
	MsgSignature string
	Timestamp    int64
	Nonce        string
	EchoStr      string
}

type CallBackHandler struct {
	token     string // 回调 token
	encryptor *encryptor.WorkWXEncryptor
	ep        *envelope.Processor
}

func NewCallbackHandler(token string, encodingAESKey string) (*CallBackHandler, error) {
	enc, err := encryptor.NewWorkWXEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	ep, err := envelope.NewProcessor(token, encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &CallBackHandler{token: token, encryptor: enc, ep: ep}, nil
}

// 解析并获取回调数据
func (cb *CallBackHandler) GetCallBackMsg(r *http.Request) (CallbackMessage, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CallbackMessage{}, err
	}

	// 验签
	ev, err := cb.ep.HandleIncomingMsg(r.URL, body)
	if err != nil {
		return CallbackMessage{}, err
	}

	// 解析Xml
	message, err := CallbackMessage{}.ParseMessageFromXml(ev.Msg)
	if err != nil {
		return message, err
	}

	if message.AgentID == 0 && ev.AgentID != "" {
		message.AgentID, _ = strconv.ParseInt(ev.AgentID, 10, 64)
	}

	return message, nil
}

// 后台回调配置URL，申请校验
func (cb *CallBackHandler) EchoTestHandler(rw http.ResponseWriter, r *http.Request) {
	if !signature.VerifyHTTPRequestSignature(cb.token, r.URL, "") {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	args, err := cb.parseUrlVars(r.URL.Query())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := cb.encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(payload.Msg)
}

func (cb *CallBackHandler) parseUrlVars(urlVars url.Values) (callBackUrlVars, error) {
	var errMalformedArgs = errors.New("malformed arguments for echo test API")

	var msgSignature string
	{
		l := urlVars["msg_signature"]
		if len(l) != 1 {
			return callBackUrlVars{}, errMalformedArgs
		}
		msgSignature = l[0]
	}

	var timestamp int64
	{
		l := urlVars["timestamp"]
		if len(l) != 1 {
			return callBackUrlVars{}, errMalformedArgs
		}
		timestampStr := l[0]

		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return callBackUrlVars{}, errMalformedArgs
		}

		timestamp = timestampInt
	}

	var nonce string
	{
		l := urlVars["nonce"]
		if len(l) != 1 {
			return callBackUrlVars{}, errMalformedArgs
		}
		nonce = l[0]
	}

	var echoStr string
	{
		l := urlVars["echostr"]
		if len(l) != 1 {
			return callBackUrlVars{}, errMalformedArgs
		}
		echoStr = l[0]
	}

	return callBackUrlVars{
		MsgSignature: msgSignature,
		Timestamp:    timestamp,
		Nonce:        nonce,
		EchoStr:      echoStr,
	}, nil
}
