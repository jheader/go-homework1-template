package homework02

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 指针
func Add(a *int) {

	*a++
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func DoubleElem(slic *[]int) {

	for indx, item := range *slic {
		fmt.Println("indx:", indx)
		(*slic)[indx] *= 2 //切片索引 [] 的优先级高于解引用 *，加括号才能先解引用拿到切片，再通过索引定位元素。
		fmt.Println("item:", item)
	}

}

// 切片本来就是引用类型
func DoubleElemOne(slic []int) {
	//item只是副本，需要通过切片修改
	for indx, item := range slic {
		fmt.Println("indx:", indx)
		slic[indx] *= 2
		fmt.Println("item:", item)
	}

}

var (
	counter int
	mutex   sync.Mutex
)

func UserMutexAdd() {

	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

// 定义原子计数器（必须为uint64类型，符合atomic包要求）
var counterUint64 uint64

// 原子递增函数（无锁）
func atomicIncrement(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// 原子递增：第一个参数是计数器地址，第二个是递增步长
		atomic.AddUint64(&counterUint64, 1)
	}
}
