package mock

//只要用结构体实现了接口中的方法,就实现了这个接口
type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
