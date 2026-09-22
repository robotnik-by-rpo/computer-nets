package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var res int64
	file, err := os.Open("blocked-networks.txt")

	if err != nil{
		fmt.Println("Error reading file")
		return
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		pref, _ := strconv.Atoi(strings.Split(strings.Trim(line,"\n"),"/")[1])
		res += int64(math.Pow(2.0,float64(32.0-pref)))
	}
	fmt.Println(res)
}