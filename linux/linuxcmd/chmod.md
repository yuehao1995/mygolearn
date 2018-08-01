###修改文件权限
chmod u=r,g=r,o=w xxx 修改文件的权限为当前用户为读，同组为读，其他人为写  
chmod u=rwx xxx 修改当前用户对xxx文件的权限为rwx,同组及其他人不改  
chmod u+w xxx对当前用户权限加上可写  
chmod u-w xxx 修改当前用户对xxx的权限减去可写  
chmod 664 auth.go 修改权限r4w2x1  
