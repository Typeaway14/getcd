package main

import (
	"context" //ask Navin anna
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"time"
)

func main() {

	args := os.Args
	cli := initEtcdClient()
	defer cli.Close()

	if len(args) >= 2 {
		switch args[1] {
		case "put":
			err := Put(cli, args[2], args[3])
			if err != nil {
				log.Fatal("Put operation failed:\n", err)
			}
		case "get":
			value := Get(cli, args[2])
			if value == nil {
				log.Fatal("Get returned empty. Possibly key doesnt exist")
				break
			}
			fmt.Println(string(value.Kvs[0].Value))
		case "delete":
			err := Delete(cli, args[2])
			if err != nil {
				log.Fatal("Delete operation failed:\n", err)
				break
			}
		case "list":
			switch args[2] {
			case "endpoints":
				err := listEndpoints(cli)
				if err != nil {
					log.Fatal(err)
					break
				}
			}
		default:
			fmt.Println("error")

		}
	} else {
		fmt.Println(`
NAME:
	getcd - A minimal client for etcd written in Go

USAGE:
	getcd [COMMAND] {KEY} {VALUE}

COMMANDS:
	get 		Gets the key or a range of keys
	put 		Puts the given key into the store
	list 		Lists all the endpoints 
	`)
	}
}

func initEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("Error initializing client:\n", err)
	}
	return cli
}

func Get(cli *clientv3.Client, key string) *clientv3.GetResponse {
	value, err := cli.Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Get operation on key: ", key, " successful\n")
	return value
}

func Put(cli *clientv3.Client, key string, value string) error {
	_, err := cli.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	fmt.Printf("Put key: %s with value: %s successfully!\n", key, value)
	log.Print("Put operation on key: ", key, " with value: ", value, " successful")
	return nil
}

func Delete(cli *clientv3.Client, key string) error {
	response, err := cli.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	if response.Deleted == 1 {
		fmt.Println("Delete operation completed successfully!")
		log.Print("Delete operation on key: ", key, " successful")
	} else {
		fmt.Println("Key not found. No action performedn")
		log.Print("Delete operation failed. Key: ", key, " not found")
	}
	return nil
}

func listEndpoints(cli *clientv3.Client) error {

	endpoints := cli.Endpoints()
	fmt.Println("Endpoints: \n", endpoints)
	return nil
}
