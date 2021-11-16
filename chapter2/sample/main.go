package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	//
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
)

const sdf = "sds"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

type test struct {
	name string
}

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	///	search.Run("president")

	Integer := int(0xFFFFFFFF)
	qwer := "ssss"
	fmt.Println("sss")
	fmt.Println(strings.Replace(qwer, "s", "ww", 1))
	println(len(sdf))
	println(Integer >> 30)
	println(swap("7", "9"))
	s2, s3 := swap("2", "1")
	println(s2, s3)

	nums := []int{}
	println(len(nums))

	var ptr *int = &Integer
	println(ptr)
	fmt.Println(ptr)

	var ttest test = test{"saaaass"}
	fmt.Println(ttest.name)
	fmt.Println((&ttest).name)
	//
	//var numbers0  = [...] int {1,2,3}
	//var numbers1 [3] int
	//numbers1[2] = 1
	//

	//fmt.Println(cap(numbers0), cap(numbers1), len(numbers0), len(numbers1))
	//fmt.Println(numbers1[2:3])

	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 9)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)

	for i, n := range numbers {
		fmt.Println(i, n)
	}

	printSlice(numbers)

	numss := make([]int, 5, 5)
	numss[4] = 2
	printSlice(numss)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)

	myMap := map[string]string{"sss": "ssf", "sssf": "sfef"}
	myMap["asd"] = "sdf"
	s := myMap["asd"]

	for s2, value := range myMap {
		fmt.Println(s2, value)
	}

	fmt.Println(myMap, s)
	fmt.Println(float32(1243.222))
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func swap(a, b string) (string, string) {
	return b, a
}
