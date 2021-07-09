package main

import "fmt"

func main() {
	fmt.Printf("develop udpate something....")
	//password := "61fcb62dfa4c554fc68863c744897aa0"

	//fmt.Println(password)

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

	_, err = UpdateUserNickname(1, "baiaisdias")
	if err != nil {
		fmt.Println(err)
		return
	}
}


