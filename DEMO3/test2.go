package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var mute sync.Mutex
var wg sync.WaitGroup

// 修改数据
func intAdd(a int) int {
	//定义一个指针接收a地址
	var p *int
	p = &a
	//在a原有基础上加10赋值给*p
	*p = a + 10
	return a

}

// 修改切片中每个元素
func arrUpdate(arr []int) []int {
	var arr1 []int
	for _, v := range arr {
		var p *int
		p = &v
		*p = v * 2
		arr1 = append(arr1, v)

	}
	return arr1

}

// 打印1-10奇数和偶数
func goSyncTest() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {

			if i%2 == 1 {
				fmt.Println("线程1%v", i)
			}

		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i++ {

			if i%2 == 0 {
				fmt.Println("线程2%v", i)
			}

		}
		wg.Done()
	}()
	wg.Wait() // 函数结束前原地听命，计数器变为0才继续

}

// Task 定义任务类型，接收任务名称和执行函数
type Task struct {
	Name string
	Func func() // 任务执行函数
}

// TaskResult 存储任务执行结果（包含执行时间）
type TaskResult struct {
	TaskName string
	Duration time.Duration
	Err      error // 可扩展用于捕获任务执行错误
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []Task
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(name string, taskFunc func()) {
	s.tasks = append(s.tasks, Task{
		Name: name,
		Func: taskFunc,
	})
}

// Run 并发执行所有任务并返回结果
func (s *Scheduler) Run() []TaskResult {
	resultChan := make(chan TaskResult, len(s.tasks))
	defer close(resultChan)

	// 启动协程执行每个任务
	for _, task := range s.tasks {
		go func(t Task) {
			start := time.Now()
			// 执行任务（这里简化处理，未捕获panic，实际应用可添加recover）
			t.Func()
			duration := time.Since(start)

			resultChan <- TaskResult{
				TaskName: t.Name,
				Duration: duration,
				Err:      nil,
			}
		}(task) // 注意：循环变量捕获需要显式传参
	}

	// 收集所有任务结果
	results := make([]TaskResult, 0, len(s.tasks))
	for i := 0; i < len(s.tasks); i++ {
		results = append(results, <-resultChan)
	}

	return results
}

// 接口的定义与实现
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct{}

type Circle struct{}

func (r Rectangle) Area() {
	fmt.Println("这是Area方法")
}

func (r Circle) Perimeter() {
	fmt.Println("这是Perimeter方法")
}

// 组合的使用、方法接收者
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person     Person
}

func (e Employee) PrintInfo() {
	fmt.Println("员工信息为：", e)
}

// 通道的基本使用、协程间通信。
func ChannleAdd() {
	inChann := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	//同时开启协程插入数据
	go func() {
		for i := 1; i <= 10; i++ {
			inChann <- i
			fmt.Println("通道存入数据：", i)
		}
		close(inChann)
		wg.Done()
	}()

	//同时开启协程取出通道数据
	go func() {
		for v := range inChann {
			fmt.Println("通道取出数据：", v)

		}

		wg.Done()
	}()
	wg.Wait()
	fmt.Println("=== 所有数据读取完成 ===")

}

// 通道的缓冲机制
func PrintChannle() {
	producer := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i <= 100; i++ {
			producer <- i
			time.Sleep(100 * time.Millisecond) // 模拟缓冲
		}
		wg.Done()

		defer close(producer)
	}()

	go func() {
		for v := range producer {
			fmt.Println("取到的数据：", v)
		}
		wg.Done()
	}()
	wg.Wait()

}

// sync.Mutex 的使用、并发数据安全
var count int64

func MutexTest() {
	mute.Lock()
	defer mute.Unlock()
	for i := 0; i < 1000; i++ {
		count++
	}
	wg.Done()
}

// 原子操作、并发数据安全。
func MutexTestAtomic() {
	mute.Lock()
	defer mute.Unlock()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
	wg.Done()
}

func main() {

	// 修改数据
	a := 2
	fmt.Println(intAdd(a))

	// 修改切片中每个元素
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arrUpdate(arr))

	//打印1-10奇数和偶数
	goSyncTest()

	// 创建调度器
	scheduler := NewScheduler()

	// 添加示例任务（模拟不同执行时间）
	scheduler.AddTask("任务1", func() {
		time.Sleep(100 * time.Millisecond) // 模拟耗时操作
	})

	scheduler.AddTask("任务2", func() {
		time.Sleep(200 * time.Millisecond)
	})

	scheduler.AddTask("任务3", func() {
		time.Sleep(150 * time.Millisecond)
	})

	// 执行所有任务并获取结果
	start := time.Now()
	results := scheduler.Run()
	totalDuration := time.Since(start)

	// 打印每个任务的执行时间
	fmt.Println("任务执行时间统计：")
	for _, res := range results {
		fmt.Printf("%s: %v\n", res.TaskName, res.Duration)
	}
	// 打印总耗时（应为最长任务的时间，体现并发效果）
	fmt.Printf("\n总执行时间: %v\n", totalDuration)

	// 接口的定义与实现
	var s Rectangle
	s.Area()
	var t Circle
	t.Perimeter()

	// 组合的使用、方法接收者
	e := Employee{
		EmployeeID: 123456,
		Person: Person{
			Name: "张三",
			Age:  20,
		},
	}
	e.PrintInfo()

	// 通道的基本使用、协程间通信。
	//ChannleAdd()
	// 通道的缓冲机制
	//PrintChannle()

	// sync.Mutex 的使用、并发数据安全。
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go MutexTest()
	}
	wg.Wait()
	fmt.Println("MutexTest计数器的值:", count)

	//原子操作、并发数据安全。
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go MutexTestAtomic()
	}
	wg.Wait()
	fmt.Println("MutexTestAtomic计数器的值:", count)

}
