package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r Retriever) Get(url string) string {
	resq, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//处理请求结果
	result, err := httputil.DumpResponse(resq, true)
	resq.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
