## sdk调用示例

1. 请先配置下 demo/config.go 的相关参数
2. 运行测试用例

    ```sh
    # 进入demo文件夹
    cd ./demo

    # 测试接收回调事件的相关代码
    go test -v -run=TestCallbackMain

    # 测试企微API调用的相关代码
    go test -v -run=TestApiMain
    ```