# **Navid Pool Project Report**

This report provides an overview of the **Navid Pool** project, which is a thread-safe worker pool implementation in Go. The project is named after the teacher assistant, Navid, who assigned this project. It includes a custom mutex, a thread-safe queue, and two worker pool implementations: one using a custom approach (`navidPool`) and another using Go's native concurrency features (`workerPool`).



## **Table of Contents**
1. [Project Overview](#project-overview)
2. [Installation and Run](#installation-and-run)
3. [Project Structure](#project-structure)
5. [How to Use the Config File](#how-to-use-the-config-file)
6. [How to Use the Task File](#how-to-use-the-task-file)
7. [Acknowledgments](#acknowledgments)



## **Project Overview**
The **Navid Pool** project is designed to simulate a worker pool that processes tasks from a file. Each task has an arrival time and a burst time. The project includes:
- A custom mutex implementation (`atomicMutex` and `chanMutex`).
- A thread-safe queue for task management.
- Two worker pool implementations:
  - `navidPool`: A custom worker pool using the thread-safe queue.
  - `workerPool`: A worker pool using Go's native concurrency features (`sync.WaitGroup` and channels).
- A task generator to create random tasks for testing.


## **Installation and Run**
### **Prerequisites**
- Go installed on your system (version 1.16 or higher).

### **Steps to Run the Project**
1. **Clone the Project**:
   - Clone the project repository to your local machine.
   - Navigate to the project directory:
     ```bash
     cd path/to/navid_pool_project
     ```

2. **Build and Run**:
   - Run the project using the following command:
     ```bash
     go mod tidy
     go run main.go
     ```

3. **Output**:
   - The program will generate tasks, process them using both worker pools, and log the results to the console or a specified output file.

---

## **Project Structure**
The project is organized as follows:
```
├── config
│   ├── config.json
│   └── loadConfig.go
├── genarator
│   ├── genarator.go
│   └── generator_test.go
├── go.mod
├── main.go
├── mutex
│   ├── atomicMutex.go
│   ├── chanMutex.go
│   ├── mutex.go
│   └── mutex_test.go
├── queue
│   ├── queue.go
│   └── queu_test.go
├── resources.txt
└── threadPool
    ├── navidPool
    │   └── pool.go
    └── workerPool
        └── pool.go
```


- #### **1. Config Package Explanation**

The `config` package is responsible for loading configuration settings from a JSON file. It defines a `Config` struct with fields like `NumWorker`, `MaxQueueSize`, `InFile`, and `OutFile`. The `LoadConfig` function reads the JSON file, parses it into the `Config` struct, and returns the configuration for use in the project. While the package does not include a dedicated test file, its functionality is critical for dynamically configuring the worker pool and file paths. The `config.json` file specifies parameters such as the number of workers, queue size, and input/output file paths, making the project flexible and easy to customize.

---
- #### **2. `mutex` Package**


The `mutex` package provides two implementations of a mutual exclusion lock (mutex): `atomicMutex` and `chanMutex`. The `atomicMutex` uses Go's `sync/atomic` package to implement a lock with atomic operations, while the `chanMutex` uses a buffered channel of size 1 to achieve thread-safe locking. Both implementations satisfy the `Mutex` interface, which defines `Lock()` and `Unlock()` methods. The package also includes test files (`mutex_test.go`) to verify the correctness of the mutex implementations. These tests simulate concurrent access to a shared resource (e.g., a counter) and ensure that the mutex prevents race conditions.

---
- #### **3. Generator Package Explanation**

The `generator` package is responsible for creating test task files with random arrival and burst times. It includes the `GenerateFile` function, which generates a specified number of tasks and writes them to a file in the format `<arrival_time> <burst_time>`. The package also contains a test file (`generator_test.go`) with the `TestGenerateFile` method. This test verifies that the generated file has the correct number of tasks, adheres to the required format, and ensures that the values fall within the specified ranges. After validation, the test automatically deletes the generated file to clean up. This package is essential for generating test data to benchmark and stress-test the worker pool implementations.

---
- #### **4. Navid Pool Package Explanation**

The `navidPool` package implements a custom worker pool using the thread-safe `queue` package to manage tasks. It reads tasks from a file, assigns them to workers based on arrival time, and logs their execution. The pool uses a mutex to ensure thread safety and tracks worker usage for performance analysis. While the package does not include a dedicated test file, its functionality is implicitly tested through the main program's execution, which processes tasks and logs their start and finish times. This package demonstrates a custom approach to worker pool implementation, contrasting with Go's native concurrency features.

---
- #### **5. Worker Pool Package Explanation**
The `workerPool` package implements a worker pool using Go's native concurrency features, such as `sync.WaitGroup` and channels. It reads tasks from a file, assigns them to workers, and logs their execution. The pool uses a channel to distribute tasks among workers and tracks worker usage for performance analysis. Like `navidPool`, this package does not include a dedicated test file but is tested implicitly through the main program's execution. It serves as a comparison to the custom `navidPool`, showcasing Go's built-in concurrency tools for task management.

---
- #### **6. Main Package Explanation**

The `main.go` file serves as the entry point for the **Navid Pool** project. It initializes the program by loading configuration settings from `config.json` using the `config` package. These settings include the number of workers (`NumWorker`), queue size (`MaxQueueSize`), input file path (`InFile`), and output file path (`OutFile`). The program then generates tasks using the `generator` package and processes them using both the `navidPool` (custom implementation) and `workerPool` (Go-native implementation). It logs the total execution time and worker usage for performance comparison. While `main.go` does not have dedicated test files, its functionality is validated through the execution of the entire project, ensuring all components work together seamlessly. The `config.json` file makes the program flexible and easy to customize for different use cases.

## **How to Use the Config File**

The `config.json` file is used to configure the **Navid Pool** project. It contains the following fields:

1. **`num_worker`**: Specifies the number of workers in the pool.
2. **`queue_size`**: Defines the maximum size of the task queue.
3. **`file_int`**: Specifies the input file pathcontaining tasks.
4. **`file_out`**: Specifies the output file path where logs will be written. This field is optional, if you ommit this part logs will write in your console.

#### **Example `config.json`**
```json
{
    "num_worker": 100,
    "queue_size": 10,
    "file_int": "in.txt",
    "file_out": "out.txt"
}
```

## **How to Use the Task File**

The task file (e.g., `jobs.txt`) contains the tasks to be processed by the worker pools. Each task is represented by two numbers: **arrival time** and **burst time**, formatted as `<arrival_time> <burst_time>`.

#### **Task File Format**
```
0 2
0 1
3 5
3 1
8 5
9 10
10 6
```


### **Generating the Task File**
The `generator` package provides a function to automatically generate task files with random arrival and burst times. Use the `GenerateFile` function as follows:

```go
time, err := genarator.GenerateFile([]int{0, 0}, []int{1, 10}, 10000, "jobs.txt")
if err != nil {
    panic(err)
}
```
#### *Parameters:*

`arrivalTimeRange`: A slice specifying the range for arrival times (e.g., [0, 0] for fixed arrival times).

`exeTimeRange`: A slice specifying the range for burst times (e.g., [1, 10] for burst times between 1 and 10 seconds).

`numberOfLines`: The number of tasks to generate.

`outputFile`: The name of the file to save the tasks (e.g., jobs.txt).

## **Acknowledgments**

We would like to extend our heartfelt thanks to our teacher assistants, **Navid** and **Amir Reza**, for their guidance, support, and patience throughout this project. Your insights and feedback were invaluable in helping us navigate the complexities of this operating systems project.

This project was developed by **Saeed** and **Daniel** as part of the **Winter 2025** course. We hope our work reflects the effort and dedication we put into understanding and implementing the concepts taught in class.

```
In the land of code, where logic flows,
Navid and Amir Reza, our guides, arose.
With wisdom and patience, they lit the way,
Through threads and queues, night and day.
```

