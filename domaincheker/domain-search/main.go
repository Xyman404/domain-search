package domain_search

import (
	"./free"
	"bufio"
	"log"
	"os"
	"sync"
)


func DomainSearch(name string)(string){
	var Domens []string
	var wg sync.WaitGroup
	var result string
	result = ""

	//I use the whois service to get domain information
	const url  = "https://www.nic.ru/whois/?searchWord="
	//domenslist file stores some domain zones.If you want to add a new domain to the search database, just edit this file.

	file, err := os.Open("domainlist.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Domens = append(Domens, scanner.Text())
	}

	for _,domen:=range(Domens){
		wg.Add(1)
		go free.Freedom(url+name+domen,&wg,&result)
	}

	wg.Wait()
	return result
}
