package navidPool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"thread_pool/queue"
	"time"
)

func RunManger(workerCnt int, fileName string){
	queue := queue.New[int](100)
	done := make(chan struct{})
	go fileReader(queue, fileName,done)
	
	var wg sync.WaitGroup
	for i := 1 ; i <= workerCnt ; i++ {
		wg.Add(1)
		go worker(queue,done,i,&wg)
	}

	wg.Wait()

}

func fileReader(queue *queue.Queue[int], fileName string,done chan struct{}){
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
		
		queue.Push(num)
	}

	close(done)
}

func heavyCalculation(n int){
	time.Sleep(time.Second * (time.Duration(n)))
}

func worker(queue *queue.Queue[int], done <-chan struct{}, id int,wg *sync.WaitGroup){
	defer wg.Done()

	for {
		for !queue.IsEmpty(){
			val, _ := queue.Pop()
			log.Printf("  _  ID = %-6d   Start with input -> %-6d \n", id , val)
			heavyCalculation(val)
			log.Printf("  _  ID = %-6d   End   with input -> %-6d \n", id , val)
		}

		select{
		case <-done:
			return
		default:
			break;
		}
	}

	
}