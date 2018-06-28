#linux命令高级
###grep命令
grep '内容' test.go 搜索文件中的内容显示出含内容的行  
grep -n '搜索内容' test.go 搜索并显示行号  
grep -n 'com$' test.go com结尾的内容  
grep -n '^http' test.go http开头的内容  
grep -n '[Aa]' test.go A或a开头  
grep -n  'A.a' test.go .可以代替一个字符  
### find命令
find . -name '*.go' 从当前路径及当前下寻找.go文件  
find . -size +2M 大小大于2M  
find . -perm 0777 寻找权限为0777的文件