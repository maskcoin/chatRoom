package main

import (
	"fmt"
	"mygithub_code/chatRoom/client/process"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户的密码
var userID int
var userPwd string
var userName string

func main() {
	// 接受用户的选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("----------------------欢迎登录多人聊天系统--------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			err := up.Login(userID, userPwd)
			if err != nil {
				fmt.Println("up.Login(userID, userPwd) err = ", err)
				return
			}

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名称(nickname):")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userID, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
