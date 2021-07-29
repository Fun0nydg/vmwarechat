# vmwarechat
## 介绍
由于vmtools实在很难用，虚拟机和物理机之间复制文本和传文件很麻烦，于是就写一个简单的小工具，可以互发消息、传文件。   
首先在物理机运行server和client，然后在虚拟机中运行一个client，这样虚拟机的client和物理机的client之间可以互相发消息。欢迎提交issues，我会不断完善！！！  
### v1.1  
支持client下载server文件，提高传输文件速度  

## 编译
需要go环境  
首先下载本项目，进入到server目录，编译server：
```bash
CGO_ENABLED=0 go build -v -a -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ./server server.go
```
同理编译client，不想编译的同学可以在releases中下载。
## 使用
物理机运行一个server和一个client，虚拟机运行一个client  
运行server,默认监听8000端口:
```bash
server.exe
```
client和server必须端口互通，**注意下物理机和虚拟机的防火墙**  
client连接参数：ip参数指定地址，p参数指定端口：
```bash
client.exe -ip 10.0.0.1 -p 8000
```
先输入**用户名（3位）** 之后回车，开始选择模式：
#### 1.发消息

格式：  
post xxxx  
按回车将会发送xxxx给server,如果有两个client,那么server会转发消息给client，**每个client必须先post一次消息才能接收到别的client的消息**  

#### 2.传文件
格式：  
file 1.txt  
**会开启8001端口，请让防火墙允许通过**，按回车将会传文件给server，**文件必须在client所在的当前目录，直接输入文件名，暂不支持路径文件名格式**  

#### 3.下载文件
格式：  
down 1.txt  
按回车会下载server端的文件  
## 问题
### 1.client连接成功，无法互发消息？
请确认server和client的cmd是否卡住，如果卡住，请多次按回车。    

## 参考
- https://github.com/OctopusLian/Golang-OnlineChatRoom/tree/master/OnetoMoreChatRoom_V2
- https://www.cnblogs.com/yang-2018/p/11147418.html
