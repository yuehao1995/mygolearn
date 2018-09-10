# go包管理
### go路径设置
可以设置多个，运行项目时会自动从gopath下去找包,其中导入包的规则是从src下开始
###### windows
设置gopath
###### linux
echo $path  
在~/.bash_profile中加入以下几句话:  
export GOPATH=  
export PATH="$GOPATH/bin:$PATH"  
多个名字之间：分割
### 系统库
会自动找到
###### 被墙的处理
go get -v github.com/gpmgo/gopm  
gopm get -g  -v github.com/spf13/viper  
其中 -g代表下载到gopath，-v代表显示下载过程
######  安装
go install github.com/spf13/viper

go 每个目录下都有一个main文件

go install 可以安装当前文件下所有含有main且只有一个main入口的包