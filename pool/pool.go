package pool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"thread_pool/queue"
)

func RunManger(workerCnt int, fileName string){
	queue := queue.New[int]()
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

func fibo(n int)int{
	if n <= 1 {
		return n
	}
	return fibo(n-1)+fibo(n-2)
}

func worker(queue *queue.Queue[int], done <-chan struct{}, id int,wg *sync.WaitGroup){
	defer wg.Done()

	for {
		for !queue.IsEmpty(){
			val, _ := queue.Pop()
			log.Println("id =", id, "  find val = ", val, "start")
			new := fibo(val)
			log.Println("id =", id, "  find val = ", val, "end ", new)
		}

		select{
		case <-done:
			return
		default:
			break;
		}
	}

	
}