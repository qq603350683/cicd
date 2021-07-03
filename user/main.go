package main

import "fmt"

func main() {
	var b bool
	var err error

	err = ConnectMysql()
	if err != nil {
		panic(err)
	}

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


