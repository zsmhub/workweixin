package callbacks

var callBackParseExtraInfoMap = make(map[string]CallBackExtraInfoInterface)

type CallBackExtraInfoInterface interface {
	GetMessageType() string
	GetEventType() string
	GetChangeType() string
	GetTypeKey() string
	ParseFromXml(data []byte) (CallBackExtraInfoInterface, error)
}

func supportCallback(item CallBackExtraInfoInterface) {
	callBackParseExtraInfoMap[item.GetTypeKey()] = item
}
