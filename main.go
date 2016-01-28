package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	res, err := http.Get("http://www.litquotes.com/Random-Quote.php")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	quoteRE := regexp.MustCompile(`</TABLE><p><b>(.*)</b><br><br><i><A HREF="/quote_title_resp.*>(.*)</i></a> by <A HREF="/quote_author.*>(.*)</a></P>`)
	match := quoteRE.FindSubmatch(body)
	if match == nil {
		log.Fatal("No match")
	}
	fmt.Printf("%s\nfrom \"%s\" by %s\n", string(match[1]), string(match[2]), string(match[3]))
}
