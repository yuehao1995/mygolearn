// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd2testflag

import (
	"fmt"

	"github.com/spf13/cobra"
	)


var versionFlag bool
// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")

		if versionFlag{
			fmt.Println("启动成功")
		}else{
			fmt.Println("启动失败")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	//添加两句flag
	//第一句persistent的作用范围是该命令之后的所有子命令
	//第二句的flag仅支持当前命令使用
	testCmd.PersistentFlags().String("foo", "小白白", "A help for foo")
	testCmd.Flags().String("foolcal", "", "A help for foo")

	//flag支持下面几种用法
	//第一种调用形式   go run main.go test --f
	//第二种调用形式
	//1.go run main.go test --aaa
	//2. go run main.go test -a
	//第三种传入方式与第二种不一样,调用形式和第二种一样，详情见pflag包,
	testCmd.Flags().String("f", "", "test")
	testCmd.Flags().StringP("aaa", "a", "", "test")
	//调用方式  go run main.go test -v true
	//结果调用testcmd的run，结果为:
	//test called
	//启动成功
	testCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Display current version of fabric peer server")
	//结果
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
