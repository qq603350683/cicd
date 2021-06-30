package main

import "testing"

func TestAddUser(t *testing.T) {
	ConnectMysql()

	_, err := AddUser()
	if err != nil {
		t.Error(err)
	}
}

func TestAddUserCache(t *testing.T) {
	ConnectRedis()

	_, err := AddUserCache()
	if err != nil {
		t.Error(err)
	}
}
