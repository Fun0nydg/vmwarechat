# simplechat
## 介绍
一个简单的客户端之间发消息的小工具。   
首先运行server，之后运行client，这样client之间可以互相发消息。

## 编译
需要go环境  
首先下载本项目，进入到server目录，编译server：
```bash
CGO_ENABLED=0 go build -v -a -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ./server server.go
```
同理编译client，编译好了之后运行server,默认监听8000端口:
```bash
./server
```

client连接，ip参数指定地址，p参数指定端口：
```bash
./client -ip 10.0.0.1 -p 8000
```
输入用户名之后回车，开始发送消息，按回车发送，'Q'结束