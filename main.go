package main

import (
	"context" //ask Navin anna
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"os"
	// "strings"
	"time"
)

func main() {

	args := os.Args

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("Error initializing client:\n", err) //see how to add custom log messages later
	}
	defer cli.Close()

	if len(args) >= 2 {
		switch args[1] {
		case "put":
			err = Put(cli, args[2], args[3])
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
		fmt.Println("Help")
	}
}

func Get(cli *clientv3.Client, key string) *clientv3.GetResponse {
	value, err := cli.Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: log at normal
	return value
}

func Put(cli *clientv3.Client, key string, value string) error {
	_, err := cli.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	fmt.Printf("Put key:%s with value:%s successfully!", key, value)
	// TODO: log at normal
	return nil
}

func Delete(cli *clientv3.Client, key string) error {
	response, err := cli.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	if response.Deleted == 1 {
		fmt.Println("Delete operation completed successfully!")
	} else {
		fmt.Println("Key not found. No action performed")
	}
	// TODO: log at normal
	return nil
}

func listEndpoints(cli *clientv3.Client) error {

	endpoints := cli.Endpoints()
	fmt.Println("Endpoints: \n", endpoints)
	return nil
}

//TODO:Implement a transaction?
