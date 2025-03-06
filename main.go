package main

import (
	"fmt"
	"log"
	"os"

	"thread_pool/config"
	"thread_pool/genarator"
	"thread_pool/threadPool/navidPool"
	"thread_pool/threadPool/workerPool"
)

var (
	numberWorker int
	maxQueueSize int
	InFile       string
	OutFile      string
)

func init() {
	c, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}

	numberWorker = c.NumWorker
	maxQueueSize = c.MaxQueueSize
	InFile = c.InFile
	OutFile = c.OutFile
	
	if c.OutFile != "" {
		OutFile, err := os.Create(c.OutFile)
		if err != nil {
			panic(err)
		}
		os.Stdout = OutFile
		log.SetOutput(OutFile)
	}

}

func main() {
	time, err := genarator.GenerateFile([]int{0,0}, []int{1,10}, 10000, InFile)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("total time without worker pool", time , "seconds")
	fmt.Println("------------------------- Navid pool (base project) -------------------------")
	navidPool.RunManger(numberWorker, InFile, maxQueueSize)
	fmt.Println("-----------------------  worker pool (golang approach) ------------------------")
	workerPool.RunManger(numberWorker, InFile)
}
