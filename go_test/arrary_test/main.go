package main

import "fmt"



// user在程序中定义一个user类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}



// changeEmail 使用指针接收者实现了一个车方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func main() {
	// user 类型的值可以用来调用
	// 使用值接受者声明的方法
	bill := user{"bill", "bill@email.com"}
	bill.notify()

	//指向user类型值的指针也可以用来调用
	//使用值接收者声明的方法
	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()

	//user类型的值可以用来调用
	//使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	//指向user类型值的指针可以用来调用
	//使用指针接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

	test := user{}
	test.notify()
}
