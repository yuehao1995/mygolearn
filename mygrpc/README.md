### 下载安装protobuf软件
### 下载安装 google.golang.org/grpc
### 下载安装go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
##### 服务端
1. 定义好proto文件  
2. 生成pb.go
3. models中定义结构体,实现服务对应的rpc
4. 注册服务
5. 调用对应的服务
##### 客户端
1. 定义好proto文件  
2. 生成pb.go    
3. 新建客户端  
4. 调用相应的rpc

