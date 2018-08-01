# Docker学习
##### ubuntu18安装docker
1. 卸载就版本docker  
sudo apt-get remove docker docker-engine docker.io  
2. 安装软件包来允许apt通过HTTPS使用存储库  
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common  
3.添加Docker的官方GPG密钥  
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -  
4.添加docker的下载源，因为官方还没有ubuntu18的下载源，所以先用ubuntu17（zesty）的  
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu zesty stable"  
5.安装docker-ce  
sudo apt-get update  
sudo apt-get install docker-ce  
##### docker启动等
1. docker -v 查看docker版本   
2. systemctl docker start 启动服务  
3. systemctl  status docker 查看docker状态  
4. systemctl stop docker关闭docker  
5. systemctl enable docker开机启动  
##### docker 命令等
1.sudo docker info查看docker信息  
2.sudo docker --help docker帮助  
docker service --help docker具体命令帮助  
##### docker 镜像介绍
容器的基础是镜像，镜像是文件集合，docker引导镜像之后就是容器  
##### docker 镜像命令  
1.docker images 列出镜像  
|镜像对应列|镜像|
|--|--|
|REPOSITORY|镜像软件名|
|TAG|版本|
|IMAGE ID|镜像标识|
|CREATED|创建时间|
|SIZE|大小|






