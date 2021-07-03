package main

import "testing"

func TestAddUserCache(t *testing.T) {
	ConnectRedis()

	_, err := AddUserCache()
	if err != nil {
		t.Error(err)
	}

	t.Log("redis success")
}

func TestAddUser(t *testing.T) {
	ConnectMysql()

	_, err := AddUser()
	if err != nil {
		t.Error(err)
	}

	t.Log("mysql success")
}


