package main

func main() {
	ConnectMysql()

	ConnectRedis()

	var b bool
	var err error

	b, err = AddUser()
	if err != nil {
		panic(err)
	}

	_, err = AddUserCache()
	if err != nil {
		panic(err)
	}
}


