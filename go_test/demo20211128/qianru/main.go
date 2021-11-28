//这个示例程序展示如何将一个类型嵌入另一个类型，以及
//内部类型和外部类型的关系
package main

import "fmt"

//notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

//user 再程序里定义一个用户类型
type user struct {
	name  string
	email string
}

//notify 实现了一个可以通过user类型值的指针
//调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

//admin 代表一个拥有权限的管理员用户
type admin struct {
	user  //嵌入类型
	level string
}

//通过 admin 类型值的指针调用的方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotification(&ad)

	//我们可以直接访问内部类型的方法
	ad.user.notify()

	//内部类型的方法也被提升到外部类型
	ad.notify()
}

//sendNotification接受一个实现了notifer接口的值并发送通知
func sendNotification(n notifier) {
	n.notify()
}
