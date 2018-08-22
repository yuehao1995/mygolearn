  # 容器部署
  ### mysql
  ###### mysql容器搭建
  1. docker run -id --name=zyh_mysql -p 33306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:latest 创建mysql端口映射，并设置初始密码123456
  2.docker exec -it zyh_mysql  /bin/bash 进入容器
  3.mysql -u root -p 进入mysql
  4.quit 退出mysql
  5.exit 退出容器
  6.远程连接
  ### nginx
  ###### nginx容器搭建
  1.docker pull nginx 拉取镜像
  2.docker run -id  --name=zy_nginx -p 80:80 nginx:latest 创建容器
  3.docker tun -id
  ### redis
  ###### redis容器搭建
  1.docker run -di --name=zyh_redis -p 6379:6379 redis:latest 创建容器


