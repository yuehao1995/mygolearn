#打包，解压
* tar xxx.tar a.go b.go  打包a.go,b.go
* tar -cvf c打包v过程f 指定名称必须是.tar文件
 tar -cvf db.tar *.go 打包文件到db.tar
* tar xvf  db.tar 解压包
* gzip 压缩
  gzip db.tar 压缩到db.tar.gz
  gzip -d db.tar.gz 解压
* tar zcvf db.tar.gz *.go 打包并压缩
  tar zxvf db.tar.gz 解包并解压
  tar zxvf db.tar.gz -C ~/gzip 解压到指定路径
* tar jcvf db.tar.bz2 *.go 压缩成bz2格式
  tar jxvf db.tar.bz2 解压成bz2格式
* zip db.zip *.go 压缩成zip格式
  unzip -d db db.zip 解压到db文件夹，必须指定文件夹