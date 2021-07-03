package main

import "fmt"

func main() {
	ConnectMysql()

	ConnectRedis()

	var b bool
	var err error

	b, err = AddUser()
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	_, err = AddUserCache()
	if err != nil {
		panic(err)
	}
}


