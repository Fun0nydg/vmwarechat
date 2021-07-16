# simplechat
## 介绍
由于vmtools实在很难用，虚拟机和物理机之间复制文本和传文件很麻烦，于是就写一个简单的小工具，可以互发消息、传文件。   
首先运行server，之后运行client，这样client之间可以互相发消息。

## 编译
需要go环境  
首先下载本项目，进入到server目录，编译server：
```bash
CGO_ENABLED=0 go build -v -a -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ./server server.go
```
同理编译client
## 运行
编译好了之后运行server,默认监听8000端口:
```bash
./server
```

client连接，ip参数指定地址，p参数指定端口：
```bash
./client -ip 10.0.0.1 -p 8000
```
先输入用户名（3位）之后回车，开始选择模式：
#### 1.发消息
格式：  
post xxxx  
按回车将会发送消息给其他客户端  

#### 2.传文件
格式：  
file 1.txt  
按回车将会传文件给server，文件必须在client目录下。  

### 参考
- https://github.com/OctopusLian/Golang-OnlineChatRoom/tree/master/OnetoMoreChatRoom_V2
- https://www.cnblogs.com/yang-2018/p/11147418.html
