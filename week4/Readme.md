1、 protobuf 使用

2、Wire 构建依赖

安装 ：go get -u github.com/google/wire/cmd/wire

provider 生成组建的普通方法。接受所需依赖为参数，创建组建并将其返回

injector 对象创建及初始化函数

对象定义provider在main.go, injector依赖关系配置在wire.go, injector依赖关系初始化在wire_gen.go

执行 go run . 进行验证测试

provider 生成组建的普通方法。接受所需依赖为参数，创建组建并将其返回

参考 : https://zhuanlan.zhihu.com/p/110453784

3、go 信号处理

4、写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理
