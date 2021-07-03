package main

import "fmt"

func main() {
	err2 := ConnectMysql()
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	var b bool
	var err error

	ConnectRedis()

	b, err = AddUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b)

	_, err = AddUserCache()
	if err != nil {
		fmt.Println(err)
		return
	}
}


