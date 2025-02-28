package main

import (
	"fmt"

	"thread_pool/config"
	"thread_pool/navidPool"
	"thread_pool/workerPool"
)

var (
	numberWorker int
	maxQueueSize int
)

func init() {
	c, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}

	numberWorker = c.NumWorker
	maxQueueSize = c.MaxQueueSize
}

func main() {
	navidPool.RunManger(numberWorker, "file1.txt")
	fmt.Println("**********************************************************")
	workerPool.RunManger(numberWorker, "file1.txt")
}
