package main

import (
	"fmt"

	"thread_pool/navidPool"
	"thread_pool/workerPool"
)

func main() {
    navidPool.RunManger(20,"file1.txt")
    fmt.Println("**********************************************************")
    workerPool.RunManger(20,"file1.txt")
}