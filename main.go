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
	//I use the whois service to get domain information
	const url  = "https://www.nic.ru/whois/?searchWord="
	name = os.Args[1]
	//domenslist file stores some domain zones.If you want to add a new domain to the search database, just edit this file.
	file, err := os.Open("domainlist.txt")
	if err != nil {
		log.Print(err)
		fmt.Scan()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
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
