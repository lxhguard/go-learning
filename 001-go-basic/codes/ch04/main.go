package main // 可运行的应用程序

import "fmt"

func main() { // 应用程序的主入口
	array:=[5]string{"a","b","c","d","e"}
	a:=array[2:5] // [2, 5)
	slice1:=make([]string,4,8)
	fmt.Println(slice1)
	fmt.Println(a)
	a[1] ="f"
	for i,v:=range array{
    	fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	fmt.Println(a)
	fmt.Println(array) // 切片使用的底层数组还是原来的数组
}
