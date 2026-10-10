package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var mac string
	fmt.Print("mac: ")
	fmt.Scan(&mac)
	file, err := os.Open("mac-vendor.txt")

	mac = strings.ToUpper(strings.Join(strings.Split(mac,"-"),"")[0:6])
	if len(mac) < 6{
		fmt.Println("Строка меньше 6 символов")
		return
	}
	fmt.Println(mac)
	if err != nil{
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := strings.Trim(scanner.Text(),"\n")
		if strings.Contains(line, mac){
			brand := strings.Split(line, "\t")
			fmt.Print(brand[1])
			return
		}

	}

	fmt.Println("Нет бренда")

}