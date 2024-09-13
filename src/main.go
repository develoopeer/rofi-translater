package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedResponse struct {
	Translate   string
	Description string
}

func parseCam(word string) {
	url := "https://dictionary.cambridge.org/dictionary/english/" + word
	headers := make(map[string]string)
	headers["Content-Type"] = "json"
	resp := makeRequest(url, headers)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	prompt_string := fmt.Sprintf("\000prompt\x1f%s\n", os.Args[1])
	var count int
	fmt.Print(prompt_string)
	for i, value := range doc.Find(".def.ddef_d.db").EachIter() {
		str := strings.ToLower(strings.ReplaceAll(value.Text(), "\n", ""))
		str = strings.ReplaceAll(str, "\t", "")
		str = strings.ReplaceAll(str, ":", "")
		splitted := strings.Fields(str) // That is ugly man
		var output string
		for _, substring := range splitted {
			output += substring + " "
		}
		fmt.Println(upperFirstLetter(output))
		count = i
	}
	if count == 0 {
		fmt.Println("No output")
	}
}

func main() {
	if len(os.Args) != 1 {
		parseCam(os.Args[1])
	}
}
