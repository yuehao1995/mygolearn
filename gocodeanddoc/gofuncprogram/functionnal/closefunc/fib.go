package main

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		//两步并行,都基于之前的结果ss
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	f() //1
	f() //1
	f() //2
	f() //3
	f() //5
	f() //8
	f() //13
	f() //21
}
