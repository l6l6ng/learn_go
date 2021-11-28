package main

import "fmt"

type student interface {
	notify()
}

type user struct {
}

func (u *user) notify() {
	fmt.Println("hello")
}

func main() {
	u := user{}
	sendNotify(&u)
}

func sendNotify(u student) {
	u.notify()
}
