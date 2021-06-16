1、 protobuf 使用

生成命令
protoc --go_out=. *.proto

2、Wire 构建依赖

安装 ：go get -u github.com/google/wire/cmd/wire

provider 生成组建的普通方法。接受所需依赖为参数，创建组建并将其返回

injector 对象创建及初始化函数

对象定义provider在main.go, injector依赖关系配置在wire.go, injector依赖关系初始化在wire_gen.go

执行 go run . 进行验证测试

provider 生成组建的普通方法。接受所需依赖为参数，创建组建并将其返回

参考 : https://zhuanlan.zhihu.com/p/110453784

3、go 信号处理


golang 对信号处理的包是 os/signal

方法 一是通notify方法监听信号； 二是stop方法取消监听

go信号通知机制可通过往channel中发送os.signal来实现

SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)

SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)

SIGKILL和SIGSTOP这两个信号既不能被应用程序捕获，也不能被操作系统阻塞或忽略。

Term表明默认动作为终止进程，Ign表明默认动作为忽略该信号，Core表明默认动作为终止进程同时输出core dump，Stop表明默认动作为停止进程。



如何实现优雅地重启一个Go网络程序呢。主要要解决两个问题：

1、重启进程不需要关闭端口。

2、保持已有的连接请求不中断，让其执行完成或超时。

大概个执行步骤：

fork一个继承侦听套接字的新进程。

子进程初始化并开始接受套接字上的连接。

子进程向父进程发信号，通知父进程停止接收连接并终止。


https://blog.csdn.net/guyan0319/article/details/90240731

4、写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理

grpc的使用

go get google.golang.org/grpc

https://studygolang.com/articles/23761?fr=sidebar

https://zhuanlan.zhihu.com/p/161473581

简单 RPC
客户端使用存根发送请求到服务器并等待响应返回，就像平常的函数调用一样。
rpc GetFeature(Point) returns (Feature) {}
服务器端流式 RPC
客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。从例子中可以看出，通过在 响应 类型前插入 stream 关键字，可以指定一个服务器端的流方法。
rpc ListFeatures(Rectangle) returns (stream Feature) {}
客户端流式 RPC
客户端写入一个消息序列并将其发送到服务器，同样也是使用流。一旦客户端完成写入消息，它等待服务器完成读取返回它的响应。通过在 请求 类型前指定 stream 关键字来指定一个客户端的流方法。
rpc RecordRoute(stream Point) returns (RouteSummary) {}
双向流式 RPC
是双方使用读写流去发送一个消息序列。两个流独立操作，因此客户端和服务器可以以任意喜欢的顺序读写：比如， 服务器可以在写入响应前等待接收所有的客户端消息，或者可以交替的读取和写入消息，或者其他读写的组合。 每个流中的消息顺序被预留。你可以通过在请求和响应前加 stream 关键字去制定方法的类型。
rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}

protobuf 生成 grpc 代码命令

protoc --go_out=plugins=grpc:. *.proto

