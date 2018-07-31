# yaml学习

### yaml对象
对象的名字与值  
``` yaml
animal: pets
```
### yaml数组
一组连词线开头的行，构成一个数组。  
``` yaml
languages:  
 - Ruby  
 - Perl  
 - Python
```
### yaml 对象中包含对象
``` yaml
websites:
 YAML: yaml.org 
 Ruby: ruby-lang.org 
 Python: python.org 
 Perl: use.perl.org 
```
### yaml 基本类型
数值直接以字面量的形式表示。  
``` yaml
number:23.30
```
布尔值用true和false表示。  
``` yaml
isSet: true
```
null用~表示  
``` yaml
parent: ~ 
```
时间采用 ISO8601 格式   
``` yaml
iso8601: 2001-12-14t21:59:43.10-05:00 
```
日期采用复合 iso8601 格式的年、月、日表示。   
``` yaml
date: 1976-07-31 
```
YAML 允许使用两个感叹号，强制转换数据类型。   
``` yaml
e: !!str 123
f: !!str true
```
字符串默认不使用引号表示。   
``` yaml
str: 这是一行字符串
```
如果字符串之中包含空格或特殊字符，需要放在引号之中。   
``` yaml
str: '内容： 字符串'
```
如果字符串之中包含空格或特殊字符，需要放在引号之中。   
``` yaml
str: '内容： 字符串'
```
其它见:  
http://www.ruanyifeng.com/blog/2016/07/yaml.html