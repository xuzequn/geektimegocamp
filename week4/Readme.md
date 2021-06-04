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
