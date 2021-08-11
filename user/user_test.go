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
	err := ConnectMysql()
	if err != nil {
		t.Error(err)
	}

	_, err = AddUser()
	if err != nil {
		t.Error(err)
	}

	ggg

	t.Log("mysql successhahahahah")
}


