package navidPool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"thread_pool/queue"
	"time"
)

type Task struct {
	ID          int
	BurstTime   int
	ArrivalTime int
}

var (
	Thread_used []bool
)

func RunManger(workerCnt int, fileName string, queueSize int) {
	now := time.Now()
	Thread_used = make([]bool, workerCnt+1)

	queue := queue.New[Task](queueSize)
	done := make(chan struct{})
	finish := make(chan struct{}, workerCnt)

	go fileReader(queue, fileName, done)
	for i := range workerCnt {
		go worker(queue, done, finish, i)
	}

	for range workerCnt {
		<-finish
	}

	cnt := 0;
	for _,val := range Thread_used {
		if val {
			cnt++
		}
	}
	fmt.Println("totoal exe is : ", time.Since(now), "   and  used ", cnt)
}

func fileReader(q *queue.Queue[Task], fileName string, done chan struct{}) {
	defer close(done)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var Tasks []Task
	scanner := bufio.NewScanner(file)
	idCounter := 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}

		arrival, _ := strconv.Atoi(parts[0])
		burst, _ := strconv.Atoi(parts[1])

		Tasks = append(Tasks, Task{
			ID:          idCounter,
			ArrivalTime: arrival,
			BurstTime:   burst,
		})
		idCounter++
	}

	sort.Slice(Tasks, func(i, j int) bool {
		return Tasks[i].ArrivalTime < Tasks[j].ArrivalTime
	})

	startTime := time.Now()
	for _, Task := range Tasks {
		time.Sleep(time.Until(startTime.Add(time.Duration(Task.ArrivalTime) * time.Second)))
		q.Push(Task)
	}

}

func processTask(task Task, workerID int) {
	Thread_used[workerID] = true
	log.Printf(" Worker %-3d:\t Started task   %-3d  \t->\t (arrived at %-3ds)",
		workerID, task.ID, task.ArrivalTime)

	time.Sleep(time.Duration(task.BurstTime) * time.Second)

	log.Printf(" Worker %-3d:\t Finished task  %-3d  \t->\t (burst %-3ds)",
		workerID, task.ID, task.BurstTime)
}

func worker(q *queue.Queue[Task], done <-chan struct{}, workerDone chan<- struct{}, id int) {
	defer func() { workerDone <- struct{}{} }()

	for {
		select {
		case <-done:
			for {
				task, err := q.Pop()
				if err != nil {
					return
				}
				processTask(task, id)
			}
		default:
			task, err := q.Pop()
			if err != nil {
				continue
			}
			processTask(task, id)
		}
	}
}
