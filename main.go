package main

import (
	"context" //ask Navin anna
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/client/v3"
	"google.golang.org/genproto/googleapis/api/error_reason"
)

func main() {

	args := os.Args

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err) //see how to add custom log messages later
	}
	defer cli.Close()

	// _, err = cli.Put(context.Background(), "rinku", "tingupingu")
	// if err != nil {
	// 	log.Fatal(err) //custom message esketit
	// }
	fmt.Println("Put operation completed successfully!")
	
	//TODO: switch case type tings

	err = Put(cli, args)
	if err != nil {
		log.Fatal(err) //custom message esketit x2
	}
}

func Put(cli *clientv3.Client,args) error {
	_,err := cli.Put(context.Background(), key, value)
	if err != nil{
		return err
	}
	fmt.Println("Put Operation completed successfully!")
	// TODO: log at normal 
	return nil
}

func Delete(cli *clientv3.Client, args) error
{
	_, err := cli.Delete(context.Background(),key)
	if err != nil{
		return err
	}
	fmt.Println("Delete operation completed successfully!")
	// TODO: log at normal
	return nil
}

