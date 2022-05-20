# 生成企微api代码
api:
	go run api_generate/main.go -doc=$(doc)

# 生成企微回调事件代码
callback:
	go run callback_generate/main.go -doc=$(doc)