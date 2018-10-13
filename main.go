package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"./free"
)

func main(){
	var name string
	var Domens []string
	const url  = "https://www.nic.ru/whois/?searchWord="
	//name = os.Args[1]
	name = "python"

	file, err := os.Open("domainlist.txt")
	if err != nil {
		log.Print(err)
		fmt.Scan()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		Domens = append(Domens, scanner.Text())
	}
	for _,domen:=range(Domens){
		go free.Freedom(url+name+domen)
	}
	defer file.Close()
	time.Sleep(10000 *time.Millisecond)
	fmt.Println("Done...")
	fmt.Scan(&name)
}