package main

import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户的密码
var userID int
var userPwd string

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
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//根据用户的输入，显示新的提示信息
	if key == 1 {
		// 说明用户要登录
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userID)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		// 先把登录的函数，写到另外一个文件中，比如：login.go
		err := login(userID, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			//进行下一步的逻辑
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
