package main

import "fmt"

func learn2() {
	fmt.Println("learn2==11") // learn1 修改

	fmt.Println("learn1==11")
}

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}
