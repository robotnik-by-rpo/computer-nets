package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Println("Enter number")
	fmt.Scan(&n)
	solveMaskOfSubNet(n)
	fmt.Println("\n--------\n")
	fmt.Println("Another numbers (0-32)")
	for i:=0; i<33;i++{
		fmt.Printf("%d: ", i)
		solveMaskOfSubNet(i)
	}
}

func solveMaskOfSubNet(n int) {
	mask := strings.Repeat("1",n)+strings.Repeat("0",32-n)
	dMask := []string{mask[0:8], mask[8:16], mask[16:24], mask[24:32]}	
	dRes := make([]int64,4)
	for i, v := range dMask{
		n, _ := strconv.ParseInt(v,2,64)
		dRes[i] = n
	}
	sliceStr := make([]string, 4)
	for i, v := range dRes{
		sliceStr[i] = strconv.Itoa(int(v))
	}
	fmt.Print(strings.Join(sliceStr,"."),"\n")
}