package main

import (
	"container/list"
	"fmt"
	"hash/crc32"
)

func main() {
	res := crc32.ChecksumIEEE([]byte("string"))
	fmt.Println(res)
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0

}
