package free

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"sync"
)

func Freedom(address string, wg *sync.WaitGroup, Msg *string) {
	defer wg.Done()
	client := &http.Client{}
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.UserAgent()
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_5_8) AppleWebKit/536.17+ (KHTML, like Gecko) iCab/5.0 Safari/533.16")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Print(address + " not answer")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Print(err)
		return
	}
	
	 //this div class contains domain usage information
	result := doc.Find("._1_uOq").Contents().Text()
	
	//if it is empty, then this domain is free
	if result == "" {
		name := address[strings.Index(address, "=")+1:]
		result = name + " is free\n"
		fmt.Print(result)
		*Msg += result
	}
}
