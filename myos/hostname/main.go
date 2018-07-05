package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Hostname())
}

func RunEnv() (env string, name string) {
	host := map[string]string{
		"aghfkjhla":     "***正式服",
		"ZYH_aghfkjhla": "RTDYUAIHADI", //加前缀正式服
		"afklfdj":       "***(TEST)",   // 测试服
		"ZYH_afklfdj":   "TEST",
		"dev":           "ZYH(DEV)", // 本地开发服
		"ZYH_dev":       "ZYHDEV",
		"gfskajhsjls":   "afhklALPHA", // 预发服
		"ZYH_yd_crm":    "ZYHALPHA",
	}
	n, err := os.Hostname()
	if err != nil {
		env, name = host["ENV_dev"], host["dev"]
		return
	}
	_, ok := host[n]
	if ok {
		env, name = host["ENV_"+n], host[n]
		return
	}
	fmt.Println("----RunEnv", n, env, name)
	env, name = host["ENV_dev"], host["dev"]
	return
}
