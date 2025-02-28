package workerPool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

)

func RunManger(workerCnt int, fileName string){
	jobs   := make(chan int)

	var wg sync.WaitGroup
	for i := 1 ; i <= workerCnt ; i++ {
		wg.Add(1)
		go worker(i,jobs,&wg)
	}

	go fileReader(jobs,fileName)

	wg.Wait()
}

func fileReader(jobs chan <- int, fileName string){
	defer close(jobs)

	data,err := os.Open(fileName)
	defer data.Close()

	if err != nil{
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)

		if err != nil{
			fmt.Println(err)
			return
		}
		
		jobs <- num
	}

	
}

func heavyCalculation(n int){
	time.Sleep(time.Second * (time.Duration(n)))
}

func worker(id int, jobs <-chan int,wg *sync.WaitGroup){
	defer wg.Done()

	for val := range jobs {
		log.Printf("  _  ID = %-6d   Start with input -> %-6d \n", id , val)
		heavyCalculation(val)
		log.Printf("  _  ID = %-6d   End   with input -> %-6d \n", id , val)
	}
}