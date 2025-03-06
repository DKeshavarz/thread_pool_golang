package workerPool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	if err != nil{
		fmt.Println(err)
		return
	}
	defer data.Close()
	

	idCounter := 1
	currTime := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}

		arrival, err := strconv.Atoi(parts[0])
		if err != nil{
			fmt.Println(err)
			return
		}
		burst, err := strconv.Atoi(parts[1])
		if err != nil{
			fmt.Println(err)
			return
		}
		
		time.Sleep((time.Duration(arrival - currTime)) * time.Second)
		currTime = arrival
		jobs <- burst
		idCounter++
	}

	
}

func heavyCalculation(n int){
	time.Sleep(time.Second * (time.Duration(n)))
}

func worker(id int, jobs <-chan int,wg *sync.WaitGroup){
	defer wg.Done()

	for val := range jobs {
		log.Printf(" Worker = %-3d   Start with input -> %-6d \n", id , val)
		heavyCalculation(val)
		log.Printf("  _  ID = %-3d   End   with input -> %-6d \n", id , val)
	}
}