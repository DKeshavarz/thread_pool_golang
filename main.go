package main

import (
	"fmt"
	"os"

	"thread_pool/config"
	"thread_pool/threadPool/navidPool"
	"thread_pool/threadPool/workerPool"
)

var (
	numberWorker int
	maxQueueSize int
	InFile       string
)

func init() {
	c, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}

	numberWorker = c.NumWorker
	maxQueueSize = c.MaxQueueSize
	InFile = c.InFile

	if c.OutFile != "" {

		file, err := os.Create(c.OutFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		os.Stdout = file
	}
}

func main() {
	fmt.Println("**********************************************************")
	navidPool.RunManger(numberWorker, InFile, maxQueueSize)
	fmt.Println("**********************************************************")
	workerPool.RunManger(numberWorker, InFile)
}
