package main

import "fmt"

func main() {
	//1.第一种建立方法
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	//2.开辟空间定义空map ，m2为空map
	m2 := make(map[string]int)
	fmt.Println(m2)

	//3.变量法定义 m3为nil,go中nil可以参与运算
	var m3 map[string]int

	fmt.Println(m3)
	//遍历方法，key在map中是无序的，map是hashmap
	for k, v := range m {
		fmt.Println(k, v)
	}

	//当取map中不存在的值时，取出一个空串
	courseName := m["course"]
	fmt.Println(courseName)

	//判断map中是否有某个值
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除元素
	name, ok := m["name"]
	fmt.Println("删除之前")
	fmt.Println(name, ok)
	delete(m, "name")
	//当name，ok都定义过时，不能用：定义，当有一个未定义时可以使用
	name, ok = m["name"]
	fmt.Println("删除之后")
	fmt.Println(name, ok)

	//要让map有顺序，要手动排序，将map中数据取到slice中，然后用slice排序

	//map中除了slice，map，function的内建类型都可以作为key
	//map中Struct不包含slice，map，functio都可以作为key

}
