package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	cli := initEtcdClient()
	err := Put(cli, "test_key", "test_value")
	if err != nil {
		t.Fatalf("Put function returned error: %v", err)
	}
}

func TestGet(t *testing.T) {
	cli := initEtcdClient()
	value := Get(cli, "test_key")
	if value == nil {
		t.Fatal("Expected value to be returned, got nil")
	}
}

func TestDelete(t *testing.T) {
	cli := initEtcdClient()
	err := Delete(cli, "test_key")
	if err != nil {
		t.Fatalf("Delete function returned error: %v", err)
	}
}

func TestListEndpoints(t *testing.T) {
	cli := initEtcdClient()
	err := listEndpoints(cli)
	if err != nil {
		t.Fatalf("ListEndpoints function returned error: %v", err)
	}
}
