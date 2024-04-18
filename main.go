package main

import (
	"context" //ask Navin anna
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err) //see how to add custom log messages later
	}
	defer cli.Close()

	_, err = cli.Put(context.Background(), "rinku", "tingupingu")
	if err != nil {
		log.Fatal(err) //custom message esketit
	}
	fmt.Println("Put operation completed successfully!")
}
