# 客户端协议代理功能

## 功能介绍

#### 开启本功能后，网关服务器以及游戏服务器等其他服务器，将预先对客户端上行和服务器下行的协议数据做前置转换，采用任意版本的协议文件(必要字段名必须与现有的协议保持一致)均可，避免了因协议序号混淆等频繁变动，而造成游戏服务器代码不必要的频繁改动

## 使用方法

> 1. 在此目录下建立proto目录
> 2. 将对应版本的proto协议文件和client_cmd.csv协议号文件复制到proto目录下
> 3. 到项目根目录下执行`make gen_client_proto`(本操作可能会修改proto文件，请注意备份)
> 4. 执行`protoc --go_out=. *.proto`，将proto目录下的proto协议文件编译成pb.go
> 5. 将gate服务器的配置文件中开启client_proto_proxy_enable客户端协议代理功能
