package free

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)
func Freedom(address string){
	res, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Print(address)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
		return
	}
	result:=doc.Find("._1_uOq").Contents().Text()
	if result==""{
		name:=address[strings.Index(address,"=")+1:]
		result = name + " is free"
		fmt.Println(result)
	}
}

