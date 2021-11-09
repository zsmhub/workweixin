#生成api代码
api:
	go run api_generate/main.go -doc=$(doc)

callback:
	go run callback_generate/main.go -doc=$(doc)

