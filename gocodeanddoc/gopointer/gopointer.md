# go指针
### go指针定义及使用
var a int =2  
var pa *int=&a  
*pa=3  
fmt.Println(a) a为3 
###### go语言指针不能做运算  
###### go语言只有值传递这一种方式
###### go通过指针传递展现出相当于引用传递的效果
###### go语言 