###go-micro安装及使用
#####安装

* 1.安装protobuf:https://github.com/google/protobuf/releases
* 2.安装protoc-gen-go:go get  github.com/golang/protobuf/protoc-gen-go
* 3.安装go-micro生成工具:go get github.com/micro/protoc-gen-micro

#####编译
* 4.src下生成pb.go及micro.go : protoc -I ./ mygolearn/mygomicro/proto/*.proto --micro_out=. --go_out=plugins=grpc:.

#####服务端使用
实现Greeter服务中的Hello接口  
新建服务,服务初始化  
给服务注册一个处理函数  
启动服务
#####客户端的使用  
新建一个服务  
服务初始化
对服务注册一个客户端  
调用rpc
