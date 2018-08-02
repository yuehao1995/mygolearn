# docker镜像
### docker镜像的含义
* 镜像是一系列文件的集合  
1. 文件集合里包含了安装装软件及相应信息
2. Docker引导镜像之后就形成了容器
3. 容器的基础是镜像
### docker镜像操作
* 列出镜像  
 docker images  
|镜像对应列|镜像|  
|--|--|  
|REPOSITORY|镜像软件名|  
|TAG|版本|  
|IMAGE ID|镜像标识|  
|CREATED|创建时间|  
|SIZE|大小|  

* 搜索镜像（查网络里的）
docker search mysql

### 拉取镜像
* 拉取镜像的含义
1. 拉取就是从注册中心下载镜像
2. 注册中心别名Docker Hub

* 拉取镜像命令
docker pull ubuntu等 

* 镜像加速器 
1. vi /etc/docker/daemon.json  
加入:
```json
{
"registry-mirrors":["https://docker.mirrors.ustc.edu.cn"]
}
```
2.重启docker服务  
systemctl restart docker  
3.拉取镜像

docker pull zookeeper



