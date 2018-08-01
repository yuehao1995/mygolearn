### 用户操作
#####增加用户
useradd  
-m 自动建立目录  
切换用户  
su zyh 切换到zyh  
sudo useradd ly -m 建立用户，在home下  
useradd -d /home/ly 无则创建  
sudo passwd ly 修改用户名密码  
passwd 修改自己密码  
userdel ly 删除用户  
userdel ly -r  
sudo -s 切到超级管理员  
#####用户组
groupadd ly添加一个组  
groupdel ly删除组  
usermod 修改用户组  
groupadd test 添加用户组test  
usermod -g test ly 将ly加入组  




