# 容器与宿主机之间交流资源
### 从宿主机拷贝文件到docker容器内
###### 文件拷贝
1.docker cp testfile mycentos:/usr/local 从宿主机拷贝到容器  
2.docker cp mycentos:/usr/local/testfile testf 从容器拷贝到宿主机
3.不容许容器之间直接拷贝
###### 目录挂载
* 宿主机目录映射给容器内某目录
1. docker run -id --name=mycentos3 -v /usr/local/myhtml:/usr/local/mh centos:latest
把宿主机文件映射到容器内  
此时容器内读文件里内容会被拒绝，无权限
2. docker run -id --name=mycentos2 -v /usr/local/myhtml:/usr/local/md --privileged=true  centos:latest   
容器内有权限读取文件及其下内容
### 容器之间通信 
###### 容器ip
1. docker inspect mycentos 查询容器Ip,在容器信息最后的ipaddress,此ip与其他机器无法接通
2. docker inspect mycentos --format='{{.NetworkSettings.IPAddress}}' mycentos
筛选仅显示ip
### 删除容器
1. docker rm mycentos3 删除容器
2. 运行，开机的容器是无法删除的

