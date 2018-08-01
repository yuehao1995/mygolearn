package main

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/spf13/viper"
)

const cmdRoot = "core"

/*
1 viper.SetConfigName 设置配置文件的名字

2 viper.AddConfigPath 设置配置文件的路径，为什么是add不是set呢，它的实现是append，就是它可以支持多目录搜索

3 viper.ReadInConfig 读取配置文件的内容然后放在viper.config内
*/
/*
读取内容
单个变量的读取比较简单，直接GetX(key)就行了，比如GetString，GetBool等
viper.GetString("encoding")
viper还支持结构体读取
viper.Unmarshal(&conf)
*/
func main() {
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./")
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	for _, p := range filepath.SplitList(gopath) {
		fmt.Println("-------------", p)
		peerpath := filepath.Join(p, "src/mygolearn/myviper/test_viper")
		viper.AddConfigPath(peerpath)
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file: %s\n", cmdRoot, err))
	}

	fullstate := viper.GetString("statetransfer.timeout.fullstate")
	fmt.Println("fullstate:", fullstate)

	abcdValuea := viper.GetString("peer.abcd")
	fmt.Println("abcdValuea is:", abcdValuea)

}
