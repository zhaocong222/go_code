package main

import "fmt"

/*
 * 1:使用 map[string]map[string]string 的map类型
   2:key表示用户名，是唯一的，不可以重复
   3:如果某个用户名存在，就将其密码修改 888888，如果不存在就增加这个用户信息
   4:编写一个函数modifyUser(users map[string]map[string]string, name string)完成上述功能
*/
func modifyUser(users map[string]map[string]string, name string) {
	//判断users中是否有name的key
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "111111"
		users[name]["nickname"] = "昵称" + name

	}

}

func main() {
	user := make(map[string]map[string]string, 10)
	modifyUser(user, "tom")
	modifyUser(user, "mary")

	fmt.Println(user)
}
