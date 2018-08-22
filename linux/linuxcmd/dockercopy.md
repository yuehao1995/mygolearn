# 容器与镜像
### 容器保存为镜像
1. docker commit zyh_mysql zyhsql  将zyh_mysql保存为镜像
### 镜像备份与恢复
1.docker save -o zyhsql.tar zyhsql 将镜像导出为一个文件  
2.docker rmi zyh_mysql 删除镜像  
3.docker load -i zyhsql.tar  将文件恢复为镜像
