package main

import "fmt"

func main() {
	err2 := ConnectMysql()
	if err2 != nil {
		panic(err)
	}

	var b bool
	var err error

	ConnectRedis()

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


