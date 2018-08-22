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
docker pull ubuntu:tag等 

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
docker pull  zookeeper
4. 删除镜像
删除镜像时，假如有容器在运行，不会删除  
docker rmi 镜像id   删除指定镜像  
docker rmi `docker images  -q` 删除所有镜像  
``里边代表命令的结果
docker images -q 查找所有镜像id
### 容器操作
* 容器就是运行状态的镜像
1. 查看容器
docker ps 查看正在运行的容器  
docker ps -a 查看所有的容器(历史运行过的容器)  
docker ps -l 查看最后一次运行的容器  
docker ps -f status=exited 查看停止状态的容器
2. 创建和使用容器  
docker run 创建容器命令    
docker run -i 表示运行容器
docker run  -t 表示启动后进入命令行,-i -t启动后进入容器  
docker run --name= 表示为创建的容器命名  
docker run -d 表示创建一个守护式容器在后台运行,不进入  
docker run -v表示目录映射关系，宿主机共享文件给容器  
docker run -p表示端口映射，前者是宿主机端口，后者是容器内的映射端口,从而使宿主机外的计算机能访问容器  
3 交互式容器与守护式容器  
* 创建后能立即进入容器为交互式,创建后不立即进入而在后台运行为守护式容器  
实际操作  
* 交互式容器
1. 启动centos，从镜像名centos标签为latest的镜像启动
docker run -it --name=mycentos centos:latest /bin/bash  
2. exit 退出容器  
docker start mycentos启动关机后的容器  
docker stop mycentos停止容器
3. 一个镜像可以运行多个容器
4.交互式容器退出时会自动关机
* 守护式容器  
1. docker run -id  --name=mycentos2 centos:latest  创建守护式容器
2. docker exec -it mycentos2 /bin/bash 启动后台容器  
3. exit退出守护式容器时依旧开机