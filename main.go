package main

import (
	"fmt"

	"thread_pool/config"
	"thread_pool/threadPool/navidPool"
	"thread_pool/threadPool/workerPool"
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
	fmt.Println("**********************************************************")
	navidPool.RunManger(numberWorker, "jobs.txt", maxQueueSize)
	fmt.Println("**********************************************************")
	workerPool.RunManger(numberWorker, "jobs.txt")
}
